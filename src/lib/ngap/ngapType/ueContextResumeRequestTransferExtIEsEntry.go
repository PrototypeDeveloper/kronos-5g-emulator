// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UEContextResumeRequestTransferExtIEsEntryPresent uint32

const (
	UEContextResumeRequestTransferExtIEsEntryPresentNothing UEContextResumeRequestTransferExtIEsEntryPresent = iota
	UEContextResumeRequestTransferExtIEsEntryPresentUnknown
)

type UEContextResumeRequestTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      UEContextResumeRequestTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEContextResumeRequestTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEContextResumeRequestTransfer-ExtIEsEntry", v)
}

// DecodeUEContextResumeRequestTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUEContextResumeRequestTransferExtIEsEntry(data []byte) (UEContextResumeRequestTransferExtIEsEntry, int, error) {
	var v UEContextResumeRequestTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("UEContextResumeRequestTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEContextResumeRequestTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
