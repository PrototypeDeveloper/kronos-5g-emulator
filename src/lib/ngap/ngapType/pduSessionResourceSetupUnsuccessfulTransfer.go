// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceSetupUnsuccessfulTransfer represents ASN.1 PDUSessionResourceSetupUnsuccessfulTransfer.
type PDUSessionResourceSetupUnsuccessfulTransfer struct {
	Cause                     Cause
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    CriticalityDiagnostics
	HasIEExtensions           bool
	IEExtensions              []PDUSessionResourceSetupUnsuccessfulTransferExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupUnsuccessfulTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupUnsuccessfulTransfer", v)
}

// DecodePDUSessionResourceSetupUnsuccessfulTransfer decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupUnsuccessfulTransfer(data []byte) (PDUSessionResourceSetupUnsuccessfulTransfer, int, error) {
	var v PDUSessionResourceSetupUnsuccessfulTransfer
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupUnsuccessfulTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupUnsuccessfulTransfer{}, 0, err
	}
	return v, len(data), err
}
