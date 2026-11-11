// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SharedNGUMulticastTNLInformationExtIEsEntryPresent uint32

const (
	SharedNGUMulticastTNLInformationExtIEsEntryPresentNothing SharedNGUMulticastTNLInformationExtIEsEntryPresent = iota
	SharedNGUMulticastTNLInformationExtIEsEntryPresentUnknown
)

type SharedNGUMulticastTNLInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SharedNGUMulticastTNLInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SharedNGUMulticastTNLInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SharedNGU-MulticastTNLInformation-ExtIEsEntry", v)
}

// DecodeSharedNGUMulticastTNLInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSharedNGUMulticastTNLInformationExtIEsEntry(data []byte) (SharedNGUMulticastTNLInformationExtIEsEntry, int, error) {
	var v SharedNGUMulticastTNLInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("SharedNGU-MulticastTNLInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SharedNGUMulticastTNLInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
