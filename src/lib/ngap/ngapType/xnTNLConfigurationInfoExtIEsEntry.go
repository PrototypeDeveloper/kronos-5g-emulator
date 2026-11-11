// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type XnTNLConfigurationInfoExtIEsEntryPresent uint32

const (
	XnTNLConfigurationInfoExtIEsEntryPresentNothing XnTNLConfigurationInfoExtIEsEntryPresent = iota
	XnTNLConfigurationInfoExtIEsEntryPresentUnknown
)

type XnTNLConfigurationInfoExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      XnTNLConfigurationInfoExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v XnTNLConfigurationInfoExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("XnTNLConfigurationInfo-ExtIEsEntry", v)
}

// DecodeXnTNLConfigurationInfoExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeXnTNLConfigurationInfoExtIEsEntry(data []byte) (XnTNLConfigurationInfoExtIEsEntry, int, error) {
	var v XnTNLConfigurationInfoExtIEsEntry
	err := ngapcodec.DecodeParameter("XnTNLConfigurationInfo-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return XnTNLConfigurationInfoExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
