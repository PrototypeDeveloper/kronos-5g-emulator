// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionUsageReportRATType represents ASN.1 PDUSessionUsageReportRATType.
const (
	PDUSessionUsageReportRATTypePresentNr              Enumerated = 0
	PDUSessionUsageReportRATTypePresentEutra           Enumerated = 1
	PDUSessionUsageReportRATTypePresentNrUnlicensed    Enumerated = 2
	PDUSessionUsageReportRATTypePresentEUtraUnlicensed Enumerated = 3
)

type PDUSessionUsageReportRATType struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v PDUSessionUsageReportRATType) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:PDUSessionUsageReportRATType", v)
}

// DecodePDUSessionUsageReportRATType decodes a standalone APER value and reports consumed octets.
func DecodePDUSessionUsageReportRATType(data []byte) (PDUSessionUsageReportRATType, int, error) {
	var v PDUSessionUsageReportRATType
	err := ngapcodec.DecodeParameter("inline:PDUSessionUsageReportRATType", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionUsageReportRATType{}, 0, err
	}
	return v, len(data), err
}
