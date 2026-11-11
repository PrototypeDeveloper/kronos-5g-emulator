// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// HandoverRequiredTransfer represents ASN.1 HandoverRequiredTransfer.
type HandoverRequiredTransfer struct {
	HasDirectForwardingPathAvailability bool
	DirectForwardingPathAvailability    DirectForwardingPathAvailability
	HasIEExtensions                     bool
	IEExtensions                        []HandoverRequiredTransferExtIEsEntry
	UnknownExtensions                   []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HandoverRequiredTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("HandoverRequiredTransfer", v)
}

// DecodeHandoverRequiredTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeHandoverRequiredTransfer(data []byte) (HandoverRequiredTransfer, int, error) {
	var v HandoverRequiredTransfer
	err := ngapcodec.DecodeParameter("HandoverRequiredTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return HandoverRequiredTransfer{}, 0, err
	}
	return v, len(data), err
}
