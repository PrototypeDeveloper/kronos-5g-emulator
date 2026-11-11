// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TraceActivationExtIEsEntryPresent uint32

const (
	TraceActivationExtIEsEntryPresentNothing TraceActivationExtIEsEntryPresent = iota
	TraceActivationExtIEsEntryPresentMDTConfiguration
	TraceActivationExtIEsEntryPresentTraceCollectionEntityURI
	TraceActivationExtIEsEntryPresentUnknown
)

type TraceActivationExtIEsEntry struct {
	ID                       ProtocolExtensionID
	Criticality              Criticality
	Present                  TraceActivationExtIEsEntryPresent
	MDTConfiguration         MDTConfiguration `ngap:"id:255,criticality:ignore,presence:optional"`
	TraceCollectionEntityURI URIAddress       `ngap:"id:257,criticality:ignore,presence:optional"`
	UnknownValue             []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TraceActivationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TraceActivation-ExtIEsEntry", v)
}

// DecodeTraceActivationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTraceActivationExtIEsEntry(data []byte) (TraceActivationExtIEsEntry, int, error) {
	var v TraceActivationExtIEsEntry
	err := ngapcodec.DecodeParameter("TraceActivation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TraceActivationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
