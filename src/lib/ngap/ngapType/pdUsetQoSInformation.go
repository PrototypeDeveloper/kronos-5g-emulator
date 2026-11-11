// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUsetQoSInformation represents ASN.1 PDUsetQoSInformation.
type PDUsetQoSInformation struct {
	HasPduSetDelayBudget                   bool
	PduSetDelayBudget                      ExtendedPacketDelayBudget
	HasPduSetErrorRate                     bool
	PduSetErrorRate                        PacketErrorRate
	HasPduSetIntegratedHandlingInformation bool
	PduSetIntegratedHandlingInformation    PDUsetQoSInformationPduSetIntegratedHandlingInformation `aper:"valueExt,valueLB:0,valueUB:1"`
	HasIEExtensions                        bool
	IEExtensions                           []PDUsetQoSInformationExtIEsEntry
	UnknownExtensions                      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUsetQoSInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUsetQoSInformation", v)
}

// DecodePDUsetQoSInformation decodes a standalone APER value and reports the consumed octets.
func DecodePDUsetQoSInformation(data []byte) (PDUsetQoSInformation, int, error) {
	var v PDUsetQoSInformation
	err := ngapcodec.DecodeParameter("PDUsetQoSInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUsetQoSInformation{}, 0, err
	}
	return v, len(data), err
}
