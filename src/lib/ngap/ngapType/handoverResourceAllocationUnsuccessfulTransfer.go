// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// HandoverResourceAllocationUnsuccessfulTransfer represents ASN.1 HandoverResourceAllocationUnsuccessfulTransfer.
type HandoverResourceAllocationUnsuccessfulTransfer struct {
	Cause                     Cause
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    CriticalityDiagnostics
	HasIEExtensions           bool
	IEExtensions              []HandoverResourceAllocationUnsuccessfulTransferExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HandoverResourceAllocationUnsuccessfulTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("HandoverResourceAllocationUnsuccessfulTransfer", v)
}

// DecodeHandoverResourceAllocationUnsuccessfulTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeHandoverResourceAllocationUnsuccessfulTransfer(data []byte) (HandoverResourceAllocationUnsuccessfulTransfer, int, error) {
	var v HandoverResourceAllocationUnsuccessfulTransfer
	err := ngapcodec.DecodeParameter("HandoverResourceAllocationUnsuccessfulTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return HandoverResourceAllocationUnsuccessfulTransfer{}, 0, err
	}
	return v, len(data), err
}
