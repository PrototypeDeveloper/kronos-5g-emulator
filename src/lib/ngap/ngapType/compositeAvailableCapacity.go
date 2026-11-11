// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CompositeAvailableCapacity represents ASN.1 CompositeAvailableCapacity.
type CompositeAvailableCapacity struct {
	HasCellCapacityClassValue bool
	CellCapacityClassValue    int64 `aper:"valueExt,valueLB:1,valueUB:100"`
	CapacityValue             int64 `aper:"valueLB:0,valueUB:100"`
	HasIEExtensions           bool
	IEExtensions              []CompositeAvailableCapacityExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CompositeAvailableCapacity) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CompositeAvailableCapacity", v)
}

// DecodeCompositeAvailableCapacity decodes a standalone APER value and reports the consumed octets.
func DecodeCompositeAvailableCapacity(data []byte) (CompositeAvailableCapacity, int, error) {
	var v CompositeAvailableCapacity
	err := ngapcodec.DecodeParameter("CompositeAvailableCapacity", data, &v)
	if ngapcodec.IsFatal(err) {
		return CompositeAvailableCapacity{}, 0, err
	}
	return v, len(data), err
}
