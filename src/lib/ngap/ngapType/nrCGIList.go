// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRCGIList represents ASN.1 NR-CGIList.
type NRCGIList struct {
	List []NRCGI `aper:"sizeLB:1,sizeUB:16384"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRCGIList) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("NR-CGIList", v) }

// DecodeNRCGIList decodes a standalone APER value and reports the consumed octets.
func DecodeNRCGIList(data []byte) (NRCGIList, int, error) {
	var v NRCGIList
	err := ngapcodec.DecodeParameter("NR-CGIList", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRCGIList{}, 0, err
	}
	return v, len(data), err
}
