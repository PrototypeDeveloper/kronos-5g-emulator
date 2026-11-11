// Constants of package security.

package security

// Key and identity lengths.
const (
	// milenageKeyLen is the length of K, OP and OPc in octets (128 bits).
	milenageKeyLen = 16
	// maxRoutingIndicatorDigits is the largest number of routing indicator
	// digits; shorter ones are padded with F to this many nibbles.
	maxRoutingIndicatorDigits = 4
	// imeisvDigits is the number of digits of an IMEISV.
	imeisvDigits = 16
	// shortMNCDigits is the length of a 2-digit MNC, which the serving
	// network name pads to three digits.
	shortMNCDigits = 2
)

// 5GS mobile identity values (TS 24.501 §9.11.3.4).
const (
	// identityTypeSUCI is the first octet of a SUCI: SUPI format IMSI, type of identity SUCI.
	identityTypeSUCI byte = 0x01
	// identityTypeIMEISV is the type of identity IMEISV (with the odd/even indicator clear).
	identityTypeIMEISV byte = 0x05
	// identityTypeIMEI is the type of identity IMEI, with the odd indicator (15 digits).
	identityTypeIMEI byte = 0x08 | 0x03
	// imeiBodyDigits is the number of digits of an IMEI before its check digit (TAC and SNR).
	imeiBodyDigits = 14
	// luhnWeight doubles every second digit of the IMEI check digit (TS 23.003 annex B).
	luhnWeight = 2
	// decimalBase is the base of the digits.
	decimalBase = 10
	// bcdFiller pads BCD digits to a whole number of nibbles.
	bcdFiller = 0xf
)

// KDF function codes FC (TS 33.501 Annex A).
const (
	fcKausf   byte = 0x6A // K_AUSF (A.2)
	fcResStar byte = 0x6B // RES* and XRES* (A.4)
	fcKseaf   byte = 0x6C // K_SEAF (A.6)
	fcKamf    byte = 0x6D // K_AMF (A.7)
	fcAlgKey  byte = 0x69 // algorithm keys (A.8)
	fcKgNB    byte = 0x6E // K_gNB (A.9)
)

// Algorithm type distinguishers (TS 33.501 Table A.8-1).
const (
	algNASEnc byte = 0x01
	algNASInt byte = 0x02
)

// accessType3GPP is the access type distinguisher of 3GPP access in the
// K_gNB derivation (TS 33.501 A.9).
const accessType3GPP byte = 0x01

// SUCI protection schemes (TS 33.501 Annex C).
const (
	SchemeNull     = 0 // the MSIN in clear
	SchemeProfileA = 1 // ECIES with X25519
	SchemeProfileB = 2 // ECIES with secp256r1
)

// ECIES of the SUCI profiles (TS 33.501 C.3.4): the KDF (ANSI X9.63, SHA-256)
// gives the AES-128-CTR key, its initial counter block and the HMAC-SHA-256
// key; the MAC tag is the first 64 bits.
const (
	eciesEncKeyLen = 16
	eciesICBLen    = 16
	eciesMacKeyLen = 32
	eciesMacLen    = 8
	kdfCounterLen  = 4 // octets of the X9.63 counter
)

// Public key encodings.
const (
	x25519KeyLen      = 32   // an X25519 public key
	p256CompressedLen = 33   // a compressed secp256r1 point: 0x02 / 0x03, then x
	p256FullLen       = 65   // an uncompressed one: 0x04, x, y
	p256CoordLen      = 32   // octets of a coordinate
	p256Uncompressed  = 0x04 // the prefix of an uncompressed point
	p256EvenY         = 0x02 // the prefix of a compressed point with an even y (0x03: odd)
)
