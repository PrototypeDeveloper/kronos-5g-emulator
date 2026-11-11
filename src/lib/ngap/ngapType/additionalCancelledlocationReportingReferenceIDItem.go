// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AdditionalCancelledlocationReportingReferenceIDItem represents ASN.1 AdditionalCancelledlocationReportingReferenceIDItem.
type AdditionalCancelledlocationReportingReferenceIDItem struct {
	LocationReportingReferenceIDToBeCancelled LocationReportingReferenceID
	HasIEExtensions                           bool
	IEExtensions                              []AdditionalCancelledlocationReportingReferenceIDItemExtIEsEntry
	UnknownExtensions                         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AdditionalCancelledlocationReportingReferenceIDItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AdditionalCancelledlocationReportingReferenceIDItem", v)
}

// DecodeAdditionalCancelledlocationReportingReferenceIDItem decodes a standalone APER value and reports the consumed octets.
func DecodeAdditionalCancelledlocationReportingReferenceIDItem(data []byte) (AdditionalCancelledlocationReportingReferenceIDItem, int, error) {
	var v AdditionalCancelledlocationReportingReferenceIDItem
	err := ngapcodec.DecodeParameter("AdditionalCancelledlocationReportingReferenceIDItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return AdditionalCancelledlocationReportingReferenceIDItem{}, 0, err
	}
	return v, len(data), err
}
