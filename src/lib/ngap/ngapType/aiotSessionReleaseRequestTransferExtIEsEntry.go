// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AIOTSessionReleaseRequestTransferExtIEsEntryPresent uint32

const (
	AIOTSessionReleaseRequestTransferExtIEsEntryPresentNothing AIOTSessionReleaseRequestTransferExtIEsEntryPresent = iota
	AIOTSessionReleaseRequestTransferExtIEsEntryPresentUnknown
)

type AIOTSessionReleaseRequestTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AIOTSessionReleaseRequestTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIOTSessionReleaseRequestTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIOTSessionReleaseRequestTransfer-ExtIEsEntry", v)
}

// DecodeAIOTSessionReleaseRequestTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAIOTSessionReleaseRequestTransferExtIEsEntry(data []byte) (AIOTSessionReleaseRequestTransferExtIEsEntry, int, error) {
	var v AIOTSessionReleaseRequestTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("AIOTSessionReleaseRequestTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIOTSessionReleaseRequestTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
