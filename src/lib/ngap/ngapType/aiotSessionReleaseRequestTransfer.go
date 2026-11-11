// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIOTSessionReleaseRequestTransfer represents ASN.1 AIOTSessionReleaseRequestTransfer.
type AIOTSessionReleaseRequestTransfer struct {
	CorrelationIdentifier AIoTCorrelationIdentifier
	Cause                 Cause
	HasIEExtensions       bool
	IEExtensions          []AIOTSessionReleaseRequestTransferExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIOTSessionReleaseRequestTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIOTSessionReleaseRequestTransfer", v)
}

// DecodeAIOTSessionReleaseRequestTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeAIOTSessionReleaseRequestTransfer(data []byte) (AIOTSessionReleaseRequestTransfer, int, error) {
	var v AIOTSessionReleaseRequestTransfer
	err := ngapcodec.DecodeParameter("AIOTSessionReleaseRequestTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIOTSessionReleaseRequestTransfer{}, 0, err
	}
	return v, len(data), err
}
