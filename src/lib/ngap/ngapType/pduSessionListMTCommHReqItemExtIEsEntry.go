// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionListMTCommHReqItemExtIEsEntryPresent uint32

const (
	PDUSessionListMTCommHReqItemExtIEsEntryPresentNothing PDUSessionListMTCommHReqItemExtIEsEntryPresent = iota
	PDUSessionListMTCommHReqItemExtIEsEntryPresentUnknown
)

type PDUSessionListMTCommHReqItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionListMTCommHReqItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionListMTCommHReqItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionListMTCommHReq-Item-ExtIEsEntry", v)
}

// DecodePDUSessionListMTCommHReqItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionListMTCommHReqItemExtIEsEntry(data []byte) (PDUSessionListMTCommHReqItemExtIEsEntry, int, error) {
	var v PDUSessionListMTCommHReqItemExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionListMTCommHReq-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionListMTCommHReqItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
