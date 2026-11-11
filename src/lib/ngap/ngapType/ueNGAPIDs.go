// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UENGAPIDs represents ASN.1 UE-NGAP-IDs.
type UENGAPIDsPresent uint32

const (
	UENGAPIDsPresentNothing UENGAPIDsPresent = iota
	UENGAPIDsPresentUENGAPIDPair
	UENGAPIDsPresentAMFUENGAPID
	UENGAPIDsPresentChoiceExtensions
)

type UENGAPIDs struct {
	Present          UENGAPIDsPresent
	UENGAPIDPair     UENGAPIDPair
	AMFUENGAPID      AMFUENGAPID
	ChoiceExtensions UENGAPIDsExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UENGAPIDs) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("UE-NGAP-IDs", v) }

// DecodeUENGAPIDs decodes a standalone APER value and reports the consumed octets.
func DecodeUENGAPIDs(data []byte) (UENGAPIDs, int, error) {
	var v UENGAPIDs
	err := ngapcodec.DecodeParameter("UE-NGAP-IDs", data, &v)
	if ngapcodec.IsFatal(err) {
		return UENGAPIDs{}, 0, err
	}
	return v, len(data), err
}
