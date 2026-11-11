// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QMCDeactivation represents ASN.1 QMCDeactivation.
type QMCDeactivation struct {
	QoEReferenceList  QoEReferenceList
	HasIEExtensions   bool
	IEExtensions      []QMCDeactivationExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QMCDeactivation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QMCDeactivation", v)
}

// DecodeQMCDeactivation decodes a standalone APER value and reports the consumed octets.
func DecodeQMCDeactivation(data []byte) (QMCDeactivation, int, error) {
	var v QMCDeactivation
	err := ngapcodec.DecodeParameter("QMCDeactivation", data, &v)
	if ngapcodec.IsFatal(err) {
		return QMCDeactivation{}, 0, err
	}
	return v, len(data), err
}
