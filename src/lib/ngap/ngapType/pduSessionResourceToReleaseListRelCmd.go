// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceToReleaseListRelCmd represents ASN.1 PDUSessionResourceToReleaseListRelCmd.
type PDUSessionResourceToReleaseListRelCmd struct {
	List []PDUSessionResourceToReleaseItemRelCmd `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceToReleaseListRelCmd) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceToReleaseListRelCmd", v)
}

// DecodePDUSessionResourceToReleaseListRelCmd decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceToReleaseListRelCmd(data []byte) (PDUSessionResourceToReleaseListRelCmd, int, error) {
	var v PDUSessionResourceToReleaseListRelCmd
	err := ngapcodec.DecodeParameter("PDUSessionResourceToReleaseListRelCmd", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceToReleaseListRelCmd{}, 0, err
	}
	return v, len(data), err
}
