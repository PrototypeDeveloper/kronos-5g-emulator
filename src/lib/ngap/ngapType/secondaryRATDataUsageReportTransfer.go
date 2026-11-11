// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SecondaryRATDataUsageReportTransfer represents ASN.1 SecondaryRATDataUsageReportTransfer.
type SecondaryRATDataUsageReportTransfer struct {
	HasSecondaryRATUsageInformation bool
	SecondaryRATUsageInformation    SecondaryRATUsageInformation
	HasIEExtensions                 bool
	IEExtensions                    []SecondaryRATDataUsageReportTransferExtIEsEntry
	UnknownExtensions               []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SecondaryRATDataUsageReportTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SecondaryRATDataUsageReportTransfer", v)
}

// DecodeSecondaryRATDataUsageReportTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeSecondaryRATDataUsageReportTransfer(data []byte) (SecondaryRATDataUsageReportTransfer, int, error) {
	var v SecondaryRATDataUsageReportTransfer
	err := ngapcodec.DecodeParameter("SecondaryRATDataUsageReportTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return SecondaryRATDataUsageReportTransfer{}, 0, err
	}
	return v, len(data), err
}
