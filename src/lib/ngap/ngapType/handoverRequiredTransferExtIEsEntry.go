// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type HandoverRequiredTransferExtIEsEntryPresent uint32

const (
	HandoverRequiredTransferExtIEsEntryPresentNothing HandoverRequiredTransferExtIEsEntryPresent = iota
	HandoverRequiredTransferExtIEsEntryPresentUnknown
)

type HandoverRequiredTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      HandoverRequiredTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HandoverRequiredTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("HandoverRequiredTransfer-ExtIEsEntry", v)
}

// DecodeHandoverRequiredTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeHandoverRequiredTransferExtIEsEntry(data []byte) (HandoverRequiredTransferExtIEsEntry, int, error) {
	var v HandoverRequiredTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("HandoverRequiredTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return HandoverRequiredTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
