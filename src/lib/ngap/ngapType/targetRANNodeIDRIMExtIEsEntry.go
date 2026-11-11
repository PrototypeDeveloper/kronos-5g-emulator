// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TargetRANNodeIDRIMExtIEsEntryPresent uint32

const (
	TargetRANNodeIDRIMExtIEsEntryPresentNothing TargetRANNodeIDRIMExtIEsEntryPresent = iota
	TargetRANNodeIDRIMExtIEsEntryPresentUnknown
)

type TargetRANNodeIDRIMExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TargetRANNodeIDRIMExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetRANNodeIDRIMExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargetRANNodeID-RIM-ExtIEsEntry", v)
}

// DecodeTargetRANNodeIDRIMExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTargetRANNodeIDRIMExtIEsEntry(data []byte) (TargetRANNodeIDRIMExtIEsEntry, int, error) {
	var v TargetRANNodeIDRIMExtIEsEntry
	err := ngapcodec.DecodeParameter("TargetRANNodeID-RIM-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetRANNodeIDRIMExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
