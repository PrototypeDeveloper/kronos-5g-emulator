// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LastVisitedGERANCellInformation represents ASN.1 LastVisitedGERANCellInformation.
type LastVisitedGERANCellInformation struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LastVisitedGERANCellInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LastVisitedGERANCellInformation", v)
}

// DecodeLastVisitedGERANCellInformation decodes a standalone APER value and reports the consumed octets.
func DecodeLastVisitedGERANCellInformation(data []byte) (LastVisitedGERANCellInformation, int, error) {
	var v LastVisitedGERANCellInformation
	err := ngapcodec.DecodeParameter("LastVisitedGERANCellInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return LastVisitedGERANCellInformation{}, 0, err
	}
	return v, len(data), err
}
