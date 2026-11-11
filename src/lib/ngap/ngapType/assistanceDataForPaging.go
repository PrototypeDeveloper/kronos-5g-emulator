// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AssistanceDataForPaging represents ASN.1 AssistanceDataForPaging.
type AssistanceDataForPaging struct {
	HasAssistanceDataForRecommendedCells bool
	AssistanceDataForRecommendedCells    AssistanceDataForRecommendedCells
	HasPagingAttemptInformation          bool
	PagingAttemptInformation             PagingAttemptInformation
	HasIEExtensions                      bool
	IEExtensions                         []AssistanceDataForPagingExtIEsEntry
	UnknownExtensions                    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AssistanceDataForPaging) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AssistanceDataForPaging", v)
}

// DecodeAssistanceDataForPaging decodes a standalone APER value and reports the consumed octets.
func DecodeAssistanceDataForPaging(data []byte) (AssistanceDataForPaging, int, error) {
	var v AssistanceDataForPaging
	err := ngapcodec.DecodeParameter("AssistanceDataForPaging", data, &v)
	if ngapcodec.IsFatal(err) {
		return AssistanceDataForPaging{}, 0, err
	}
	return v, len(data), err
}
