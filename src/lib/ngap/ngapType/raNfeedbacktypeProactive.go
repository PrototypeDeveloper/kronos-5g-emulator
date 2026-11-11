// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RANfeedbacktypeProactive represents ASN.1 RANfeedbacktype-proactive.
type RANfeedbacktypeProactive struct {
	BurstArrivalTimeWindow BurstArrivalTimeWindow
	HasPeriodicityRange    bool
	PeriodicityRange       PeriodicityRange
	HasIEExtension         bool
	IEExtension            []RANfeedbacktypeProactiveExtIEsEntry
	UnknownExtensions      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANfeedbacktypeProactive) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RANfeedbacktype-proactive", v)
}

// DecodeRANfeedbacktypeProactive decodes a standalone APER value and reports the consumed octets.
func DecodeRANfeedbacktypeProactive(data []byte) (RANfeedbacktypeProactive, int, error) {
	var v RANfeedbacktypeProactive
	err := ngapcodec.DecodeParameter("RANfeedbacktype-proactive", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANfeedbacktypeProactive{}, 0, err
	}
	return v, len(data), err
}
