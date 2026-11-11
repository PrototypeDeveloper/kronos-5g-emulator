// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRPagingLongEDRXCycleForRRCINACTIVE represents ASN.1 NR-Paging-Long-eDRX-Cycle-for-RRC-INACTIVE.
const (
	NRPagingLongEDRXCycleForRRCINACTIVEPresentHf2    Enumerated = 0
	NRPagingLongEDRXCycleForRRCINACTIVEPresentHf4    Enumerated = 1
	NRPagingLongEDRXCycleForRRCINACTIVEPresentHf8    Enumerated = 2
	NRPagingLongEDRXCycleForRRCINACTIVEPresentHf16   Enumerated = 3
	NRPagingLongEDRXCycleForRRCINACTIVEPresentHf32   Enumerated = 4
	NRPagingLongEDRXCycleForRRCINACTIVEPresentHf64   Enumerated = 5
	NRPagingLongEDRXCycleForRRCINACTIVEPresentHf128  Enumerated = 6
	NRPagingLongEDRXCycleForRRCINACTIVEPresentHf256  Enumerated = 7
	NRPagingLongEDRXCycleForRRCINACTIVEPresentHf512  Enumerated = 8
	NRPagingLongEDRXCycleForRRCINACTIVEPresentHf1024 Enumerated = 9
)

type NRPagingLongEDRXCycleForRRCINACTIVE struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:9"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRPagingLongEDRXCycleForRRCINACTIVE) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NR-Paging-Long-eDRX-Cycle-for-RRC-INACTIVE", v)
}

// DecodeNRPagingLongEDRXCycleForRRCINACTIVE decodes a standalone APER value and reports the consumed octets.
func DecodeNRPagingLongEDRXCycleForRRCINACTIVE(data []byte) (NRPagingLongEDRXCycleForRRCINACTIVE, int, error) {
	var v NRPagingLongEDRXCycleForRRCINACTIVE
	err := ngapcodec.DecodeParameter("NR-Paging-Long-eDRX-Cycle-for-RRC-INACTIVE", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRPagingLongEDRXCycleForRRCINACTIVE{}, 0, err
	}
	return v, len(data), err
}
