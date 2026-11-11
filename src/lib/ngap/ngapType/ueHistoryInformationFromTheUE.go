// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEHistoryInformationFromTheUE represents ASN.1 UEHistoryInformationFromTheUE.
type UEHistoryInformationFromTheUEPresent uint32

const (
	UEHistoryInformationFromTheUEPresentNothing UEHistoryInformationFromTheUEPresent = iota
	UEHistoryInformationFromTheUEPresentNR
	UEHistoryInformationFromTheUEPresentChoiceExtensions
)

type UEHistoryInformationFromTheUE struct {
	Present          UEHistoryInformationFromTheUEPresent
	NR               NRMobilityHistoryReport
	ChoiceExtensions UEHistoryInformationFromTheUEExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEHistoryInformationFromTheUE) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEHistoryInformationFromTheUE", v)
}

// DecodeUEHistoryInformationFromTheUE decodes a standalone APER value and reports the consumed octets.
func DecodeUEHistoryInformationFromTheUE(data []byte) (UEHistoryInformationFromTheUE, int, error) {
	var v UEHistoryInformationFromTheUE
	err := ngapcodec.DecodeParameter("UEHistoryInformationFromTheUE", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEHistoryInformationFromTheUE{}, 0, err
	}
	return v, len(data), err
}
