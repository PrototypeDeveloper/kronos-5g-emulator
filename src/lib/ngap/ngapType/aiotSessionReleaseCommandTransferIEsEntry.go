// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AIOTSessionReleaseCommandTransferIEsEntryPresent uint32

const (
	AIOTSessionReleaseCommandTransferIEsEntryPresentNothing AIOTSessionReleaseCommandTransferIEsEntryPresent = iota
	AIOTSessionReleaseCommandTransferIEsEntryPresentAIoTCorrelationIdentifier
	AIOTSessionReleaseCommandTransferIEsEntryPresentCause
	AIOTSessionReleaseCommandTransferIEsEntryPresentUnknown
)

type AIOTSessionReleaseCommandTransferIEsEntry struct {
	ID                        ProtocolIEID
	Criticality               Criticality
	Present                   AIOTSessionReleaseCommandTransferIEsEntryPresent
	AIoTCorrelationIdentifier AIoTCorrelationIdentifier `ngap:"id:466,criticality:reject,presence:mandatory"`
	Cause                     Cause                     `ngap:"id:15,criticality:ignore,presence:mandatory"`
	UnknownValue              []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIOTSessionReleaseCommandTransferIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIOTSessionReleaseCommandTransferIEsEntry", v)
}

// DecodeAIOTSessionReleaseCommandTransferIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAIOTSessionReleaseCommandTransferIEsEntry(data []byte) (AIOTSessionReleaseCommandTransferIEsEntry, int, error) {
	var v AIOTSessionReleaseCommandTransferIEsEntry
	err := ngapcodec.DecodeParameter("AIOTSessionReleaseCommandTransferIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIOTSessionReleaseCommandTransferIEsEntry{}, 0, err
	}
	return v, len(data), err
}
