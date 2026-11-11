// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// WAGFID represents ASN.1 W-AGF-ID.
type WAGFIDPresent uint32

const (
	WAGFIDPresentNothing WAGFIDPresent = iota
	WAGFIDPresentWAGFID
	WAGFIDPresentChoiceExtensions
)

type WAGFID struct {
	Present          WAGFIDPresent
	WAGFID           BitString `aper:"sizeExt,sizeLB:16,sizeUB:16"`
	ChoiceExtensions WAGFIDExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v WAGFID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("W-AGF-ID", v) }

// DecodeWAGFID decodes a standalone APER value and reports the consumed octets.
func DecodeWAGFID(data []byte) (WAGFID, int, error) {
	var v WAGFID
	err := ngapcodec.DecodeParameter("W-AGF-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return WAGFID{}, 0, err
	}
	return v, len(data), err
}
