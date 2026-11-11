// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerExtIEsEntryPresent uint32

const (
	TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerExtIEsEntryPresentNothing TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerExtIEsEntryPresent = iota
	TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerExtIEsEntryPresentNGAPIESupportInformationResponseList
	TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerExtIEsEntryPresentUnknown
)

type TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerExtIEsEntry struct {
	ID                                   ProtocolExtensionID
	Criticality                          Criticality
	Present                              TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerExtIEsEntryPresent
	NGAPIESupportInformationResponseList NGAPIESupportInformationResponseList `ngap:"id:356,criticality:ignore,presence:optional"`
	UnknownValue                         []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargetNGRANNode-ToSourceNGRANNode-FailureTransparentContainer-ExtIEsEntry", v)
}

// DecodeTargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerExtIEsEntry(data []byte) (TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerExtIEsEntry, int, error) {
	var v TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerExtIEsEntry
	err := ngapcodec.DecodeParameter("TargetNGRANNode-ToSourceNGRANNode-FailureTransparentContainer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
