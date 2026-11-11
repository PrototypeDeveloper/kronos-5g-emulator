// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSSessionTNLInfoNGRANExtIEsEntryPresent uint32

const (
	MBSSessionTNLInfoNGRANExtIEsEntryPresentNothing MBSSessionTNLInfoNGRANExtIEsEntryPresent = iota
	MBSSessionTNLInfoNGRANExtIEsEntryPresentUnknown
)

type MBSSessionTNLInfoNGRANExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      MBSSessionTNLInfoNGRANExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionTNLInfoNGRANExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-SessionTNLInfoNGRAN-ExtIEsEntry", v)
}

// DecodeMBSSessionTNLInfoNGRANExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionTNLInfoNGRANExtIEsEntry(data []byte) (MBSSessionTNLInfoNGRANExtIEsEntry, int, error) {
	var v MBSSessionTNLInfoNGRANExtIEsEntry
	err := ngapcodec.DecodeParameter("MBS-SessionTNLInfoNGRAN-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionTNLInfoNGRANExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
