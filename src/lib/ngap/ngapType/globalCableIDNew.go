// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GlobalCableIDNew represents ASN.1 GlobalCable-ID-new.
type GlobalCableIDNew struct {
	GlobalCableID     GlobalCableID
	TAI               TAI
	HasIEExtensions   bool
	IEExtensions      []GlobalCableIDNewExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalCableIDNew) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GlobalCable-ID-new", v)
}

// DecodeGlobalCableIDNew decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalCableIDNew(data []byte) (GlobalCableIDNew, int, error) {
	var v GlobalCableIDNew
	err := ngapcodec.DecodeParameter("GlobalCable-ID-new", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalCableIDNew{}, 0, err
	}
	return v, len(data), err
}
