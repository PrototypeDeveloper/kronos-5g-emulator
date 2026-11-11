// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RedundantPDUSessionInformation represents ASN.1 RedundantPDUSessionInformation.
type RedundantPDUSessionInformation struct {
	RSN               RSN
	HasIEExtensions   bool
	IEExtensions      []RedundantPDUSessionInformationExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RedundantPDUSessionInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RedundantPDUSessionInformation", v)
}

// DecodeRedundantPDUSessionInformation decodes a standalone APER value and reports the consumed octets.
func DecodeRedundantPDUSessionInformation(data []byte) (RedundantPDUSessionInformation, int, error) {
	var v RedundantPDUSessionInformation
	err := ngapcodec.DecodeParameter("RedundantPDUSessionInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return RedundantPDUSessionInformation{}, 0, err
	}
	return v, len(data), err
}
