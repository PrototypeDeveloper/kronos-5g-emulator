// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type RIMInformationTransferExtIEsEntryPresent uint32

const (
	RIMInformationTransferExtIEsEntryPresentNothing RIMInformationTransferExtIEsEntryPresent = iota
	RIMInformationTransferExtIEsEntryPresentUnknown
)

type RIMInformationTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      RIMInformationTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RIMInformationTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RIMInformationTransfer-ExtIEsEntry", v)
}

// DecodeRIMInformationTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeRIMInformationTransferExtIEsEntry(data []byte) (RIMInformationTransferExtIEsEntry, int, error) {
	var v RIMInformationTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("RIMInformationTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return RIMInformationTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
