// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionSetupOrModResponseTransfer represents ASN.1 MBSSessionSetupOrModResponseTransfer.
type MBSSessionSetupOrModResponseTransfer struct {
	HasMBSSessionTNLInfoNGRAN bool
	MBSSessionTNLInfoNGRAN    MBSSessionTNLInfoNGRAN
	HasIEExtensions           bool
	IEExtensions              []MBSSessionSetupOrModResponseTransferExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionSetupOrModResponseTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionSetupOrModResponseTransfer", v)
}

// DecodeMBSSessionSetupOrModResponseTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionSetupOrModResponseTransfer(data []byte) (MBSSessionSetupOrModResponseTransfer, int, error) {
	var v MBSSessionSetupOrModResponseTransfer
	err := ngapcodec.DecodeParameter("MBSSessionSetupOrModResponseTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionSetupOrModResponseTransfer{}, 0, err
	}
	return v, len(data), err
}
