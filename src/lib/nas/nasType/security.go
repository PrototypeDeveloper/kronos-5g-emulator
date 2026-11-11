package nasType

// SecurityHeaderType is the low nibble of octet 2 in a 5GMM NAS PDU.
type SecurityHeaderType uint8

const (
	SecurityHeaderPlain SecurityHeaderType = iota
	SecurityHeaderIntegrityProtected
	SecurityHeaderIntegrityProtectedAndCiphered
	SecurityHeaderIntegrityProtectedWithNewContext
	SecurityHeaderIntegrityProtectedAndCipheredWithNewContext
)

type Direction uint8

const (
	Uplink Direction = iota
	Downlink
)

// Count is the next unused NAS COUNT, not the last transmitted/received COUNT.
// 0..0xffffff are usable; 0x1000000 is the exhausted sentinel.
// Keep separate contexts at each endpoint and for each access connection.
type Count uint32

const MaxCount Count = 0xffffff

func (c Count) SequenceNumber() uint8 { return uint8(c) }
func (c Count) Overflow() uint16      { return uint16(c >> 8) }
