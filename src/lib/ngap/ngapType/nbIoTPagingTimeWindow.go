// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NBIoTPagingTimeWindow represents ASN.1 NB-IoT-Paging-TimeWindow.
const (
	NBIoTPagingTimeWindowPresentS1  Enumerated = 0
	NBIoTPagingTimeWindowPresentS2  Enumerated = 1
	NBIoTPagingTimeWindowPresentS3  Enumerated = 2
	NBIoTPagingTimeWindowPresentS4  Enumerated = 3
	NBIoTPagingTimeWindowPresentS5  Enumerated = 4
	NBIoTPagingTimeWindowPresentS6  Enumerated = 5
	NBIoTPagingTimeWindowPresentS7  Enumerated = 6
	NBIoTPagingTimeWindowPresentS8  Enumerated = 7
	NBIoTPagingTimeWindowPresentS9  Enumerated = 8
	NBIoTPagingTimeWindowPresentS10 Enumerated = 9
	NBIoTPagingTimeWindowPresentS11 Enumerated = 10
	NBIoTPagingTimeWindowPresentS12 Enumerated = 11
	NBIoTPagingTimeWindowPresentS13 Enumerated = 12
	NBIoTPagingTimeWindowPresentS14 Enumerated = 13
	NBIoTPagingTimeWindowPresentS15 Enumerated = 14
	NBIoTPagingTimeWindowPresentS16 Enumerated = 15
)

type NBIoTPagingTimeWindow struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:15"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NBIoTPagingTimeWindow) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NB-IoT-Paging-TimeWindow", v)
}

// DecodeNBIoTPagingTimeWindow decodes a standalone APER value and reports the consumed octets.
func DecodeNBIoTPagingTimeWindow(data []byte) (NBIoTPagingTimeWindow, int, error) {
	var v NBIoTPagingTimeWindow
	err := ngapcodec.DecodeParameter("NB-IoT-Paging-TimeWindow", data, &v)
	if ngapcodec.IsFatal(err) {
		return NBIoTPagingTimeWindow{}, 0, err
	}
	return v, len(data), err
}
