// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EUTRACGIList represents ASN.1 EUTRA-CGIList.
type EUTRACGIList struct {
	List []EUTRACGI `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRACGIList) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("EUTRA-CGIList", v) }

// DecodeEUTRACGIList decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRACGIList(data []byte) (EUTRACGIList, int, error) {
	var v EUTRACGIList
	err := ngapcodec.DecodeParameter("EUTRA-CGIList", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRACGIList{}, 0, err
	}
	return v, len(data), err
}
