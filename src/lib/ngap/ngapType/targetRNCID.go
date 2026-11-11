// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TargetRNCID represents ASN.1 TargetRNC-ID.
type TargetRNCID struct {
	LAI               LAI
	RNCID             RNCID
	HasExtendedRNCID  bool
	ExtendedRNCID     ExtendedRNCID
	HasIEExtensions   bool
	IEExtensions      []TargetRNCIDExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetRNCID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TargetRNC-ID", v) }

// DecodeTargetRNCID decodes a standalone APER value and reports the consumed octets.
func DecodeTargetRNCID(data []byte) (TargetRNCID, int, error) {
	var v TargetRNCID
	err := ngapcodec.DecodeParameter("TargetRNC-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetRNCID{}, 0, err
	}
	return v, len(data), err
}
