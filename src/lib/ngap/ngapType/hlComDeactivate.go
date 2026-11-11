// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// HLComDeactivate represents ASN.1 HLComDeactivate.
type HLComDeactivate struct {
	UEReachabilityIndication HLComDeactivateUEReachabilityIndication `aper:"valueExt,valueLB:0,valueUB:0"`
	HasIEExtensions          bool
	IEExtensions             []HLComDeactivateExtIEsEntry
	UnknownExtensions        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HLComDeactivate) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("HLComDeactivate", v)
}

// DecodeHLComDeactivate decodes a standalone APER value and reports the consumed octets.
func DecodeHLComDeactivate(data []byte) (HLComDeactivate, int, error) {
	var v HLComDeactivate
	err := ngapcodec.DecodeParameter("HLComDeactivate", data, &v)
	if ngapcodec.IsFatal(err) {
		return HLComDeactivate{}, 0, err
	}
	return v, len(data), err
}
