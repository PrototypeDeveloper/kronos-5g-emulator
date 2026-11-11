// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EUTRANCompositeAvailableCapacityGroup represents ASN.1 EUTRAN-CompositeAvailableCapacityGroup.
type EUTRANCompositeAvailableCapacityGroup struct {
	DLCompositeAvailableCapacity CompositeAvailableCapacity
	ULCompositeAvailableCapacity CompositeAvailableCapacity
	HasIEExtensions              bool
	IEExtensions                 []EUTRANCompositeAvailableCapacityGroupExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRANCompositeAvailableCapacityGroup) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRAN-CompositeAvailableCapacityGroup", v)
}

// DecodeEUTRANCompositeAvailableCapacityGroup decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRANCompositeAvailableCapacityGroup(data []byte) (EUTRANCompositeAvailableCapacityGroup, int, error) {
	var v EUTRANCompositeAvailableCapacityGroup
	err := ngapcodec.DecodeParameter("EUTRAN-CompositeAvailableCapacityGroup", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRANCompositeAvailableCapacityGroup{}, 0, err
	}
	return v, len(data), err
}
