// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UserPlaneFailureIndication represents ASN.1 UserPlaneFailureIndication.
type UserPlaneFailureIndication struct {
	UserPlaneFailureType  UserPlaneFailureType
	ULNGUUPTNLInformation UPTransportLayerInformation
	DLNGUUPTNLInformation UPTransportLayerInformation
	HasIEExtensions       bool
	IEExtensions          []UserPlaneFailureIndicationExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserPlaneFailureIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserPlaneFailureIndication", v)
}

// DecodeUserPlaneFailureIndication decodes a standalone APER value and reports the consumed octets.
func DecodeUserPlaneFailureIndication(data []byte) (UserPlaneFailureIndication, int, error) {
	var v UserPlaneFailureIndication
	err := ngapcodec.DecodeParameter("UserPlaneFailureIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserPlaneFailureIndication{}, 0, err
	}
	return v, len(data), err
}
