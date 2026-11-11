// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TSCAssistanceInformation represents ASN.1 TSCAssistanceInformation.
type TSCAssistanceInformation struct {
	Periodicity         Periodicity
	HasBurstArrivalTime bool
	BurstArrivalTime    BurstArrivalTime
	HasIEExtensions     bool
	IEExtensions        []TSCAssistanceInformationExtIEsEntry
	UnknownExtensions   []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TSCAssistanceInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TSCAssistanceInformation", v)
}

// DecodeTSCAssistanceInformation decodes a standalone APER value and reports the consumed octets.
func DecodeTSCAssistanceInformation(data []byte) (TSCAssistanceInformation, int, error) {
	var v TSCAssistanceInformation
	err := ngapcodec.DecodeParameter("TSCAssistanceInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return TSCAssistanceInformation{}, 0, err
	}
	return v, len(data), err
}
