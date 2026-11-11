// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DLPDUSetInformationMarkingSupportIndication represents ASN.1 DLPDUSetInformationMarkingSupportIndication.
const (
	DLPDUSetInformationMarkingSupportIndicationPresentTrue Enumerated = 0
)

type DLPDUSetInformationMarkingSupportIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DLPDUSetInformationMarkingSupportIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DLPDUSetInformationMarkingSupportIndication", v)
}

// DecodeDLPDUSetInformationMarkingSupportIndication decodes a standalone APER value and reports the consumed octets.
func DecodeDLPDUSetInformationMarkingSupportIndication(data []byte) (DLPDUSetInformationMarkingSupportIndication, int, error) {
	var v DLPDUSetInformationMarkingSupportIndication
	err := ngapcodec.DecodeParameter("DLPDUSetInformationMarkingSupportIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return DLPDUSetInformationMarkingSupportIndication{}, 0, err
	}
	return v, len(data), err
}
