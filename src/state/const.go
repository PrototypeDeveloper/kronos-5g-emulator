// Constants of package state.

package state

// Defaults applied when a key is not set.
const (
	// defaultSubscription is the number of subscribers when ue.subscription is not set.
	defaultSubscription = 1
	// deprecatedSessionID is the PDU session ID given to the deprecated ue.type / ue.apn.
	deprecatedSessionID = 1
	// defaultGNBIDLength is the gNB ID length in bits when gnodeb.gnodebidlength is not set.
	defaultGNBIDLength = 32
)

// PLMN and subscriber identities.
const (
	// mccDigits is the number of digits of an MCC.
	mccDigits = 3
	// mncMinDigits and mncMaxDigits bound the number of digits of an MNC.
	mncMinDigits = 2
	mncMaxDigits = 3
	// imsiDigits is the number of digits of an IMSI.
	imsiDigits = 15
	// imeisvDigits is the number of digits of an IMEISV.
	imeisvDigits = 16
	// minSubscription is the smallest allowed ue.subscription.
	minSubscription = 1
)

// SUCI.
const (
	// suciNullScheme, suciProfileA and suciProfileB are the protection scheme IDs.
	suciNullScheme = 0
	suciProfileA   = 1
	suciProfileB   = 2
	// maxSUCIKeyID is the largest home network public key ID.
	maxSUCIKeyID = 255
	// maxRoutingIndicatorDigits is the largest number of routing indicator digits.
	maxRoutingIndicatorDigits = 4
)

// Authentication credentials.
const (
	// keyHexDigits is the length of K and OP/OPc in hex digits (128 bits).
	keyHexDigits = 32
	// amfHexDigits is the length of the AMF field in hex digits.
	amfHexDigits = 4
	// amfSeparationBit is the separation bit (the top bit) of the AMF field.
	amfSeparationBit = 0x8000
)

// S-NSSAIs and PDU sessions.
const (
	// maxSST is the largest slice/service type.
	maxSST = 255
	// sdHexDigits is the length of an SD in hex digits.
	sdHexDigits = 6
	// minPDUSessionID and maxPDUSessionID bound the PDU session ID.
	minPDUSessionID = 1
	maxPDUSessionID = 15
)

// gNB ID, NCI, RAN node name and TAC.
const (
	// minGNBIDLength and maxGNBIDLength bound the gNB ID length in bits.
	minGNBIDLength = 22
	maxGNBIDLength = 32
	// maxGNBIDHexDigits is the largest number of hex digits of gnodeb.gnodebid.
	maxGNBIDHexDigits = 8
	// nciBits is the length of the NR Cell Identity in bits.
	nciBits = 36
	// nciHexDigits is the length of gnodeb.nci in hex digits.
	nciHexDigits = 9
	// maxRANNodeNameLength is the largest RANNodeName length in characters.
	maxRANNodeNameLength = 150
	// minPrintableASCII and maxPrintableASCII bound the characters of the RAN node name.
	minPrintableASCII = 0x20
	maxPrintableASCII = 0x7e
	// maxTAC is the largest TAC (3 octets).
	maxTAC = 0xffffff
)

// Endpoints.
const (
	// maxPort is the largest port number.
	maxPort = 65535
)
