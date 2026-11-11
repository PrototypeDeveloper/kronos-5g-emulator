// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ServedGUAMIItem represents ASN.1 ServedGUAMIItem.
type ServedGUAMIItem struct {
	GUAMI             GUAMI
	HasBackupAMFName  bool
	BackupAMFName     AMFName
	HasIEExtensions   bool
	IEExtensions      []ServedGUAMIItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ServedGUAMIItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ServedGUAMIItem", v)
}

// DecodeServedGUAMIItem decodes a standalone APER value and reports the consumed octets.
func DecodeServedGUAMIItem(data []byte) (ServedGUAMIItem, int, error) {
	var v ServedGUAMIItem
	err := ngapcodec.DecodeParameter("ServedGUAMIItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return ServedGUAMIItem{}, 0, err
	}
	return v, len(data), err
}
