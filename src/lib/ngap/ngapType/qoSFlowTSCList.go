// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QoSFlowTSCList represents ASN.1 QoSFlowTSCList.
type QoSFlowTSCList struct {
	List []QoSFlowTSCItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QoSFlowTSCList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QoSFlowTSCList", v)
}

// DecodeQoSFlowTSCList decodes a standalone APER value and reports the consumed octets.
func DecodeQoSFlowTSCList(data []byte) (QoSFlowTSCList, int, error) {
	var v QoSFlowTSCList
	err := ngapcodec.DecodeParameter("QoSFlowTSCList", data, &v)
	if ngapcodec.IsFatal(err) {
		return QoSFlowTSCList{}, 0, err
	}
	return v, len(data), err
}
