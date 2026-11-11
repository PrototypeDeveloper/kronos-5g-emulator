// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NBIoTPagingEDRXCycle represents ASN.1 NB-IoT-Paging-eDRXCycle.
const (
	NBIoTPagingEDRXCyclePresentHf2    Enumerated = 0
	NBIoTPagingEDRXCyclePresentHf4    Enumerated = 1
	NBIoTPagingEDRXCyclePresentHf6    Enumerated = 2
	NBIoTPagingEDRXCyclePresentHf8    Enumerated = 3
	NBIoTPagingEDRXCyclePresentHf10   Enumerated = 4
	NBIoTPagingEDRXCyclePresentHf12   Enumerated = 5
	NBIoTPagingEDRXCyclePresentHf14   Enumerated = 6
	NBIoTPagingEDRXCyclePresentHf16   Enumerated = 7
	NBIoTPagingEDRXCyclePresentHf32   Enumerated = 8
	NBIoTPagingEDRXCyclePresentHf64   Enumerated = 9
	NBIoTPagingEDRXCyclePresentHf128  Enumerated = 10
	NBIoTPagingEDRXCyclePresentHf256  Enumerated = 11
	NBIoTPagingEDRXCyclePresentHf512  Enumerated = 12
	NBIoTPagingEDRXCyclePresentHf1024 Enumerated = 13
)

type NBIoTPagingEDRXCycle struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:13"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NBIoTPagingEDRXCycle) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NB-IoT-Paging-eDRXCycle", v)
}

// DecodeNBIoTPagingEDRXCycle decodes a standalone APER value and reports the consumed octets.
func DecodeNBIoTPagingEDRXCycle(data []byte) (NBIoTPagingEDRXCycle, int, error) {
	var v NBIoTPagingEDRXCycle
	err := ngapcodec.DecodeParameter("NB-IoT-Paging-eDRXCycle", data, &v)
	if ngapcodec.IsFatal(err) {
		return NBIoTPagingEDRXCycle{}, 0, err
	}
	return v, len(data), err
}
