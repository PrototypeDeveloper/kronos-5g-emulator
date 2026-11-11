// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionFSAID represents ASN.1 MBS-SessionFSAID.
type MBSSessionFSAID struct {
	Value []byte `aper:"sizeLB:3,sizeUB:3"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionFSAID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-SessionFSAID", v)
}

// DecodeMBSSessionFSAID decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionFSAID(data []byte) (MBSSessionFSAID, int, error) {
	var v MBSSessionFSAID
	err := ngapcodec.DecodeParameter("MBS-SessionFSAID", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionFSAID{}, 0, err
	}
	return v, len(data), err
}
