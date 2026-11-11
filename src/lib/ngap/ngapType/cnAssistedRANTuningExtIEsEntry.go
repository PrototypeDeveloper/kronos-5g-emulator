// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CNAssistedRANTuningExtIEsEntryPresent uint32

const (
	CNAssistedRANTuningExtIEsEntryPresentNothing CNAssistedRANTuningExtIEsEntryPresent = iota
	CNAssistedRANTuningExtIEsEntryPresentUnknown
)

type CNAssistedRANTuningExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CNAssistedRANTuningExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CNAssistedRANTuningExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CNAssistedRANTuning-ExtIEsEntry", v)
}

// DecodeCNAssistedRANTuningExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCNAssistedRANTuningExtIEsEntry(data []byte) (CNAssistedRANTuningExtIEsEntry, int, error) {
	var v CNAssistedRANTuningExtIEsEntry
	err := ngapcodec.DecodeParameter("CNAssistedRANTuning-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CNAssistedRANTuningExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
