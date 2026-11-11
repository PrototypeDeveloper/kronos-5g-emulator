// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AIOTSessionReleaseCompleteTransferExtIEsEntryPresent uint32

const (
	AIOTSessionReleaseCompleteTransferExtIEsEntryPresentNothing AIOTSessionReleaseCompleteTransferExtIEsEntryPresent = iota
	AIOTSessionReleaseCompleteTransferExtIEsEntryPresentUnknown
)

type AIOTSessionReleaseCompleteTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AIOTSessionReleaseCompleteTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIOTSessionReleaseCompleteTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIOTSessionReleaseCompleteTransfer-ExtIEsEntry", v)
}

// DecodeAIOTSessionReleaseCompleteTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAIOTSessionReleaseCompleteTransferExtIEsEntry(data []byte) (AIOTSessionReleaseCompleteTransferExtIEsEntry, int, error) {
	var v AIOTSessionReleaseCompleteTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("AIOTSessionReleaseCompleteTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIOTSessionReleaseCompleteTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
