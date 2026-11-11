// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntryPresent uint32

const (
	SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntryPresentNothing SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntryPresent = iota
	SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntryPresentSgNBUEX2APID
	SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntryPresentUEHistoryInformationFromTheUE
	SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntryPresentSourceNodeID
	SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntryPresentUEContextReferenceAtSource
	SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntryPresentMBSActiveSessionInformationSourcetoTargetList
	SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntryPresentQMCConfigInfo
	SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntryPresentNGAPIESupportInformationRequestList
	SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntryPresentCandidateRelayUEInformationList
	SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntryPresentTimeBasedHandoverInformation
	SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntryPresentSourceSNToTargetSNQMCInfo
	SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntryPresentUnknown
)

type SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntry struct {
	ID                                            ProtocolExtensionID
	Criticality                                   Criticality
	Present                                       SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntryPresent
	SgNBUEX2APID                                  SgNBUEX2APID                                  `ngap:"id:182,criticality:ignore,presence:optional"`
	UEHistoryInformationFromTheUE                 UEHistoryInformationFromTheUE                 `ngap:"id:253,criticality:ignore,presence:optional"`
	SourceNodeID                                  SourceNodeID                                  `ngap:"id:286,criticality:ignore,presence:optional"`
	UEContextReferenceAtSource                    RANUENGAPID                                   `ngap:"id:288,criticality:ignore,presence:optional"`
	MBSActiveSessionInformationSourcetoTargetList MBSActiveSessionInformationSourcetoTargetList `ngap:"id:323,criticality:ignore,presence:optional"`
	QMCConfigInfo                                 QMCConfigInfo                                 `ngap:"id:328,criticality:ignore,presence:optional"`
	NGAPIESupportInformationRequestList           NGAPIESupportInformationRequestList           `ngap:"id:355,criticality:ignore,presence:optional"`
	CandidateRelayUEInformationList               CandidateRelayUEInformationList               `ngap:"id:382,criticality:reject,presence:optional"`
	TimeBasedHandoverInformation                  TimeBasedHandoverInformation                  `ngap:"id:420,criticality:ignore,presence:optional"`
	SourceSNToTargetSNQMCInfo                     QMCConfigInfo                                 `ngap:"id:437,criticality:ignore,presence:optional"`
	UnknownValue                                  []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SourceNGRANNode-ToTargetNGRANNode-TransparentContainer-ExtIEsEntry", v)
}

// DecodeSourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntry(data []byte) (SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntry, int, error) {
	var v SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntry
	err := ngapcodec.DecodeParameter("SourceNGRANNode-ToTargetNGRANNode-TransparentContainer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
