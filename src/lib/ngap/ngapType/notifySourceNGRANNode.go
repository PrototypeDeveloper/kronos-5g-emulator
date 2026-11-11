// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NotifySourceNGRANNode represents ASN.1 NotifySourceNGRANNode.
const (
	NotifySourceNGRANNodePresentNotifySource Enumerated = 0
)

type NotifySourceNGRANNode struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NotifySourceNGRANNode) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NotifySourceNGRANNode", v)
}

// DecodeNotifySourceNGRANNode decodes a standalone APER value and reports the consumed octets.
func DecodeNotifySourceNGRANNode(data []byte) (NotifySourceNGRANNode, int, error) {
	var v NotifySourceNGRANNode
	err := ngapcodec.DecodeParameter("NotifySourceNGRANNode", data, &v)
	if ngapcodec.IsFatal(err) {
		return NotifySourceNGRANNode{}, 0, err
	}
	return v, len(data), err
}
