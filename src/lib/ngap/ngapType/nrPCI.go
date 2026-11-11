// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRPCI represents ASN.1 NR-PCI.
type NRPCI struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:1007"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRPCI) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("NR-PCI", v) }

// DecodeNRPCI decodes a standalone APER value and reports the consumed octets.
func DecodeNRPCI(data []byte) (NRPCI, int, error) {
	var v NRPCI
	err := ngapcodec.DecodeParameter("NR-PCI", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRPCI{}, 0, err
	}
	return v, len(data), err
}
