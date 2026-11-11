// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SharedNGUMulticastTNLInformation represents ASN.1 SharedNGU-MulticastTNLInformation.
type SharedNGUMulticastTNLInformation struct {
	IPMulticastAddress TransportLayerAddress
	IPSourceAddress    TransportLayerAddress
	GTPTEID            GTPTEID
	HasIEExtensions    bool
	IEExtensions       []SharedNGUMulticastTNLInformationExtIEsEntry
	UnknownExtensions  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SharedNGUMulticastTNLInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SharedNGU-MulticastTNLInformation", v)
}

// DecodeSharedNGUMulticastTNLInformation decodes a standalone APER value and reports the consumed octets.
func DecodeSharedNGUMulticastTNLInformation(data []byte) (SharedNGUMulticastTNLInformation, int, error) {
	var v SharedNGUMulticastTNLInformation
	err := ngapcodec.DecodeParameter("SharedNGU-MulticastTNLInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return SharedNGUMulticastTNLInformation{}, 0, err
	}
	return v, len(data), err
}
