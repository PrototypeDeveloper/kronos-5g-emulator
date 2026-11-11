// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSDistributionReleaseRequestTransfer represents ASN.1 MBS-DistributionReleaseRequestTransfer.
type MBSDistributionReleaseRequestTransfer struct {
	MBSSessionID                      MBSSessionID
	HasMBSAreaSessionID               bool
	MBSAreaSessionID                  MBSAreaSessionID
	HasSharedNGUUnicastTNLInformation bool
	SharedNGUUnicastTNLInformation    UPTransportLayerInformation
	Cause                             Cause
	HasIEExtensions                   bool
	IEExtensions                      []MBSDistributionReleaseRequesTransferExtIEsEntry
	UnknownExtensions                 []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSDistributionReleaseRequestTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-DistributionReleaseRequestTransfer", v)
}

// DecodeMBSDistributionReleaseRequestTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeMBSDistributionReleaseRequestTransfer(data []byte) (MBSDistributionReleaseRequestTransfer, int, error) {
	var v MBSDistributionReleaseRequestTransfer
	err := ngapcodec.DecodeParameter("MBS-DistributionReleaseRequestTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSDistributionReleaseRequestTransfer{}, 0, err
	}
	return v, len(data), err
}
