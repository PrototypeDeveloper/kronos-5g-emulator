// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type RedundantPDUSessionInformationExtIEsEntryPresent uint32

const (
	RedundantPDUSessionInformationExtIEsEntryPresentNothing RedundantPDUSessionInformationExtIEsEntryPresent = iota
	RedundantPDUSessionInformationExtIEsEntryPresentPDUSessionPairID
	RedundantPDUSessionInformationExtIEsEntryPresentUnknown
)

type RedundantPDUSessionInformationExtIEsEntry struct {
	ID               ProtocolExtensionID
	Criticality      Criticality
	Present          RedundantPDUSessionInformationExtIEsEntryPresent
	PDUSessionPairID PDUSessionPairID `ngap:"id:331,criticality:ignore,presence:optional"`
	UnknownValue     []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RedundantPDUSessionInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RedundantPDUSessionInformation-ExtIEsEntry", v)
}

// DecodeRedundantPDUSessionInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeRedundantPDUSessionInformationExtIEsEntry(data []byte) (RedundantPDUSessionInformationExtIEsEntry, int, error) {
	var v RedundantPDUSessionInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("RedundantPDUSessionInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return RedundantPDUSessionInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
