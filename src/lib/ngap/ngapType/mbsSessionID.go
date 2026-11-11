// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionID represents ASN.1 MBS-SessionID.
type MBSSessionID struct {
	TMGI              TMGI
	HasNID            bool
	NID               NID
	HasIEExtensions   bool
	IEExtensions      []MBSSessionIDExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("MBS-SessionID", v) }

// DecodeMBSSessionID decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionID(data []byte) (MBSSessionID, int, error) {
	var v MBSSessionID
	err := ngapcodec.DecodeParameter("MBS-SessionID", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionID{}, 0, err
	}
	return v, len(data), err
}
