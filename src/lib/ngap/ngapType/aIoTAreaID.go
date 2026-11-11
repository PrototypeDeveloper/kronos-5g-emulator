// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIoTAreaID represents ASN.1 AIoTAreaID.
type AIoTAreaID struct {
	PLMNidentity      PLMNIdentity
	HasNID            bool
	NID               NID
	AIoTAreaCode      []byte `aper:"sizeLB:3,sizeUB:3"`
	HasIEExtensions   bool
	IEExtensions      []AIoTAreaIDExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTAreaID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("AIoTAreaID", v) }

// DecodeAIoTAreaID decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTAreaID(data []byte) (AIoTAreaID, int, error) {
	var v AIoTAreaID
	err := ngapcodec.DecodeParameter("AIoTAreaID", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTAreaID{}, 0, err
	}
	return v, len(data), err
}
