// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceReleasedItemPSFail represents ASN.1 PDUSessionResourceReleasedItemPSFail.
type PDUSessionResourceReleasedItemPSFail struct {
	PDUSessionID                          PDUSessionID
	PathSwitchRequestUnsuccessfulTransfer []byte
	HasIEExtensions                       bool
	IEExtensions                          []PDUSessionResourceReleasedItemPSFailExtIEsEntry
	UnknownExtensions                     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceReleasedItemPSFail) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceReleasedItemPSFail", v)
}

// DecodePDUSessionResourceReleasedItemPSFail decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceReleasedItemPSFail(data []byte) (PDUSessionResourceReleasedItemPSFail, int, error) {
	var v PDUSessionResourceReleasedItemPSFail
	err := ngapcodec.DecodeParameter("PDUSessionResourceReleasedItemPSFail", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceReleasedItemPSFail{}, 0, err
	}
	return v, len(data), err
}
