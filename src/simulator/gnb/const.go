// Constants of package gnb.

package gnb

// Node identity and user location.
const (
	nciBits        = 36         // bits of an NR Cell Identity
	amfSetIDBits   = 10         // bits of an AMF Set ID
	amfPointerBits = 6          // bits of an AMF Pointer
	sdLen          = 3          // octets of the SD of an S-NSSAI
	ntpEpochOffset = 2208988800 // seconds from the NTP epoch (1900) to the Unix epoch (1970)
)

// The ngKSI and the half octet that carries it.
const (
	noKey         = 0x7  // ngKSI "no key is available" (TS 24.501 §9.11.3.32)
	ksiMask       = 0x7  // the key set identifier of an ngKSI (without TSC)
	ngKSIShift    = 4    // position of an ngKSI in the high half octet
	lowHalfOctet  = 0x0f // mask of the low half octet
	highHalfOctet = 0xf0 // mask of the high half octet
)

// 5GMM message fields.
const (
	identityTypeMask   = 7    // type of identity of a 5GS mobile identity
	gutiLen            = 11   // octets of the 5GS mobile identity value of a 5G-GUTI
	sTMSIIdentityOctet = 0xf4 // first octet of a 5G-S-TMSI mobile identity: spare bits, type 4
	switchOffBit       = 0x08 // switch off bit of the de-registration type
	secCapAlg0Bit      = 0x80 // EA0 / IA0 bit of a UE security capability octet; the next algorithms follow
	payloadN1SMInfo    = 1    // payload container type: N1 SM information
)

// NAS security.
const (
	shtMask                = 0x0f // security header type: the low half of the second octet of a 5GMM message
	shtCiphered            = 2    // security header type: integrity protected and ciphered
	shtIntegrityNewContext = 3    // security header type: integrity protected with new 5G NAS security context
	shtCipheredNewContext  = 4    // security header type: integrity protected and ciphered with new 5G NAS security context
	protectedHeaderLen     = 7    // octets of a security protected 5GMM header: EPD, header type, MAC, sequence number
	sqnMask                = 0xff // NAS sequence number: the low octet of a NAS COUNT
	sqnRange               = 256  // values of the NAS sequence number before it wraps
	bearer3GPP             = 1    // NAS connection identifier (BEARER) of 3GPP access, used when none is set
	nasAlgMask             = 0x7  // type of algorithm in the selected NAS security algorithms
	cipheringAlgShift      = 4    // position of the ciphering algorithm in the selected NAS security algorithms
	randAUTNLen            = 16   // octets of RAND and of AUTN
	amfSeparationBit       = 0x80 // separation bit of the AMF field of AUTN (set for 5G authentication)
)

// PDU sessions.
const (
	pduTypeMask   = 0x7 // PDU session type in the first octet of a PDU address
	pduTypeIPv4   = 1   // PDU session type IPv4
	pduTypeIPv6   = 2   // PDU session type IPv6
	pduTypeIPv4v6 = 3   // PDU session type IPv4v6
	maxPTI        = 254 // largest PTI the UE allocates (0 and 255 are reserved)
	defaultQFI    = 1   // QoS flow used when the setup request lists none
)

// gtpuPort is the GTP-U port (TS 29.281 §4.4.2).
const gtpuPort = 2152
