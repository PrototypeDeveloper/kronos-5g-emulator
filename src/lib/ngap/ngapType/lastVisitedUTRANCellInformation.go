// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LastVisitedUTRANCellInformation represents ASN.1 LastVisitedUTRANCellInformation.
type LastVisitedUTRANCellInformation struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LastVisitedUTRANCellInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LastVisitedUTRANCellInformation", v)
}

// DecodeLastVisitedUTRANCellInformation decodes a standalone APER value and reports the consumed octets.
func DecodeLastVisitedUTRANCellInformation(data []byte) (LastVisitedUTRANCellInformation, int, error) {
	var v LastVisitedUTRANCellInformation
	err := ngapcodec.DecodeParameter("LastVisitedUTRANCellInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return LastVisitedUTRANCellInformation{}, 0, err
	}
	return v, len(data), err
}
