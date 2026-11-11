// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EUTRAPagingTimeWindow represents ASN.1 EUTRA-Paging-Time-Window.
const (
	EUTRAPagingTimeWindowPresentS1  Enumerated = 0
	EUTRAPagingTimeWindowPresentS2  Enumerated = 1
	EUTRAPagingTimeWindowPresentS3  Enumerated = 2
	EUTRAPagingTimeWindowPresentS4  Enumerated = 3
	EUTRAPagingTimeWindowPresentS5  Enumerated = 4
	EUTRAPagingTimeWindowPresentS6  Enumerated = 5
	EUTRAPagingTimeWindowPresentS7  Enumerated = 6
	EUTRAPagingTimeWindowPresentS8  Enumerated = 7
	EUTRAPagingTimeWindowPresentS9  Enumerated = 8
	EUTRAPagingTimeWindowPresentS10 Enumerated = 9
	EUTRAPagingTimeWindowPresentS11 Enumerated = 10
	EUTRAPagingTimeWindowPresentS12 Enumerated = 11
	EUTRAPagingTimeWindowPresentS13 Enumerated = 12
	EUTRAPagingTimeWindowPresentS14 Enumerated = 13
	EUTRAPagingTimeWindowPresentS15 Enumerated = 14
	EUTRAPagingTimeWindowPresentS16 Enumerated = 15
)

type EUTRAPagingTimeWindow struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:15"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRAPagingTimeWindow) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRA-Paging-Time-Window", v)
}

// DecodeEUTRAPagingTimeWindow decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRAPagingTimeWindow(data []byte) (EUTRAPagingTimeWindow, int, error) {
	var v EUTRAPagingTimeWindow
	err := ngapcodec.DecodeParameter("EUTRA-Paging-Time-Window", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRAPagingTimeWindow{}, 0, err
	}
	return v, len(data), err
}
