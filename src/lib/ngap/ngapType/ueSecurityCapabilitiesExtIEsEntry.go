// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UESecurityCapabilitiesExtIEsEntryPresent uint32

const (
	UESecurityCapabilitiesExtIEsEntryPresentNothing UESecurityCapabilitiesExtIEsEntryPresent = iota
	UESecurityCapabilitiesExtIEsEntryPresentUnknown
)

type UESecurityCapabilitiesExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      UESecurityCapabilitiesExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UESecurityCapabilitiesExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UESecurityCapabilities-ExtIEsEntry", v)
}

// DecodeUESecurityCapabilitiesExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUESecurityCapabilitiesExtIEsEntry(data []byte) (UESecurityCapabilitiesExtIEsEntry, int, error) {
	var v UESecurityCapabilitiesExtIEsEntry
	err := ngapcodec.DecodeParameter("UESecurityCapabilities-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UESecurityCapabilitiesExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
