// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// InventoryFailureTransfer represents ASN.1 InventoryFailureTransfer.
type InventoryFailureTransfer struct {
	CorrelationIdentifier     AIoTCorrelationIdentifier
	Cause                     Cause
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    CriticalityDiagnostics
	HasIEExtensions           bool
	IEExtensions              []InventoryFailureTransferExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InventoryFailureTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InventoryFailureTransfer", v)
}

// DecodeInventoryFailureTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeInventoryFailureTransfer(data []byte) (InventoryFailureTransfer, int, error) {
	var v InventoryFailureTransfer
	err := ngapcodec.DecodeParameter("InventoryFailureTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return InventoryFailureTransfer{}, 0, err
	}
	return v, len(data), err
}
