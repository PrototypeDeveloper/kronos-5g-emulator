// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PNINPNAreaScopeofMDTExtIEsEntryPresent uint32

const (
	PNINPNAreaScopeofMDTExtIEsEntryPresentNothing PNINPNAreaScopeofMDTExtIEsEntryPresent = iota
	PNINPNAreaScopeofMDTExtIEsEntryPresentUnknown
)

type PNINPNAreaScopeofMDTExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PNINPNAreaScopeofMDTExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PNINPNAreaScopeofMDTExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PNI-NPN-AreaScopeofMDT-ExtIEsEntry", v)
}

// DecodePNINPNAreaScopeofMDTExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePNINPNAreaScopeofMDTExtIEsEntry(data []byte) (PNINPNAreaScopeofMDTExtIEsEntry, int, error) {
	var v PNINPNAreaScopeofMDTExtIEsEntry
	err := ngapcodec.DecodeParameter("PNI-NPN-AreaScopeofMDT-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PNINPNAreaScopeofMDTExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
