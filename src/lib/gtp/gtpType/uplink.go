package gtpType

import "fmt"

// Bounds shared by the N3 userspace receiver and the XDP parser. These are
// implementation limits, not limits imposed by 3GPP.
const MaxUplinkExtensionHeaders = 16
const MaxUplinkExtensionBytes = 256
const MaxNewIEFlagOctets = 8

type UnsupportedExtensionError struct{ Type uint8 }

func (e *UnsupportedExtensionError) Error() string {
	return fmt.Sprintf("unsupported comprehension-required GTP extension 0x%02x", e.Type)
}

// ValidateUplinkExtensions validates the complete chain before reporting an
// unsupported extension. No user data is forwarded on that error.
func ValidateUplinkExtensions(headers []ExtensionHeader) error {
	if len(headers) > MaxUplinkExtensionHeaders {
		return fmt.Errorf("GTP extension count exceeds limit")
	}
	total, psc := 0, false
	var unsupported *UnsupportedExtensionError
	for _, h := range headers {
		n := int(h.Header.Length) * 4
		if n < 4 || len(h.Content) != n-2 {
			return fmt.Errorf("invalid GTP extension length")
		}
		total += n
		if total > MaxUplinkExtensionBytes {
			return fmt.Errorf("GTP extension bytes exceed limit")
		}
		if h.Header.Type == ExtensionHeaderPDUSessionContainer {
			if psc {
				return fmt.Errorf("duplicate PDU Session Container")
			}
			psc = true
			if err := ValidateULPDUSessionContent(h.Content); err != nil {
				return err
			}
		} else if h.Header.Type&0x80 != 0 {
			unsupported = &UnsupportedExtensionError{h.Header.Type}
		}
	}
	if unsupported != nil {
		return unsupported
	}
	return nil
}

// ValidateULPDUSessionContent implements TS 38.415 5.5.2.2. The GTP sequence
// number and this container's optional 24-bit QFI sequence number are distinct.
// Future extension/padding bytes are skipped; all indicated known fields must
// fit. Monitoring results are accepted without advertising monitoring support.
func ValidateULPDUSessionContent(content []byte) error {
	if len(content) < 2 || content[0]>>4 != PDUSessionInformationUL {
		return fmt.Errorf("UL PDU Session Information required")
	}
	a, b := content[0], content[1]
	n := 2 + int((a>>3)&1)*24 + int((a>>2)&1)*4 + int((a>>1)&1)*4 + int(a&1)*3 + int(b>>7)*4
	if n > len(content) {
		return fmt.Errorf("truncated UL PDU Session Information")
	}
	if b&0x40 != 0 {
		first := byte(0)
		for i := 0; ; i++ {
			if i == MaxNewIEFlagOctets || n >= len(content) {
				return fmt.Errorf("invalid New IE Flags chain")
			}
			flags := content[n]
			n++
			if i == 0 {
				first = flags
			}
			if flags&0x80 == 0 {
				break
			}
		}
		n += int(first&1) + int((first>>1)&1)*2 + int((first>>2)&1)*2
		if n > len(content) {
			return fmt.Errorf("truncated new UL PDU Session fields")
		}
	}
	return nil
}
