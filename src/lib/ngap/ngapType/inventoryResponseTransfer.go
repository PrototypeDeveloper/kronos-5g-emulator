// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// InventoryResponseTransfer represents ASN.1 InventoryResponseTransfer.
type InventoryResponseTransfer struct {
	CorrelationIdentifier     AIoTCorrelationIdentifier
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    CriticalityDiagnostics
	HasIEExtensions           bool
	IEExtensions              []InventoryResponseTransferExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InventoryResponseTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InventoryResponseTransfer", v)
}

// DecodeInventoryResponseTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeInventoryResponseTransfer(data []byte) (InventoryResponseTransfer, int, error) {
	var v InventoryResponseTransfer
	err := ngapcodec.DecodeParameter("InventoryResponseTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return InventoryResponseTransfer{}, 0, err
	}
	return v, len(data), err
}
