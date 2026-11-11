// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BroadcastTransportFailureTransfer represents ASN.1 BroadcastTransportFailureTransfer.
type BroadcastTransportFailureTransfer struct {
	MBSSessionID              MBSSessionID
	Cause                     Cause
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    CriticalityDiagnostics
	HasIEExtensions           bool
	IEExtensions              []BroadcastTransportFailureTransferExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BroadcastTransportFailureTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BroadcastTransportFailureTransfer", v)
}

// DecodeBroadcastTransportFailureTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeBroadcastTransportFailureTransfer(data []byte) (BroadcastTransportFailureTransfer, int, error) {
	var v BroadcastTransportFailureTransfer
	err := ngapcodec.DecodeParameter("BroadcastTransportFailureTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return BroadcastTransportFailureTransfer{}, 0, err
	}
	return v, len(data), err
}
