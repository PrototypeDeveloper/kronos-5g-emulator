// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIOTSessionReleaseCompleteTransfer represents ASN.1 AIOTSessionReleaseCompleteTransfer.
type AIOTSessionReleaseCompleteTransfer struct {
	CorrelationIdentifier     AIoTCorrelationIdentifier
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    CriticalityDiagnostics
	HasIEExtensions           bool
	IEExtensions              []AIOTSessionReleaseCompleteTransferExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIOTSessionReleaseCompleteTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIOTSessionReleaseCompleteTransfer", v)
}

// DecodeAIOTSessionReleaseCompleteTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeAIOTSessionReleaseCompleteTransfer(data []byte) (AIOTSessionReleaseCompleteTransfer, int, error) {
	var v AIOTSessionReleaseCompleteTransfer
	err := ngapcodec.DecodeParameter("AIOTSessionReleaseCompleteTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIOTSessionReleaseCompleteTransfer{}, 0, err
	}
	return v, len(data), err
}
