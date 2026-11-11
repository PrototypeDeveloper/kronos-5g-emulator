// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EPSTAI represents ASN.1 EPS-TAI.
type EPSTAI struct {
	PLMNIdentity      PLMNIdentity
	EPSTAC            EPSTAC
	HasIEExtensions   bool
	IEExtensions      []EPSTAIExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EPSTAI) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("EPS-TAI", v) }

// DecodeEPSTAI decodes a standalone APER value and reports the consumed octets.
func DecodeEPSTAI(data []byte) (EPSTAI, int, error) {
	var v EPSTAI
	err := ngapcodec.DecodeParameter("EPS-TAI", data, &v)
	if ngapcodec.IsFatal(err) {
		return EPSTAI{}, 0, err
	}
	return v, len(data), err
}
