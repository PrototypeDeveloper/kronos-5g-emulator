// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SLPositioningRangingServiceInfoExtIEsEntryPresent uint32

const (
	SLPositioningRangingServiceInfoExtIEsEntryPresentNothing SLPositioningRangingServiceInfoExtIEsEntryPresent = iota
	SLPositioningRangingServiceInfoExtIEsEntryPresentUnknown
)

type SLPositioningRangingServiceInfoExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SLPositioningRangingServiceInfoExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SLPositioningRangingServiceInfoExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SLPositioningRangingServiceInfo-ExtIEsEntry", v)
}

// DecodeSLPositioningRangingServiceInfoExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSLPositioningRangingServiceInfoExtIEsEntry(data []byte) (SLPositioningRangingServiceInfoExtIEsEntry, int, error) {
	var v SLPositioningRangingServiceInfoExtIEsEntry
	err := ngapcodec.DecodeParameter("SLPositioningRangingServiceInfo-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SLPositioningRangingServiceInfoExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
