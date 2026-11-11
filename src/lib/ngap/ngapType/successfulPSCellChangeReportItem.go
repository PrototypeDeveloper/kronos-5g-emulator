// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SuccessfulPSCellChangeReportItem represents ASN.1 SuccessfulPSCellChangeReport-Item.
type SuccessfulPSCellChangeReportItem struct {
	SuccessfulPSCellChangeReportContainer []byte
	HasIEExtensions                       bool
	IEExtensions                          []SuccessfulPSCellChangeReportItemExtIEsEntry
	UnknownExtensions                     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SuccessfulPSCellChangeReportItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SuccessfulPSCellChangeReport-Item", v)
}

// DecodeSuccessfulPSCellChangeReportItem decodes a standalone APER value and reports the consumed octets.
func DecodeSuccessfulPSCellChangeReportItem(data []byte) (SuccessfulPSCellChangeReportItem, int, error) {
	var v SuccessfulPSCellChangeReportItem
	err := ngapcodec.DecodeParameter("SuccessfulPSCellChangeReport-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return SuccessfulPSCellChangeReportItem{}, 0, err
	}
	return v, len(data), err
}
