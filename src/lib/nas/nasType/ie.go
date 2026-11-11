package nasType

// IE is a low-level wire representation used by nascodec. NF callers use
// named message fields and their Has flags instead. Value contains the IE value only,
// without IEI or length. For a half-octet TV, ID is the high nibble (e.g. 0x90)
// and Value is one byte containing the low nibble. Message-specific schemas
// determine the wire format; callers cannot choose an inconsistent format.
// Decoded byte slices are owned by the message.
type IE struct {
	ID    uint8
	Value []byte
}
