// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEContextResumeRequestTransfer represents ASN.1 UEContextResumeRequestTransfer.
type UEContextResumeRequestTransfer struct {
	HasQosFlowFailedToResumeList bool
	QosFlowFailedToResumeList    QosFlowListWithCause
	HasIEExtensions              bool
	IEExtensions                 []UEContextResumeRequestTransferExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEContextResumeRequestTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEContextResumeRequestTransfer", v)
}

// DecodeUEContextResumeRequestTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeUEContextResumeRequestTransfer(data []byte) (UEContextResumeRequestTransfer, int, error) {
	var v UEContextResumeRequestTransfer
	err := ngapcodec.DecodeParameter("UEContextResumeRequestTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEContextResumeRequestTransfer{}, 0, err
	}
	return v, len(data), err
}
