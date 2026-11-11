// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QMCConfigInfo represents ASN.1 QMCConfigInfo.
type QMCConfigInfo struct {
	UEAppLayerMeasInfoList UEAppLayerMeasInfoList
	HasIEExtensions        bool
	IEExtensions           []QMCConfigInfoExtIEsEntry
	UnknownExtensions      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QMCConfigInfo) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("QMCConfigInfo", v) }

// DecodeQMCConfigInfo decodes a standalone APER value and reports the consumed octets.
func DecodeQMCConfigInfo(data []byte) (QMCConfigInfo, int, error) {
	var v QMCConfigInfo
	err := ngapcodec.DecodeParameter("QMCConfigInfo", data, &v)
	if ngapcodec.IsFatal(err) {
		return QMCConfigInfo{}, 0, err
	}
	return v, len(data), err
}
