// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type COUNTValueForPDCPSN18ExtIEsEntryPresent uint32

const (
	COUNTValueForPDCPSN18ExtIEsEntryPresentNothing COUNTValueForPDCPSN18ExtIEsEntryPresent = iota
	COUNTValueForPDCPSN18ExtIEsEntryPresentUnknown
)

type COUNTValueForPDCPSN18ExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      COUNTValueForPDCPSN18ExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v COUNTValueForPDCPSN18ExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("COUNTValueForPDCP-SN18-ExtIEsEntry", v)
}

// DecodeCOUNTValueForPDCPSN18ExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCOUNTValueForPDCPSN18ExtIEsEntry(data []byte) (COUNTValueForPDCPSN18ExtIEsEntry, int, error) {
	var v COUNTValueForPDCPSN18ExtIEsEntry
	err := ngapcodec.DecodeParameter("COUNTValueForPDCP-SN18-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return COUNTValueForPDCPSN18ExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
