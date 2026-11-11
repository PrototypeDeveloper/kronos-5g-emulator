// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UERLFReportContainer represents ASN.1 UERLFReportContainer.
type UERLFReportContainerPresent uint32

const (
	UERLFReportContainerPresentNothing UERLFReportContainerPresent = iota
	UERLFReportContainerPresentNR
	UERLFReportContainerPresentLTE
	UERLFReportContainerPresentChoiceExtensions
)

type UERLFReportContainer struct {
	Present          UERLFReportContainerPresent
	NR               NRUERLFReportContainer
	LTE              LTEUERLFReportContainer
	ChoiceExtensions UERLFReportContainerExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UERLFReportContainer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UERLFReportContainer", v)
}

// DecodeUERLFReportContainer decodes a standalone APER value and reports the consumed octets.
func DecodeUERLFReportContainer(data []byte) (UERLFReportContainer, int, error) {
	var v UERLFReportContainer
	err := ngapcodec.DecodeParameter("UERLFReportContainer", data, &v)
	if ngapcodec.IsFatal(err) {
		return UERLFReportContainer{}, 0, err
	}
	return v, len(data), err
}
