// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRPagingLongEDRXInformationForRRCINACTIVE represents ASN.1 NR-Paging-Long-eDRX-Information-for-RRC-INACTIVE.
type NRPagingLongEDRXInformationForRRCINACTIVE struct {
	NRPagingLongEDRXCycleForRRCINACTIVE NRPagingLongEDRXCycleForRRCINACTIVE
	NRPagingTimeWindowForRRCINACTIVE    NRPagingTimeWindowForRRCINACTIVE
	HasIEExtensions                     bool
	IEExtensions                        []NRPagingLongEDRXInformationForRRCINACTIVEExtIEsEntry
	UnknownExtensions                   []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRPagingLongEDRXInformationForRRCINACTIVE) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NR-Paging-Long-eDRX-Information-for-RRC-INACTIVE", v)
}

// DecodeNRPagingLongEDRXInformationForRRCINACTIVE decodes a standalone APER value and reports the consumed octets.
func DecodeNRPagingLongEDRXInformationForRRCINACTIVE(data []byte) (NRPagingLongEDRXInformationForRRCINACTIVE, int, error) {
	var v NRPagingLongEDRXInformationForRRCINACTIVE
	err := ngapcodec.DecodeParameter("NR-Paging-Long-eDRX-Information-for-RRC-INACTIVE", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRPagingLongEDRXInformationForRRCINACTIVE{}, 0, err
	}
	return v, len(data), err
}
