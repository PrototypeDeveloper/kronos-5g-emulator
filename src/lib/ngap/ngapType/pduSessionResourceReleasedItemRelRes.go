// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceReleasedItemRelRes represents ASN.1 PDUSessionResourceReleasedItemRelRes.
type PDUSessionResourceReleasedItemRelRes struct {
	PDUSessionID                              PDUSessionID
	PDUSessionResourceReleaseResponseTransfer []byte
	HasIEExtensions                           bool
	IEExtensions                              []PDUSessionResourceReleasedItemRelResExtIEsEntry
	UnknownExtensions                         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceReleasedItemRelRes) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceReleasedItemRelRes", v)
}

// DecodePDUSessionResourceReleasedItemRelRes decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceReleasedItemRelRes(data []byte) (PDUSessionResourceReleasedItemRelRes, int, error) {
	var v PDUSessionResourceReleasedItemRelRes
	err := ngapcodec.DecodeParameter("PDUSessionResourceReleasedItemRelRes", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceReleasedItemRelRes{}, 0, err
	}
	return v, len(data), err
}
