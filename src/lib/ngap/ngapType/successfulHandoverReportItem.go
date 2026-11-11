// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SuccessfulHandoverReportItem represents ASN.1 SuccessfulHandoverReport-Item.
type SuccessfulHandoverReportItem struct {
	SuccessfulHOReportContainer []byte
	HasIEExtensions             bool
	IEExtensions                []SuccessfulHandoverReportItemExtIEsEntry
	UnknownExtensions           []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SuccessfulHandoverReportItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SuccessfulHandoverReport-Item", v)
}

// DecodeSuccessfulHandoverReportItem decodes a standalone APER value and reports the consumed octets.
func DecodeSuccessfulHandoverReportItem(data []byte) (SuccessfulHandoverReportItem, int, error) {
	var v SuccessfulHandoverReportItem
	err := ngapcodec.DecodeParameter("SuccessfulHandoverReport-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return SuccessfulHandoverReportItem{}, 0, err
	}
	return v, len(data), err
}
