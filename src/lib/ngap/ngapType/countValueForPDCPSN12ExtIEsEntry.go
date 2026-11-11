// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type COUNTValueForPDCPSN12ExtIEsEntryPresent uint32

const (
	COUNTValueForPDCPSN12ExtIEsEntryPresentNothing COUNTValueForPDCPSN12ExtIEsEntryPresent = iota
	COUNTValueForPDCPSN12ExtIEsEntryPresentUnknown
)

type COUNTValueForPDCPSN12ExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      COUNTValueForPDCPSN12ExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v COUNTValueForPDCPSN12ExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("COUNTValueForPDCP-SN12-ExtIEsEntry", v)
}

// DecodeCOUNTValueForPDCPSN12ExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCOUNTValueForPDCPSN12ExtIEsEntry(data []byte) (COUNTValueForPDCPSN12ExtIEsEntry, int, error) {
	var v COUNTValueForPDCPSN12ExtIEsEntry
	err := ngapcodec.DecodeParameter("COUNTValueForPDCP-SN12-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return COUNTValueForPDCPSN12ExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
