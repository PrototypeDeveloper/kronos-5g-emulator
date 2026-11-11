// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEContextResumeResponseTransfer represents ASN.1 UEContextResumeResponseTransfer.
type UEContextResumeResponseTransfer struct {
	HasQosFlowFailedToResumeList bool
	QosFlowFailedToResumeList    QosFlowListWithCause
	HasIEExtensions              bool
	IEExtensions                 []UEContextResumeResponseTransferExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEContextResumeResponseTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEContextResumeResponseTransfer", v)
}

// DecodeUEContextResumeResponseTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeUEContextResumeResponseTransfer(data []byte) (UEContextResumeResponseTransfer, int, error) {
	var v UEContextResumeResponseTransfer
	err := ngapcodec.DecodeParameter("UEContextResumeResponseTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEContextResumeResponseTransfer{}, 0, err
	}
	return v, len(data), err
}
