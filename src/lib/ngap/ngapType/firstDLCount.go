// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FirstDLCount represents ASN.1 FirstDLCount.
type FirstDLCount struct {
	DRBsSubjectToEarlyStatusTransfer DRBsSubjectToEarlyStatusTransferList
	HasIEExtension                   bool
	IEExtension                      []FirstDLCountExtIEsEntry
	UnknownExtensions                []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FirstDLCount) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("FirstDLCount", v) }

// DecodeFirstDLCount decodes a standalone APER value and reports the consumed octets.
func DecodeFirstDLCount(data []byte) (FirstDLCount, int, error) {
	var v FirstDLCount
	err := ngapcodec.DecodeParameter("FirstDLCount", data, &v)
	if ngapcodec.IsFatal(err) {
		return FirstDLCount{}, 0, err
	}
	return v, len(data), err
}
