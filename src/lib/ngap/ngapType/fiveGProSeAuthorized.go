// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveGProSeAuthorized represents ASN.1 FiveG-ProSeAuthorized.
type FiveGProSeAuthorized struct {
	HasFiveGProSeDirectDiscovery        bool
	FiveGProSeDirectDiscovery           FiveGProSeDirectDiscovery
	HasFiveGProSeDirectCommunication    bool
	FiveGProSeDirectCommunication       FiveGProSeDirectCommunication
	HasFiveGProSeLayer2UEtoNetworkRelay bool
	FiveGProSeLayer2UEtoNetworkRelay    FiveGProSeLayer2UEtoNetworkRelay
	HasFiveGProSeLayer3UEtoNetworkRelay bool
	FiveGProSeLayer3UEtoNetworkRelay    FiveGProSeLayer3UEtoNetworkRelay
	HasFiveGProSeLayer2RemoteUE         bool
	FiveGProSeLayer2RemoteUE            FiveGProSeLayer2RemoteUE
	HasIEExtensions                     bool
	IEExtensions                        []FiveGProSeAuthorizedExtIEsEntry
	UnknownExtensions                   []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSeAuthorized) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveG-ProSeAuthorized", v)
}

// DecodeFiveGProSeAuthorized decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSeAuthorized(data []byte) (FiveGProSeAuthorized, int, error) {
	var v FiveGProSeAuthorized
	err := ngapcodec.DecodeParameter("FiveG-ProSeAuthorized", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSeAuthorized{}, 0, err
	}
	return v, len(data), err
}
