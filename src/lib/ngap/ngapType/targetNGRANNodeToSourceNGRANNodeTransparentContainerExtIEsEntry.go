// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsEntryPresent uint32

const (
	TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsEntryPresentNothing TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsEntryPresent = iota
	TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsEntryPresentDAPSResponseInfoList
	TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsEntryPresentDirectForwardingPathAvailability
	TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsEntryPresentMBSActiveSessionInformationTargettoSourceList
	TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsEntryPresentNGAPIESupportInformationResponseList
	TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsEntryPresentUnknown
)

type TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsEntry struct {
	ID                                            ProtocolExtensionID
	Criticality                                   Criticality
	Present                                       TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsEntryPresent
	DAPSResponseInfoList                          DAPSResponseInfoList                          `ngap:"id:267,criticality:ignore,presence:optional"`
	DirectForwardingPathAvailability              DirectForwardingPathAvailability              `ngap:"id:22,criticality:ignore,presence:optional"`
	MBSActiveSessionInformationTargettoSourceList MBSActiveSessionInformationTargettoSourceList `ngap:"id:324,criticality:ignore,presence:optional"`
	NGAPIESupportInformationResponseList          NGAPIESupportInformationResponseList          `ngap:"id:356,criticality:ignore,presence:optional"`
	UnknownValue                                  []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargetNGRANNode-ToSourceNGRANNode-TransparentContainer-ExtIEsEntry", v)
}

// DecodeTargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsEntry(data []byte) (TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsEntry, int, error) {
	var v TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsEntry
	err := ngapcodec.DecodeParameter("TargetNGRANNode-ToSourceNGRANNode-TransparentContainer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
