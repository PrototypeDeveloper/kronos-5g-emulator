// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GUAMI represents ASN.1 GUAMI.
type GUAMI struct {
	PLMNIdentity      PLMNIdentity
	AMFRegionID       AMFRegionID
	AMFSetID          AMFSetID
	AMFPointer        AMFPointer
	HasIEExtensions   bool
	IEExtensions      []GUAMIExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GUAMI) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("GUAMI", v) }

// DecodeGUAMI decodes a standalone APER value and reports the consumed octets.
func DecodeGUAMI(data []byte) (GUAMI, int, error) {
	var v GUAMI
	err := ngapcodec.DecodeParameter("GUAMI", data, &v)
	if ngapcodec.IsFatal(err) {
		return GUAMI{}, 0, err
	}
	return v, len(data), err
}
