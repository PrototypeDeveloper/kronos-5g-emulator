// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TNGFID represents ASN.1 TNGF-ID.
type TNGFIDPresent uint32

const (
	TNGFIDPresentNothing TNGFIDPresent = iota
	TNGFIDPresentTNGFID
	TNGFIDPresentChoiceExtensions
)

type TNGFID struct {
	Present          TNGFIDPresent
	TNGFID           BitString `aper:"sizeExt,sizeLB:32,sizeUB:32"`
	ChoiceExtensions TNGFIDExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TNGFID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TNGF-ID", v) }

// DecodeTNGFID decodes a standalone APER value and reports the consumed octets.
func DecodeTNGFID(data []byte) (TNGFID, int, error) {
	var v TNGFID
	err := ngapcodec.DecodeParameter("TNGF-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return TNGFID{}, 0, err
	}
	return v, len(data), err
}
