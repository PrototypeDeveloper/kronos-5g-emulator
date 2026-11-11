// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRPagingTimeWindow represents ASN.1 NR-Paging-Time-Window.
const (
	NRPagingTimeWindowPresentS1  Enumerated = 0
	NRPagingTimeWindowPresentS2  Enumerated = 1
	NRPagingTimeWindowPresentS3  Enumerated = 2
	NRPagingTimeWindowPresentS4  Enumerated = 3
	NRPagingTimeWindowPresentS5  Enumerated = 4
	NRPagingTimeWindowPresentS6  Enumerated = 5
	NRPagingTimeWindowPresentS7  Enumerated = 6
	NRPagingTimeWindowPresentS8  Enumerated = 7
	NRPagingTimeWindowPresentS9  Enumerated = 8
	NRPagingTimeWindowPresentS10 Enumerated = 9
	NRPagingTimeWindowPresentS11 Enumerated = 10
	NRPagingTimeWindowPresentS12 Enumerated = 11
	NRPagingTimeWindowPresentS13 Enumerated = 12
	NRPagingTimeWindowPresentS14 Enumerated = 13
	NRPagingTimeWindowPresentS15 Enumerated = 14
	NRPagingTimeWindowPresentS16 Enumerated = 15
	NRPagingTimeWindowPresentS17 Enumerated = 16
	NRPagingTimeWindowPresentS18 Enumerated = 17
	NRPagingTimeWindowPresentS19 Enumerated = 18
	NRPagingTimeWindowPresentS20 Enumerated = 19
	NRPagingTimeWindowPresentS21 Enumerated = 20
	NRPagingTimeWindowPresentS22 Enumerated = 21
	NRPagingTimeWindowPresentS23 Enumerated = 22
	NRPagingTimeWindowPresentS24 Enumerated = 23
	NRPagingTimeWindowPresentS25 Enumerated = 24
	NRPagingTimeWindowPresentS26 Enumerated = 25
	NRPagingTimeWindowPresentS27 Enumerated = 26
	NRPagingTimeWindowPresentS28 Enumerated = 27
	NRPagingTimeWindowPresentS29 Enumerated = 28
	NRPagingTimeWindowPresentS30 Enumerated = 29
	NRPagingTimeWindowPresentS31 Enumerated = 30
	NRPagingTimeWindowPresentS32 Enumerated = 31
)

type NRPagingTimeWindow struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:15"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRPagingTimeWindow) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NR-Paging-Time-Window", v)
}

// DecodeNRPagingTimeWindow decodes a standalone APER value and reports the consumed octets.
func DecodeNRPagingTimeWindow(data []byte) (NRPagingTimeWindow, int, error) {
	var v NRPagingTimeWindow
	err := ngapcodec.DecodeParameter("NR-Paging-Time-Window", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRPagingTimeWindow{}, 0, err
	}
	return v, len(data), err
}
