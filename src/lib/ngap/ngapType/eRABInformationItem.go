// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ERABInformationItem represents ASN.1 E-RABInformationItem.
type ERABInformationItem struct {
	ERABID            ERABID
	HasDLForwarding   bool
	DLForwarding      DLForwarding
	HasIEExtensions   bool
	IEExtensions      []ERABInformationItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ERABInformationItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("E-RABInformationItem", v)
}

// DecodeERABInformationItem decodes a standalone APER value and reports the consumed octets.
func DecodeERABInformationItem(data []byte) (ERABInformationItem, int, error) {
	var v ERABInformationItem
	err := ngapcodec.DecodeParameter("E-RABInformationItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return ERABInformationItem{}, 0, err
	}
	return v, len(data), err
}
