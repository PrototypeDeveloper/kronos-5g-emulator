// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSSessionSetupOrModResponseTransferExtIEsEntryPresent uint32

const (
	MBSSessionSetupOrModResponseTransferExtIEsEntryPresentNothing MBSSessionSetupOrModResponseTransferExtIEsEntryPresent = iota
	MBSSessionSetupOrModResponseTransferExtIEsEntryPresentUnknown
)

type MBSSessionSetupOrModResponseTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MBSSessionSetupOrModResponseTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionSetupOrModResponseTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionSetupOrModResponseTransfer-ExtIEsEntry", v)
}

// DecodeMBSSessionSetupOrModResponseTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionSetupOrModResponseTransferExtIEsEntry(data []byte) (MBSSessionSetupOrModResponseTransferExtIEsEntry, int, error) {
	var v MBSSessionSetupOrModResponseTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("MBSSessionSetupOrModResponseTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionSetupOrModResponseTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
