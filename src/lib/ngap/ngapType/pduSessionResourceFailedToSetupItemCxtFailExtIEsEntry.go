// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceFailedToSetupItemCxtFailExtIEsEntryPresent uint32

const (
	PDUSessionResourceFailedToSetupItemCxtFailExtIEsEntryPresentNothing PDUSessionResourceFailedToSetupItemCxtFailExtIEsEntryPresent = iota
	PDUSessionResourceFailedToSetupItemCxtFailExtIEsEntryPresentUnknown
)

type PDUSessionResourceFailedToSetupItemCxtFailExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceFailedToSetupItemCxtFailExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToSetupItemCxtFailExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToSetupItemCxtFail-ExtIEsEntry", v)
}

// DecodePDUSessionResourceFailedToSetupItemCxtFailExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToSetupItemCxtFailExtIEsEntry(data []byte) (PDUSessionResourceFailedToSetupItemCxtFailExtIEsEntry, int, error) {
	var v PDUSessionResourceFailedToSetupItemCxtFailExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToSetupItemCxtFail-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToSetupItemCxtFailExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
