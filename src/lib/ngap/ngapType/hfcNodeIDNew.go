// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// HFCNodeIDNew represents ASN.1 HFCNode-ID-new.
type HFCNodeIDNew struct {
	HFCNodeID         HFCNodeID
	TAI               TAI
	HasIEExtensions   bool
	IEExtensions      []HFCNodeIDNewExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HFCNodeIDNew) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("HFCNode-ID-new", v) }

// DecodeHFCNodeIDNew decodes a standalone APER value and reports the consumed octets.
func DecodeHFCNodeIDNew(data []byte) (HFCNodeIDNew, int, error) {
	var v HFCNodeIDNew
	err := ngapcodec.DecodeParameter("HFCNode-ID-new", data, &v)
	if ngapcodec.IsFatal(err) {
		return HFCNodeIDNew{}, 0, err
	}
	return v, len(data), err
}
