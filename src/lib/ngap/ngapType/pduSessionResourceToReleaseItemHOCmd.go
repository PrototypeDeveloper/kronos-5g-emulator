// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceToReleaseItemHOCmd represents ASN.1 PDUSessionResourceToReleaseItemHOCmd.
type PDUSessionResourceToReleaseItemHOCmd struct {
	PDUSessionID                            PDUSessionID
	HandoverPreparationUnsuccessfulTransfer []byte
	HasIEExtensions                         bool
	IEExtensions                            []PDUSessionResourceToReleaseItemHOCmdExtIEsEntry
	UnknownExtensions                       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceToReleaseItemHOCmd) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceToReleaseItemHOCmd", v)
}

// DecodePDUSessionResourceToReleaseItemHOCmd decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceToReleaseItemHOCmd(data []byte) (PDUSessionResourceToReleaseItemHOCmd, int, error) {
	var v PDUSessionResourceToReleaseItemHOCmd
	err := ngapcodec.DecodeParameter("PDUSessionResourceToReleaseItemHOCmd", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceToReleaseItemHOCmd{}, 0, err
	}
	return v, len(data), err
}
