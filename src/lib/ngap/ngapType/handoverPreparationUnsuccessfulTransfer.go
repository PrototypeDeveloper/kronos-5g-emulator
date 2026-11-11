// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// HandoverPreparationUnsuccessfulTransfer represents ASN.1 HandoverPreparationUnsuccessfulTransfer.
type HandoverPreparationUnsuccessfulTransfer struct {
	Cause             Cause
	HasIEExtensions   bool
	IEExtensions      []HandoverPreparationUnsuccessfulTransferExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HandoverPreparationUnsuccessfulTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("HandoverPreparationUnsuccessfulTransfer", v)
}

// DecodeHandoverPreparationUnsuccessfulTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeHandoverPreparationUnsuccessfulTransfer(data []byte) (HandoverPreparationUnsuccessfulTransfer, int, error) {
	var v HandoverPreparationUnsuccessfulTransfer
	err := ngapcodec.DecodeParameter("HandoverPreparationUnsuccessfulTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return HandoverPreparationUnsuccessfulTransfer{}, 0, err
	}
	return v, len(data), err
}
