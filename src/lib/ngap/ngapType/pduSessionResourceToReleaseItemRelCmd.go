// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceToReleaseItemRelCmd represents ASN.1 PDUSessionResourceToReleaseItemRelCmd.
type PDUSessionResourceToReleaseItemRelCmd struct {
	PDUSessionID                             PDUSessionID
	PDUSessionResourceReleaseCommandTransfer []byte
	HasIEExtensions                          bool
	IEExtensions                             []PDUSessionResourceToReleaseItemRelCmdExtIEsEntry
	UnknownExtensions                        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceToReleaseItemRelCmd) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceToReleaseItemRelCmd", v)
}

// DecodePDUSessionResourceToReleaseItemRelCmd decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceToReleaseItemRelCmd(data []byte) (PDUSessionResourceToReleaseItemRelCmd, int, error) {
	var v PDUSessionResourceToReleaseItemRelCmd
	err := ngapcodec.DecodeParameter("PDUSessionResourceToReleaseItemRelCmd", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceToReleaseItemRelCmd{}, 0, err
	}
	return v, len(data), err
}
