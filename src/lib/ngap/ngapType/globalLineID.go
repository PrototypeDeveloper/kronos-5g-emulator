// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GlobalLineID represents ASN.1 GlobalLine-ID.
type GlobalLineID struct {
	GlobalLineIdentity GlobalLineIdentity
	HasLineType        bool
	LineType           LineType
	HasIEExtensions    bool
	IEExtensions       []GlobalLineIDExtIEsEntry
	UnknownExtensions  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalLineID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("GlobalLine-ID", v) }

// DecodeGlobalLineID decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalLineID(data []byte) (GlobalLineID, int, error) {
	var v GlobalLineID
	err := ngapcodec.DecodeParameter("GlobalLine-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalLineID{}, 0, err
	}
	return v, len(data), err
}
