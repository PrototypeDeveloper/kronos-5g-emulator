// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TargeteNBID represents ASN.1 TargeteNB-ID.
type TargeteNBID struct {
	GlobalENBID       GlobalNgENBID
	SelectedEPSTAI    EPSTAI
	HasIEExtensions   bool
	IEExtensions      []TargeteNBIDExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargeteNBID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TargeteNB-ID", v) }

// DecodeTargeteNBID decodes a standalone APER value and reports the consumed octets.
func DecodeTargeteNBID(data []byte) (TargeteNBID, int, error) {
	var v TargeteNBID
	err := ngapcodec.DecodeParameter("TargeteNB-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargeteNBID{}, 0, err
	}
	return v, len(data), err
}
