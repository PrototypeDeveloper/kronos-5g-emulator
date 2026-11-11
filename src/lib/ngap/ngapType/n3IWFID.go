// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// N3IWFID represents ASN.1 N3IWF-ID.
type N3IWFIDPresent uint32

const (
	N3IWFIDPresentNothing N3IWFIDPresent = iota
	N3IWFIDPresentN3IWFID
	N3IWFIDPresentChoiceExtensions
)

type N3IWFID struct {
	Present          N3IWFIDPresent
	N3IWFID          BitString `aper:"sizeLB:16,sizeUB:16"`
	ChoiceExtensions N3IWFIDExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v N3IWFID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("N3IWF-ID", v) }

// DecodeN3IWFID decodes a standalone APER value and reports the consumed octets.
func DecodeN3IWFID(data []byte) (N3IWFID, int, error) {
	var v N3IWFID
	err := ngapcodec.DecodeParameter("N3IWF-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return N3IWFID{}, 0, err
	}
	return v, len(data), err
}
