// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NGRANCellReportItem represents ASN.1 NGRAN-CellReportItem.
type NGRANCellReportItem struct {
	NGRANCGI                             NGRANCGI
	NGRANCompositeAvailableCapacityGroup EUTRANCompositeAvailableCapacityGroup
	HasNGRANNumberOfActiveUEs            bool
	NGRANNumberOfActiveUEs               NGRANNumberOfActiveUEs
	HasNGRANNoofRRCConnections           bool
	NGRANNoofRRCConnections              NGRANNoofRRCConnections
	HasNGRANRadioResourceStatus          bool
	NGRANRadioResourceStatus             NGRANRadioResourceStatus
	HasIEExtensions                      bool
	IEExtensions                         []NGRANCellReportItemExtIEsEntry
	UnknownExtensions                    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGRANCellReportItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGRAN-CellReportItem", v)
}

// DecodeNGRANCellReportItem decodes a standalone APER value and reports the consumed octets.
func DecodeNGRANCellReportItem(data []byte) (NGRANCellReportItem, int, error) {
	var v NGRANCellReportItem
	err := ngapcodec.DecodeParameter("NGRAN-CellReportItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGRANCellReportItem{}, 0, err
	}
	return v, len(data), err
}
