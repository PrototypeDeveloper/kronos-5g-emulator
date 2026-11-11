// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SLPositioningRangingServiceInfo represents ASN.1 SLPositioningRangingServiceInfo.
type SLPositioningRangingServiceInfo struct {
	SLPositioningRangingAuthorized       SLPositioningRangingAuthorized
	HasSLPositioningRangingQoSParameters bool
	SLPositioningRangingQoSParameters    SLPositioningRangingQoSParameters
	HasIEExtensions                      bool
	IEExtensions                         []SLPositioningRangingServiceInfoExtIEsEntry
	UnknownExtensions                    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SLPositioningRangingServiceInfo) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SLPositioningRangingServiceInfo", v)
}

// DecodeSLPositioningRangingServiceInfo decodes a standalone APER value and reports the consumed octets.
func DecodeSLPositioningRangingServiceInfo(data []byte) (SLPositioningRangingServiceInfo, int, error) {
	var v SLPositioningRangingServiceInfo
	err := ngapcodec.DecodeParameter("SLPositioningRangingServiceInfo", data, &v)
	if ngapcodec.IsFatal(err) {
		return SLPositioningRangingServiceInfo{}, 0, err
	}
	return v, len(data), err
}
