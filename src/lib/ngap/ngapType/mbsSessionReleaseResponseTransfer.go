// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionReleaseResponseTransfer represents ASN.1 MBSSessionReleaseResponseTransfer.
type MBSSessionReleaseResponseTransfer struct {
	HasMBSSessionTNLInfoNGRAN bool
	MBSSessionTNLInfoNGRAN    MBSSessionTNLInfoNGRAN
	HasIEExtensions           bool
	IEExtensions              []MBSSessionReleaseResponseTransferExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionReleaseResponseTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionReleaseResponseTransfer", v)
}

// DecodeMBSSessionReleaseResponseTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionReleaseResponseTransfer(data []byte) (MBSSessionReleaseResponseTransfer, int, error) {
	var v MBSSessionReleaseResponseTransfer
	err := ngapcodec.DecodeParameter("MBSSessionReleaseResponseTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionReleaseResponseTransfer{}, 0, err
	}
	return v, len(data), err
}
