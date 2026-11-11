// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NBIoTPagingEDRXInfo represents ASN.1 NB-IoT-Paging-eDRXInfo.
type NBIoTPagingEDRXInfo struct {
	NBIoTPagingEDRXCycle     NBIoTPagingEDRXCycle
	HasNBIoTPagingTimeWindow bool
	NBIoTPagingTimeWindow    NBIoTPagingTimeWindow
	HasIEExtensions          bool
	IEExtensions             []NBIoTPagingEDRXInfoExtIEsEntry
	UnknownExtensions        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NBIoTPagingEDRXInfo) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NB-IoT-Paging-eDRXInfo", v)
}

// DecodeNBIoTPagingEDRXInfo decodes a standalone APER value and reports the consumed octets.
func DecodeNBIoTPagingEDRXInfo(data []byte) (NBIoTPagingEDRXInfo, int, error) {
	var v NBIoTPagingEDRXInfo
	err := ngapcodec.DecodeParameter("NB-IoT-Paging-eDRXInfo", data, &v)
	if ngapcodec.IsFatal(err) {
		return NBIoTPagingEDRXInfo{}, 0, err
	}
	return v, len(data), err
}
