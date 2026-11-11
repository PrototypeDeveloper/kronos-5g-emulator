// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PCIListForMDT represents ASN.1 PCIListForMDT.
type PCIListForMDT struct {
	List []NRPCI `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PCIListForMDT) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("PCIListForMDT", v) }

// DecodePCIListForMDT decodes a standalone APER value and reports the consumed octets.
func DecodePCIListForMDT(data []byte) (PCIListForMDT, int, error) {
	var v PCIListForMDT
	err := ngapcodec.DecodeParameter("PCIListForMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return PCIListForMDT{}, 0, err
	}
	return v, len(data), err
}
