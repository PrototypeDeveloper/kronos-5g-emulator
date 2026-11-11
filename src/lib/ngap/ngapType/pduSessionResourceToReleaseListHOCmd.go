// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceToReleaseListHOCmd represents ASN.1 PDUSessionResourceToReleaseListHOCmd.
type PDUSessionResourceToReleaseListHOCmd struct {
	List []PDUSessionResourceToReleaseItemHOCmd `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceToReleaseListHOCmd) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceToReleaseListHOCmd", v)
}

// DecodePDUSessionResourceToReleaseListHOCmd decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceToReleaseListHOCmd(data []byte) (PDUSessionResourceToReleaseListHOCmd, int, error) {
	var v PDUSessionResourceToReleaseListHOCmd
	err := ngapcodec.DecodeParameter("PDUSessionResourceToReleaseListHOCmd", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceToReleaseListHOCmd{}, 0, err
	}
	return v, len(data), err
}
