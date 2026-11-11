// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EUTRANCellReportItem represents ASN.1 EUTRAN-CellReportItem.
type EUTRANCellReportItem struct {
	ECGI                                  EUTRACGI
	EUTRANCompositeAvailableCapacityGroup EUTRANCompositeAvailableCapacityGroup
	HasEUTRANNumberOfActiveUEs            bool
	EUTRANNumberOfActiveUEs               EUTRANNumberOfActiveUEs
	HasEUTRANNoofRRCConnections           bool
	EUTRANNoofRRCConnections              NGRANNoofRRCConnections
	HasEUTRANRadioResourceStatus          bool
	EUTRANRadioResourceStatus             EUTRANRadioResourceStatus
	HasIEExtensions                       bool
	IEExtensions                          []EUTRANCellReportItemExtIEsEntry
	UnknownExtensions                     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRANCellReportItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRAN-CellReportItem", v)
}

// DecodeEUTRANCellReportItem decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRANCellReportItem(data []byte) (EUTRANCellReportItem, int, error) {
	var v EUTRANCellReportItem
	err := ngapcodec.DecodeParameter("EUTRAN-CellReportItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRANCellReportItem{}, 0, err
	}
	return v, len(data), err
}
