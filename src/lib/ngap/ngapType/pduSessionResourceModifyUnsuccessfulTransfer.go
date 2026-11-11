// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceModifyUnsuccessfulTransfer represents ASN.1 PDUSessionResourceModifyUnsuccessfulTransfer.
type PDUSessionResourceModifyUnsuccessfulTransfer struct {
	Cause                     Cause
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    CriticalityDiagnostics
	HasIEExtensions           bool
	IEExtensions              []PDUSessionResourceModifyUnsuccessfulTransferExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyUnsuccessfulTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyUnsuccessfulTransfer", v)
}

// DecodePDUSessionResourceModifyUnsuccessfulTransfer decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyUnsuccessfulTransfer(data []byte) (PDUSessionResourceModifyUnsuccessfulTransfer, int, error) {
	var v PDUSessionResourceModifyUnsuccessfulTransfer
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyUnsuccessfulTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyUnsuccessfulTransfer{}, 0, err
	}
	return v, len(data), err
}
