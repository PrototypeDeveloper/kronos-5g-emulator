// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRPagingEDRXCycle represents ASN.1 NR-Paging-eDRX-Cycle.
const (
	NRPagingEDRXCyclePresentHfquarter Enumerated = 0
	NRPagingEDRXCyclePresentHfhalf    Enumerated = 1
	NRPagingEDRXCyclePresentHf1       Enumerated = 2
	NRPagingEDRXCyclePresentHf2       Enumerated = 3
	NRPagingEDRXCyclePresentHf4       Enumerated = 4
	NRPagingEDRXCyclePresentHf8       Enumerated = 5
	NRPagingEDRXCyclePresentHf16      Enumerated = 6
	NRPagingEDRXCyclePresentHf32      Enumerated = 7
	NRPagingEDRXCyclePresentHf64      Enumerated = 8
	NRPagingEDRXCyclePresentHf128     Enumerated = 9
	NRPagingEDRXCyclePresentHf256     Enumerated = 10
	NRPagingEDRXCyclePresentHf512     Enumerated = 11
	NRPagingEDRXCyclePresentHf1024    Enumerated = 12
)

type NRPagingEDRXCycle struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:12"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRPagingEDRXCycle) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NR-Paging-eDRX-Cycle", v)
}

// DecodeNRPagingEDRXCycle decodes a standalone APER value and reports the consumed octets.
func DecodeNRPagingEDRXCycle(data []byte) (NRPagingEDRXCycle, int, error) {
	var v NRPagingEDRXCycle
	err := ngapcodec.DecodeParameter("NR-Paging-eDRX-Cycle", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRPagingEDRXCycle{}, 0, err
	}
	return v, len(data), err
}
