// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AdditionalCancelledlocationReportingReferenceIDList represents ASN.1 AdditionalCancelledlocationReportingReferenceIDList.
type AdditionalCancelledlocationReportingReferenceIDList struct {
	List []AdditionalCancelledlocationReportingReferenceIDItem `aper:"sizeLB:1,sizeUB:63"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AdditionalCancelledlocationReportingReferenceIDList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AdditionalCancelledlocationReportingReferenceIDList", v)
}

// DecodeAdditionalCancelledlocationReportingReferenceIDList decodes a standalone APER value and reports the consumed octets.
func DecodeAdditionalCancelledlocationReportingReferenceIDList(data []byte) (AdditionalCancelledlocationReportingReferenceIDList, int, error) {
	var v AdditionalCancelledlocationReportingReferenceIDList
	err := ngapcodec.DecodeParameter("AdditionalCancelledlocationReportingReferenceIDList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AdditionalCancelledlocationReportingReferenceIDList{}, 0, err
	}
	return v, len(data), err
}
