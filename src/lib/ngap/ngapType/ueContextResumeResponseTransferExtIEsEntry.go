// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UEContextResumeResponseTransferExtIEsEntryPresent uint32

const (
	UEContextResumeResponseTransferExtIEsEntryPresentNothing UEContextResumeResponseTransferExtIEsEntryPresent = iota
	UEContextResumeResponseTransferExtIEsEntryPresentUnknown
)

type UEContextResumeResponseTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      UEContextResumeResponseTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEContextResumeResponseTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEContextResumeResponseTransfer-ExtIEsEntry", v)
}

// DecodeUEContextResumeResponseTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUEContextResumeResponseTransferExtIEsEntry(data []byte) (UEContextResumeResponseTransferExtIEsEntry, int, error) {
	var v UEContextResumeResponseTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("UEContextResumeResponseTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEContextResumeResponseTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
