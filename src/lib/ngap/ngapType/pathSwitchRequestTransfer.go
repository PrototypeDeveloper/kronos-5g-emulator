// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PathSwitchRequestTransfer represents ASN.1 PathSwitchRequestTransfer.
type PathSwitchRequestTransfer struct {
	DLNGUUPTNLInformation           UPTransportLayerInformation
	HasDLNGUTNLInformationReused    bool
	DLNGUTNLInformationReused       DLNGUTNLInformationReused
	HasUserPlaneSecurityInformation bool
	UserPlaneSecurityInformation    UserPlaneSecurityInformation
	QosFlowAcceptedList             QosFlowAcceptedList
	HasIEExtensions                 bool
	IEExtensions                    []PathSwitchRequestTransferExtIEsEntry
	UnknownExtensions               []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PathSwitchRequestTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PathSwitchRequestTransfer", v)
}

// DecodePathSwitchRequestTransfer decodes a standalone APER value and reports the consumed octets.
func DecodePathSwitchRequestTransfer(data []byte) (PathSwitchRequestTransfer, int, error) {
	var v PathSwitchRequestTransfer
	err := ngapcodec.DecodeParameter("PathSwitchRequestTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return PathSwitchRequestTransfer{}, 0, err
	}
	return v, len(data), err
}
