// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EUTRACGIListForWarning represents ASN.1 EUTRA-CGIListForWarning.
type EUTRACGIListForWarning struct {
	List []EUTRACGI `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRACGIListForWarning) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRA-CGIListForWarning", v)
}

// DecodeEUTRACGIListForWarning decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRACGIListForWarning(data []byte) (EUTRACGIListForWarning, int, error) {
	var v EUTRACGIListForWarning
	err := ngapcodec.DecodeParameter("EUTRA-CGIListForWarning", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRACGIListForWarning{}, 0, err
	}
	return v, len(data), err
}
