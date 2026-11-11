// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NPNSupport represents ASN.1 NPN-Support.
type NPNSupportPresent uint32

const (
	NPNSupportPresentNothing NPNSupportPresent = iota
	NPNSupportPresentSNPN
	NPNSupportPresentChoiceExtensions
)

type NPNSupport struct {
	Present          NPNSupportPresent
	SNPN             NID
	ChoiceExtensions NPNSupportExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NPNSupport) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("NPN-Support", v) }

// DecodeNPNSupport decodes a standalone APER value and reports the consumed octets.
func DecodeNPNSupport(data []byte) (NPNSupport, int, error) {
	var v NPNSupport
	err := ngapcodec.DecodeParameter("NPN-Support", data, &v)
	if ngapcodec.IsFatal(err) {
		return NPNSupport{}, 0, err
	}
	return v, len(data), err
}
