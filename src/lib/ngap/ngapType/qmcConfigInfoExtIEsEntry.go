// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QMCConfigInfoExtIEsEntryPresent uint32

const (
	QMCConfigInfoExtIEsEntryPresentNothing QMCConfigInfoExtIEsEntryPresent = iota
	QMCConfigInfoExtIEsEntryPresentUnknown
)

type QMCConfigInfoExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      QMCConfigInfoExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QMCConfigInfoExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QMCConfigInfo-ExtIEsEntry", v)
}

// DecodeQMCConfigInfoExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQMCConfigInfoExtIEsEntry(data []byte) (QMCConfigInfoExtIEsEntry, int, error) {
	var v QMCConfigInfoExtIEsEntry
	err := ngapcodec.DecodeParameter("QMCConfigInfo-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QMCConfigInfoExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
