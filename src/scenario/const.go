// Constants of package scenario.

package scenario

// NGAP message association.
const (
	AssocNode   Assoc = iota // non-UE-associated: handled by the node block
	AssocUE                  // UE NGAP IDs are mandatory: handled per subscriber
	AssocEither              // UE NGAP IDs are optional (ErrorIndication ...): decided per message
)

// Block kinds of scenario.txt.
const (
	BlockNode  BlockKind = iota // node { ... }: the node itself
	BlockAny                    // any { ... }: per subscriber, every state
	BlockState                  // <state> { ... }: per subscriber, that state only
)

// Send modes of a signal.
const (
	ModeAuto    Mode = iota // patch every dynamic field (default)
	ModeNASKeep             // patch NGAP, keep the NAS body, re-protect it
	ModeRaw                 // send the bytes as written
)

// NAS header layout (TS 24.501).
const (
	gmmHeaderLen       = 3                           // plain 5GMM header (EPD, security header type, message type); also the shortest NAS message
	gsmHeaderLen       = 4                           // 5GSM header (EPD, PDU session ID, PTI, message type)
	secHeaderLen       = 7                           // security protected 5GMM header (EPD, security header type, MAC, sequence number)
	protectedNASMinLen = secHeaderLen + gmmHeaderLen // shortest protected NAS: the security header and a plain 5GMM header
	macLen             = 4                           // octets of the NAS message authentication code
	secHeaderTypeMask  = 0x0f                        // security header type bits of the second 5GMM octet
	secHeaderSpareMask = 0xf0                        // spare half octet beside the security header type
)

// 5GMM security header types (TS 24.501 §9.3); 0 is plain.
const (
	shtIntegrity            = 1 // integrity protected
	shtIntegrityCiphered    = 2 // integrity protected and ciphered
	shtIntegrityNew         = 3 // integrity protected with a new 5G NAS security context
	shtIntegrityCipheredNew = 4 // integrity protected and ciphered with a new 5G NAS security context
)

// payloadN1SMInfo is the payload container type of N1 SM information (TS 24.501 §9.11.3.40).
const payloadN1SMInfo = 1

// NGAP IE sizes.
const (
	plmnIDLen   = 3                         // octets of a PLMN identity
	tacLen      = 3                         // octets of a TAC
	nciBitLen   = 36                        // bits of an NR Cell Identity
	ipv4AddrLen = 4                         // octets of an IPv4 transport layer address
	ipv6AddrLen = 16                        // octets of an IPv6 transport layer address
	dualAddrLen = ipv4AddrLen + ipv6AddrLen // octets of an IPv4 and an IPv6 address together
)

// Rendering and editing of values.
const (
	pduKindUnknown      = 3       // index of "?" in pduKinds, for any other PDU kind
	maxBitsAsNumber     = 64      // longest BIT STRING shown and edited as a number
	decimalOnlyMax      = 9       // integers up to this are shown without their hex form
	maxBoundCheckedUint = 1 << 62 // larger unsigned values skip the int64 aper bound check
)

// Reading and writing signal.txt.
const (
	utf8BOM      = "\xef\xbb\xbf" // UTF-8 byte order mark, skipped as space
	utf8ContMask = 0xc0           // top two bits of a UTF-8 octet
	utf8ContBits = 0x80           // top two bits of a UTF-8 continuation octet
	newFilePerm  = 0o644          // permissions of a file WriteFileAtomic creates
)
