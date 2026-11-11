// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntersystemCellStateIndication represents ASN.1 IntersystemCellStateIndication.
type IntersystemCellStateIndication struct {
	NotificationCellList NotificationCellList
	HasIEExtensions      bool
	IEExtensions         []IntersystemCellStateIndicationExtIEsEntry
	UnknownExtensions    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemCellStateIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemCellStateIndication", v)
}

// DecodeIntersystemCellStateIndication decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemCellStateIndication(data []byte) (IntersystemCellStateIndication, int, error) {
	var v IntersystemCellStateIndication
	err := ngapcodec.DecodeParameter("IntersystemCellStateIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemCellStateIndication{}, 0, err
	}
	return v, len(data), err
}
