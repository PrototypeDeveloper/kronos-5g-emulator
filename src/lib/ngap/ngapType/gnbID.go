// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GNBID represents ASN.1 GNB-ID.
type GNBIDPresent uint32

const (
	GNBIDPresentNothing GNBIDPresent = iota
	GNBIDPresentGNBID
	GNBIDPresentChoiceExtensions
)

type GNBID struct {
	Present          GNBIDPresent
	GNBID            BitString `aper:"sizeLB:22,sizeUB:32"`
	ChoiceExtensions GNBIDExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GNBID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("GNB-ID", v) }

// DecodeGNBID decodes a standalone APER value and reports the consumed octets.
func DecodeGNBID(data []byte) (GNBID, int, error) {
	var v GNBID
	err := ngapcodec.DecodeParameter("GNB-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return GNBID{}, 0, err
	}
	return v, len(data), err
}
