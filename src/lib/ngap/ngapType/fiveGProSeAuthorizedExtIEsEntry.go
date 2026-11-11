// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type FiveGProSeAuthorizedExtIEsEntryPresent uint32

const (
	FiveGProSeAuthorizedExtIEsEntryPresentNothing FiveGProSeAuthorizedExtIEsEntryPresent = iota
	FiveGProSeAuthorizedExtIEsEntryPresentFiveGProSeLayer2Multipath
	FiveGProSeAuthorizedExtIEsEntryPresentFiveGProSeLayer2UEtoUERelay
	FiveGProSeAuthorizedExtIEsEntryPresentFiveGProSeLayer2UEtoUERemote
	FiveGProSeAuthorizedExtIEsEntryPresentFiveGProSeLayer2MHUEtoNetworkRelay
	FiveGProSeAuthorizedExtIEsEntryPresentFiveGProSeLayer2MHIntermediateUEtoNetworkRelay
	FiveGProSeAuthorizedExtIEsEntryPresentFiveGProSeLayer2MHRemote
	FiveGProSeAuthorizedExtIEsEntryPresentFiveGProSeLayer3MHUEtoNetworkRelay
	FiveGProSeAuthorizedExtIEsEntryPresentUnknown
)

type FiveGProSeAuthorizedExtIEsEntry struct {
	ID                                             ProtocolExtensionID
	Criticality                                    Criticality
	Present                                        FiveGProSeAuthorizedExtIEsEntryPresent
	FiveGProSeLayer2Multipath                      FiveGProSeLayer2Multipath                      `ngap:"id:379,criticality:ignore,presence:optional"`
	FiveGProSeLayer2UEtoUERelay                    FiveGProSeLayer2UEtoUERelay                    `ngap:"id:380,criticality:ignore,presence:optional"`
	FiveGProSeLayer2UEtoUERemote                   FiveGProSeLayer2UEtoUERemote                   `ngap:"id:381,criticality:ignore,presence:optional"`
	FiveGProSeLayer2MHUEtoNetworkRelay             FiveGProSeLayer2MHUEtoNetworkRelay             `ngap:"id:479,criticality:ignore,presence:optional"`
	FiveGProSeLayer2MHIntermediateUEtoNetworkRelay FiveGProSeLayer2MHIntermediateUEtoNetworkRelay `ngap:"id:480,criticality:ignore,presence:optional"`
	FiveGProSeLayer2MHRemote                       FiveGProSeLayer2MHRemote                       `ngap:"id:481,criticality:ignore,presence:optional"`
	FiveGProSeLayer3MHUEtoNetworkRelay             FiveGProSeLayer3MHUEtoNetworkRelay             `ngap:"id:482,criticality:ignore,presence:optional"`
	UnknownValue                                   []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSeAuthorizedExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveG-ProSeAuthorized-ExtIEsEntry", v)
}

// DecodeFiveGProSeAuthorizedExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSeAuthorizedExtIEsEntry(data []byte) (FiveGProSeAuthorizedExtIEsEntry, int, error) {
	var v FiveGProSeAuthorizedExtIEsEntry
	err := ngapcodec.DecodeParameter("FiveG-ProSeAuthorized-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSeAuthorizedExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
