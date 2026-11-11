// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SLPositioningRangingQoSParametersExtIEsEntryPresent uint32

const (
	SLPositioningRangingQoSParametersExtIEsEntryPresentNothing SLPositioningRangingQoSParametersExtIEsEntryPresent = iota
	SLPositioningRangingQoSParametersExtIEsEntryPresentUnknown
)

type SLPositioningRangingQoSParametersExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SLPositioningRangingQoSParametersExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SLPositioningRangingQoSParametersExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SLPositioningRangingQoSParameters-ExtIEsEntry", v)
}

// DecodeSLPositioningRangingQoSParametersExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSLPositioningRangingQoSParametersExtIEsEntry(data []byte) (SLPositioningRangingQoSParametersExtIEsEntry, int, error) {
	var v SLPositioningRangingQoSParametersExtIEsEntry
	err := ngapcodec.DecodeParameter("SLPositioningRangingQoSParameters-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SLPositioningRangingQoSParametersExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
