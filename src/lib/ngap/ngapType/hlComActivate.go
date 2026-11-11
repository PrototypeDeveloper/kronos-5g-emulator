// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// HLComActivate represents ASN.1 HLComActivate.
type HLComActivate struct {
	NRPagingLongEDRXInformationForRRCINACTIVE NRPagingLongEDRXInformationForRRCINACTIVE
	HasIEExtensions                           bool
	IEExtensions                              []HLComActivateExtIEsEntry
	UnknownExtensions                         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HLComActivate) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("HLComActivate", v) }

// DecodeHLComActivate decodes a standalone APER value and reports the consumed octets.
func DecodeHLComActivate(data []byte) (HLComActivate, int, error) {
	var v HLComActivate
	err := ngapcodec.DecodeParameter("HLComActivate", data, &v)
	if ngapcodec.IsFatal(err) {
		return HLComActivate{}, 0, err
	}
	return v, len(data), err
}
