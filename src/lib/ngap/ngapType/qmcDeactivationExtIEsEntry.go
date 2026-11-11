// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QMCDeactivationExtIEsEntryPresent uint32

const (
	QMCDeactivationExtIEsEntryPresentNothing QMCDeactivationExtIEsEntryPresent = iota
	QMCDeactivationExtIEsEntryPresentUnknown
)

type QMCDeactivationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      QMCDeactivationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QMCDeactivationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QMCDeactivation-ExtIEsEntry", v)
}

// DecodeQMCDeactivationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQMCDeactivationExtIEsEntry(data []byte) (QMCDeactivationExtIEsEntry, int, error) {
	var v QMCDeactivationExtIEsEntry
	err := ngapcodec.DecodeParameter("QMCDeactivation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QMCDeactivationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
