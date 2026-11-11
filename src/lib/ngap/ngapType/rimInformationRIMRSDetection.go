// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RIMInformationRIMRSDetection represents ASN.1 RIMInformationRIMRSDetection.
const (
	RIMInformationRIMRSDetectionPresentRsDetected    Enumerated = 0
	RIMInformationRIMRSDetectionPresentRsDisappeared Enumerated = 1
)

type RIMInformationRIMRSDetection struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v RIMInformationRIMRSDetection) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:RIMInformationRIMRSDetection", v)
}

// DecodeRIMInformationRIMRSDetection decodes a standalone APER value and reports consumed octets.
func DecodeRIMInformationRIMRSDetection(data []byte) (RIMInformationRIMRSDetection, int, error) {
	var v RIMInformationRIMRSDetection
	err := ngapcodec.DecodeParameter("inline:RIMInformationRIMRSDetection", data, &v)
	if ngapcodec.IsFatal(err) {
		return RIMInformationRIMRSDetection{}, 0, err
	}
	return v, len(data), err
}
