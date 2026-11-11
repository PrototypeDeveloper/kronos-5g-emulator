// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRCGIListForWarning represents ASN.1 NR-CGIListForWarning.
type NRCGIListForWarning struct {
	List []NRCGI `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRCGIListForWarning) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NR-CGIListForWarning", v)
}

// DecodeNRCGIListForWarning decodes a standalone APER value and reports the consumed octets.
func DecodeNRCGIListForWarning(data []byte) (NRCGIListForWarning, int, error) {
	var v NRCGIListForWarning
	err := ngapcodec.DecodeParameter("NR-CGIListForWarning", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRCGIListForWarning{}, 0, err
	}
	return v, len(data), err
}
