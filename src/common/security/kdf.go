package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

// KDF is the generic key derivation function of TS 33.220 Annex B.2:
// HMAC-SHA-256(Key, FC || P0 || L0 || P1 || L1 || ...), Li being 2 octets.
func KDF(key []byte, fc byte, params ...[]byte) []byte {
	s := []byte{fc}
	for _, p := range params {
		s = append(s, p...)
		s = append(s, byte(len(p)>>8), byte(len(p)))
	}
	h := hmac.New(sha256.New, key)
	h.Write(s)
	return h.Sum(nil)
}

// ServingNetworkName is "5G:mnc<MNC>.mcc<MCC>.3gppnetwork.org" with a
// three-digit MNC (TS 24.501 §9.12.1).
func ServingNetworkName(mcc, mnc string) string {
	if len(mnc) == shortMNCDigits {
		mnc = "0" + mnc
	}
	return fmt.Sprintf("5G:mnc%s.mcc%s.3gppnetwork.org", mnc, mcc)
}

// concat returns a new slice holding a followed by b.
func concat(a, b []byte) []byte { return append(append([]byte(nil), a...), b...) }

// ResStar derives RES* (or XRES*) from RES (TS 33.501 A.4): the 128 least
// significant bits of KDF(CK||IK, 0x6B, SN name, RAND, RES).
func ResStar(ck, ik []byte, snn string, rand, res []byte) []byte {
	return KDF(concat(ck, ik), fcResStar, []byte(snn), rand, res)[16:]
}

// Kausf derives K_AUSF for 5G AKA (TS 33.501 A.2).
func Kausf(ck, ik []byte, snn string, sqnXorAK []byte) []byte {
	return KDF(concat(ck, ik), fcKausf, []byte(snn), sqnXorAK)
}

// Kseaf derives K_SEAF (TS 33.501 A.6).
func Kseaf(kausf []byte, snn string) []byte {
	return KDF(kausf, fcKseaf, []byte(snn))
}

// Kamf derives K_AMF from K_SEAF, the SUPI digits and ABBA (TS 33.501 A.7).
func Kamf(kseaf []byte, supi string, abba []byte) []byte {
	return KDF(kseaf, fcKamf, []byte(supi), abba)
}

// NASKeys derives K_NASenc and K_NASint, the 128 least significant bits of
// KDF(K_AMF, 0x69, distinguisher, algorithm identity) (TS 33.501 A.8).
func NASKeys(kamf []byte, cipheringAlg, integrityAlg uint8) (enc, integrity [16]byte) {
	copy(enc[:], KDF(kamf, fcAlgKey, []byte{algNASEnc}, []byte{cipheringAlg})[16:])
	copy(integrity[:], KDF(kamf, fcAlgKey, []byte{algNASInt}, []byte{integrityAlg})[16:])
	return
}

// KgNB derives K_gNB from K_AMF and the uplink NAS COUNT (TS 33.501 A.9;
// access type 1 = 3GPP access).
func KgNB(kamf []byte, ulCount uint32) []byte {
	var c [4]byte
	binary.BigEndian.PutUint32(c[:], ulCount)
	return KDF(kamf, fcKgNB, c[:], []byte{accessType3GPP})
}
