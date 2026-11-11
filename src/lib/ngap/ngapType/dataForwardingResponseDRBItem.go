// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DataForwardingResponseDRBItem represents ASN.1 DataForwardingResponseDRBItem.
type DataForwardingResponseDRBItem struct {
	DRBID                           DRBID
	HasDLForwardingUPTNLInformation bool
	DLForwardingUPTNLInformation    UPTransportLayerInformation
	HasULForwardingUPTNLInformation bool
	ULForwardingUPTNLInformation    UPTransportLayerInformation
	HasIEExtensions                 bool
	IEExtensions                    []DataForwardingResponseDRBItemExtIEsEntry
	UnknownExtensions               []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DataForwardingResponseDRBItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DataForwardingResponseDRBItem", v)
}

// DecodeDataForwardingResponseDRBItem decodes a standalone APER value and reports the consumed octets.
func DecodeDataForwardingResponseDRBItem(data []byte) (DataForwardingResponseDRBItem, int, error) {
	var v DataForwardingResponseDRBItem
	err := ngapcodec.DecodeParameter("DataForwardingResponseDRBItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return DataForwardingResponseDRBItem{}, 0, err
	}
	return v, len(data), err
}
