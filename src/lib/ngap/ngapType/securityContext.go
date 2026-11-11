// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SecurityContext represents ASN.1 SecurityContext.
type SecurityContext struct {
	NextHopChainingCount NextHopChainingCount
	NextHopNH            SecurityKey
	HasIEExtensions      bool
	IEExtensions         []SecurityContextExtIEsEntry
	UnknownExtensions    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SecurityContext) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SecurityContext", v)
}

// DecodeSecurityContext decodes a standalone APER value and reports the consumed octets.
func DecodeSecurityContext(data []byte) (SecurityContext, int, error) {
	var v SecurityContext
	err := ngapcodec.DecodeParameter("SecurityContext", data, &v)
	if ngapcodec.IsFatal(err) {
		return SecurityContext{}, 0, err
	}
	return v, len(data), err
}
