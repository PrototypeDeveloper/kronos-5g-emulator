// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QoEReferenceList represents ASN.1 QoEReferenceList.
type QoEReferenceList struct {
	List []QoEReference `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QoEReferenceList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QoEReferenceList", v)
}

// DecodeQoEReferenceList decodes a standalone APER value and reports the consumed octets.
func DecodeQoEReferenceList(data []byte) (QoEReferenceList, int, error) {
	var v QoEReferenceList
	err := ngapcodec.DecodeParameter("QoEReferenceList", data, &v)
	if ngapcodec.IsFatal(err) {
		return QoEReferenceList{}, 0, err
	}
	return v, len(data), err
}
