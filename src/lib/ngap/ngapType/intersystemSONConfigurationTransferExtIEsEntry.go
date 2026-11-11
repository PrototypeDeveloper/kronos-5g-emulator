// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type IntersystemSONConfigurationTransferExtIEsEntryPresent uint32

const (
	IntersystemSONConfigurationTransferExtIEsEntryPresentNothing IntersystemSONConfigurationTransferExtIEsEntryPresent = iota
	IntersystemSONConfigurationTransferExtIEsEntryPresentUnknown
)

type IntersystemSONConfigurationTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      IntersystemSONConfigurationTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemSONConfigurationTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemSONConfigurationTransfer-ExtIEsEntry", v)
}

// DecodeIntersystemSONConfigurationTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemSONConfigurationTransferExtIEsEntry(data []byte) (IntersystemSONConfigurationTransferExtIEsEntry, int, error) {
	var v IntersystemSONConfigurationTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("IntersystemSONConfigurationTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemSONConfigurationTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
