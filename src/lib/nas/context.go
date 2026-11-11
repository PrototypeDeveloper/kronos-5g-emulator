package libnas

import "libnas/nasType"

type SecurityHeaderType = nasType.SecurityHeaderType
type Direction = nasType.Direction
type Count = nasType.Count

const (
	SecurityHeaderPlain                                       = nasType.SecurityHeaderPlain
	SecurityHeaderIntegrityProtected                          = nasType.SecurityHeaderIntegrityProtected
	SecurityHeaderIntegrityProtectedAndCiphered               = nasType.SecurityHeaderIntegrityProtectedAndCiphered
	SecurityHeaderIntegrityProtectedWithNewContext            = nasType.SecurityHeaderIntegrityProtectedWithNewContext
	SecurityHeaderIntegrityProtectedAndCipheredWithNewContext = nasType.SecurityHeaderIntegrityProtectedAndCipheredWithNewContext
	Uplink                                                    = nasType.Uplink
	Downlink                                                  = nasType.Downlink
)

type EncodeContext struct {
	// Observe sees the original plaintext after successful protection, without consuming another COUNT.
	Observe            func(wire, plain []byte)
	Security           *SecurityContext
	SecurityHeaderType SecurityHeaderType
	Direction          Direction
}
type DecodeContext struct {
	// Observe is called only after successful verification and decoding.
	Observe   func(wire, plain []byte)
	Security  *SecurityContext
	Direction Direction
}

// SecurityContext belongs to one endpoint and one NAS access connection.
// The caller serializes access and supplies already-derived NAS keys. COUNTs
// are next-unused values and are committed only when Encode/Decode succeeds.
// A new-context header never resets an existing context: supply a fresh context
// explicitly when the protocol procedure has established fresh keys.
type SecurityContext struct {
	IntegrityAlgorithm uint8
	CipheringAlgorithm uint8
	KnasInt            [16]byte
	KnasEnc            [16]byte
	UplinkCount        Count
	DownlinkCount      Count
	// NASConnectionID is the BEARER input: 1=3GPP, 2=non-3GPP. Zero defaults to 1.
	NASConnectionID uint8
}
