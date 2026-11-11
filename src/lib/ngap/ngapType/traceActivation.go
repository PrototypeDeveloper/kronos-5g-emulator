// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TraceActivation represents ASN.1 TraceActivation.
type TraceActivation struct {
	NGRANTraceID                   NGRANTraceID
	InterfacesToTrace              InterfacesToTrace
	TraceDepth                     TraceDepth
	TraceCollectionEntityIPAddress TransportLayerAddress
	HasIEExtensions                bool
	IEExtensions                   []TraceActivationExtIEsEntry
	UnknownExtensions              []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TraceActivation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TraceActivation", v)
}

// DecodeTraceActivation decodes a standalone APER value and reports the consumed octets.
func DecodeTraceActivation(data []byte) (TraceActivation, int, error) {
	var v TraceActivation
	err := ngapcodec.DecodeParameter("TraceActivation", data, &v)
	if ngapcodec.IsFatal(err) {
		return TraceActivation{}, 0, err
	}
	return v, len(data), err
}
