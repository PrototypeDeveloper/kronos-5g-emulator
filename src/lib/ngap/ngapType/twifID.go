// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TWIFID represents ASN.1 TWIF-ID.
type TWIFIDPresent uint32

const (
	TWIFIDPresentNothing TWIFIDPresent = iota
	TWIFIDPresentTWIFID
	TWIFIDPresentChoiceExtensions
)

type TWIFID struct {
	Present          TWIFIDPresent
	TWIFID           BitString `aper:"sizeExt,sizeLB:32,sizeUB:32"`
	ChoiceExtensions TWIFIDExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TWIFID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TWIF-ID", v) }

// DecodeTWIFID decodes a standalone APER value and reports the consumed octets.
func DecodeTWIFID(data []byte) (TWIFID, int, error) {
	var v TWIFID
	err := ngapcodec.DecodeParameter("TWIF-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return TWIFID{}, 0, err
	}
	return v, len(data), err
}
