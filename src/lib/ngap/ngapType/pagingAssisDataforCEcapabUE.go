// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PagingAssisDataforCEcapabUE represents ASN.1 PagingAssisDataforCEcapabUE.
type PagingAssisDataforCEcapabUE struct {
	EUTRACGI                 EUTRACGI
	CoverageEnhancementLevel CoverageEnhancementLevel
	HasIEExtensions          bool
	IEExtensions             []PagingAssisDataforCEcapabUEExtIEsEntry
	UnknownExtensions        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PagingAssisDataforCEcapabUE) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PagingAssisDataforCEcapabUE", v)
}

// DecodePagingAssisDataforCEcapabUE decodes a standalone APER value and reports the consumed octets.
func DecodePagingAssisDataforCEcapabUE(data []byte) (PagingAssisDataforCEcapabUE, int, error) {
	var v PagingAssisDataforCEcapabUE
	err := ngapcodec.DecodeParameter("PagingAssisDataforCEcapabUE", data, &v)
	if ngapcodec.IsFatal(err) {
		return PagingAssisDataforCEcapabUE{}, 0, err
	}
	return v, len(data), err
}
