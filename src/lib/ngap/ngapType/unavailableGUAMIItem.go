// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UnavailableGUAMIItem represents ASN.1 UnavailableGUAMIItem.
type UnavailableGUAMIItem struct {
	GUAMI                           GUAMI
	HasTimerApproachForGUAMIRemoval bool
	TimerApproachForGUAMIRemoval    TimerApproachForGUAMIRemoval
	HasBackupAMFName                bool
	BackupAMFName                   AMFName
	HasIEExtensions                 bool
	IEExtensions                    []UnavailableGUAMIItemExtIEsEntry
	UnknownExtensions               []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UnavailableGUAMIItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UnavailableGUAMIItem", v)
}

// DecodeUnavailableGUAMIItem decodes a standalone APER value and reports the consumed octets.
func DecodeUnavailableGUAMIItem(data []byte) (UnavailableGUAMIItem, int, error) {
	var v UnavailableGUAMIItem
	err := ngapcodec.DecodeParameter("UnavailableGUAMIItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return UnavailableGUAMIItem{}, 0, err
	}
	return v, len(data), err
}
