// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TSCTrafficCharacteristics represents ASN.1 TSCTrafficCharacteristics.
type TSCTrafficCharacteristics struct {
	HasTSCAssistanceInformationDL bool
	TSCAssistanceInformationDL    TSCAssistanceInformation
	HasTSCAssistanceInformationUL bool
	TSCAssistanceInformationUL    TSCAssistanceInformation
	HasIEExtensions               bool
	IEExtensions                  []TSCTrafficCharacteristicsExtIEsEntry
	UnknownExtensions             []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TSCTrafficCharacteristics) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TSCTrafficCharacteristics", v)
}

// DecodeTSCTrafficCharacteristics decodes a standalone APER value and reports the consumed octets.
func DecodeTSCTrafficCharacteristics(data []byte) (TSCTrafficCharacteristics, int, error) {
	var v TSCTrafficCharacteristics
	err := ngapcodec.DecodeParameter("TSCTrafficCharacteristics", data, &v)
	if ngapcodec.IsFatal(err) {
		return TSCTrafficCharacteristics{}, 0, err
	}
	return v, len(data), err
}
