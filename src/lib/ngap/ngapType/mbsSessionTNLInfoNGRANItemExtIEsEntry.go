// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSSessionTNLInfoNGRANItemExtIEsEntryPresent uint32

const (
	MBSSessionTNLInfoNGRANItemExtIEsEntryPresentNothing MBSSessionTNLInfoNGRANItemExtIEsEntryPresent = iota
	MBSSessionTNLInfoNGRANItemExtIEsEntryPresentUnknown
)

type MBSSessionTNLInfoNGRANItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MBSSessionTNLInfoNGRANItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionTNLInfoNGRANItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-SessionTNLInfoNGRANItem-ExtIEsEntry", v)
}

// DecodeMBSSessionTNLInfoNGRANItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionTNLInfoNGRANItemExtIEsEntry(data []byte) (MBSSessionTNLInfoNGRANItemExtIEsEntry, int, error) {
	var v MBSSessionTNLInfoNGRANItemExtIEsEntry
	err := ngapcodec.DecodeParameter("MBS-SessionTNLInfoNGRANItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionTNLInfoNGRANItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
