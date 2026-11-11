// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceFailedToSetupItemPSReq represents ASN.1 PDUSessionResourceFailedToSetupItemPSReq.
type PDUSessionResourceFailedToSetupItemPSReq struct {
	PDUSessionID                         PDUSessionID
	PathSwitchRequestSetupFailedTransfer []byte
	HasIEExtensions                      bool
	IEExtensions                         []PDUSessionResourceFailedToSetupItemPSReqExtIEsEntry
	UnknownExtensions                    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToSetupItemPSReq) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToSetupItemPSReq", v)
}

// DecodePDUSessionResourceFailedToSetupItemPSReq decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToSetupItemPSReq(data []byte) (PDUSessionResourceFailedToSetupItemPSReq, int, error) {
	var v PDUSessionResourceFailedToSetupItemPSReq
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToSetupItemPSReq", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToSetupItemPSReq{}, 0, err
	}
	return v, len(data), err
}
