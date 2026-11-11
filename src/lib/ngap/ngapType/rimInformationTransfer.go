// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RIMInformationTransfer represents ASN.1 RIMInformationTransfer.
type RIMInformationTransfer struct {
	TargetRANNodeIDRIM TargetRANNodeIDRIM
	SourceRANNodeID    SourceRANNodeID
	RIMInformation     RIMInformation
	HasIEExtensions    bool
	IEExtensions       []RIMInformationTransferExtIEsEntry
	UnknownExtensions  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RIMInformationTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RIMInformationTransfer", v)
}

// DecodeRIMInformationTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeRIMInformationTransfer(data []byte) (RIMInformationTransfer, int, error) {
	var v RIMInformationTransfer
	err := ngapcodec.DecodeParameter("RIMInformationTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return RIMInformationTransfer{}, 0, err
	}
	return v, len(data), err
}
