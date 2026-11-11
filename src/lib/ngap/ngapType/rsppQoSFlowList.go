// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RSPPQoSFlowList represents ASN.1 RSPPQoSFlowList.
type RSPPQoSFlowList struct {
	List []RSPPQoSFlowItem `aper:"sizeLB:1,sizeUB:2048"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RSPPQoSFlowList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RSPPQoSFlowList", v)
}

// DecodeRSPPQoSFlowList decodes a standalone APER value and reports the consumed octets.
func DecodeRSPPQoSFlowList(data []byte) (RSPPQoSFlowList, int, error) {
	var v RSPPQoSFlowList
	err := ngapcodec.DecodeParameter("RSPPQoSFlowList", data, &v)
	if ngapcodec.IsFatal(err) {
		return RSPPQoSFlowList{}, 0, err
	}
	return v, len(data), err
}
