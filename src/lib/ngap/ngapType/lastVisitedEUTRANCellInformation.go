// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LastVisitedEUTRANCellInformation represents ASN.1 LastVisitedEUTRANCellInformation.
type LastVisitedEUTRANCellInformation struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LastVisitedEUTRANCellInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LastVisitedEUTRANCellInformation", v)
}

// DecodeLastVisitedEUTRANCellInformation decodes a standalone APER value and reports the consumed octets.
func DecodeLastVisitedEUTRANCellInformation(data []byte) (LastVisitedEUTRANCellInformation, int, error) {
	var v LastVisitedEUTRANCellInformation
	err := ngapcodec.DecodeParameter("LastVisitedEUTRANCellInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return LastVisitedEUTRANCellInformation{}, 0, err
	}
	return v, len(data), err
}
