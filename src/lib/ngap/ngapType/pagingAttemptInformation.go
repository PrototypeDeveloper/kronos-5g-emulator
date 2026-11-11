// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PagingAttemptInformation represents ASN.1 PagingAttemptInformation.
type PagingAttemptInformation struct {
	PagingAttemptCount             PagingAttemptCount
	IntendedNumberOfPagingAttempts IntendedNumberOfPagingAttempts
	HasNextPagingAreaScope         bool
	NextPagingAreaScope            NextPagingAreaScope
	HasIEExtensions                bool
	IEExtensions                   []PagingAttemptInformationExtIEsEntry
	UnknownExtensions              []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PagingAttemptInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PagingAttemptInformation", v)
}

// DecodePagingAttemptInformation decodes a standalone APER value and reports the consumed octets.
func DecodePagingAttemptInformation(data []byte) (PagingAttemptInformation, int, error) {
	var v PagingAttemptInformation
	err := ngapcodec.DecodeParameter("PagingAttemptInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return PagingAttemptInformation{}, 0, err
	}
	return v, len(data), err
}
