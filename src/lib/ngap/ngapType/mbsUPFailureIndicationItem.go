// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSUPFailureIndicationItem represents ASN.1 MBS-UP-FailureIndicationItem.
type MBSUPFailureIndicationItem struct {
	MBSAreaSessionID       MBSAreaSessionID
	MBSUPFailureIndication MBSUPFailureIndication
	HasIEExtensions        bool
	IEExtensions           []MBSUPFailureIndicationItemExtIEsEntry
	UnknownExtensions      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSUPFailureIndicationItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-UP-FailureIndicationItem", v)
}

// DecodeMBSUPFailureIndicationItem decodes a standalone APER value and reports the consumed octets.
func DecodeMBSUPFailureIndicationItem(data []byte) (MBSUPFailureIndicationItem, int, error) {
	var v MBSUPFailureIndicationItem
	err := ngapcodec.DecodeParameter("MBS-UP-FailureIndicationItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSUPFailureIndicationItem{}, 0, err
	}
	return v, len(data), err
}
