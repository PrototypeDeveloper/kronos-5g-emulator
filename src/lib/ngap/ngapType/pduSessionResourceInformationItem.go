// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceInformationItem represents ASN.1 PDUSessionResourceInformationItem.
type PDUSessionResourceInformationItem struct {
	PDUSessionID                 PDUSessionID
	QosFlowInformationList       QosFlowInformationList
	HasDRBsToQosFlowsMappingList bool
	DRBsToQosFlowsMappingList    DRBsToQosFlowsMappingList
	HasIEExtensions              bool
	IEExtensions                 []PDUSessionResourceInformationItemExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceInformationItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceInformationItem", v)
}

// DecodePDUSessionResourceInformationItem decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceInformationItem(data []byte) (PDUSessionResourceInformationItem, int, error) {
	var v PDUSessionResourceInformationItem
	err := ngapcodec.DecodeParameter("PDUSessionResourceInformationItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceInformationItem{}, 0, err
	}
	return v, len(data), err
}
