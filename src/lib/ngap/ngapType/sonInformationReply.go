// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SONInformationReply represents ASN.1 SONInformationReply.
type SONInformationReply struct {
	HasXnTNLConfigurationInfo bool
	XnTNLConfigurationInfo    XnTNLConfigurationInfo
	HasIEExtensions           bool
	IEExtensions              []SONInformationReplyExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SONInformationReply) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SONInformationReply", v)
}

// DecodeSONInformationReply decodes a standalone APER value and reports the consumed octets.
func DecodeSONInformationReply(data []byte) (SONInformationReply, int, error) {
	var v SONInformationReply
	err := ngapcodec.DecodeParameter("SONInformationReply", data, &v)
	if ngapcodec.IsFatal(err) {
		return SONInformationReply{}, 0, err
	}
	return v, len(data), err
}
