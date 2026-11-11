// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MulticastSessionUpdateRequestTransferIEsEntryPresent uint32

const (
	MulticastSessionUpdateRequestTransferIEsEntryPresentNothing MulticastSessionUpdateRequestTransferIEsEntryPresent = iota
	MulticastSessionUpdateRequestTransferIEsEntryPresentMBSSessionID
	MulticastSessionUpdateRequestTransferIEsEntryPresentMBSServiceArea
	MulticastSessionUpdateRequestTransferIEsEntryPresentMBSQoSFlowsToBeSetupModList
	MulticastSessionUpdateRequestTransferIEsEntryPresentMBSQoSFlowToReleaseList
	MulticastSessionUpdateRequestTransferIEsEntryPresentMBSSessionTNLInfo5GC
	MulticastSessionUpdateRequestTransferIEsEntryPresentMBSUPFailureIndication
	MulticastSessionUpdateRequestTransferIEsEntryPresentUnknown
)

type MulticastSessionUpdateRequestTransferIEsEntry struct {
	ID                          ProtocolIEID
	Criticality                 Criticality
	Present                     MulticastSessionUpdateRequestTransferIEsEntryPresent
	MBSSessionID                MBSSessionID             `ngap:"id:299,criticality:reject,presence:mandatory"`
	MBSServiceArea              MBSServiceArea           `ngap:"id:298,criticality:reject,presence:optional"`
	MBSQoSFlowsToBeSetupModList MBSQoSFlowsToBeSetupList `ngap:"id:297,criticality:reject,presence:optional"`
	MBSQoSFlowToReleaseList     QosFlowListWithCause     `ngap:"id:351,criticality:reject,presence:optional"`
	MBSSessionTNLInfo5GC        MBSSessionTNLInfo5GC     `ngap:"id:352,criticality:reject,presence:optional"`
	MBSUPFailureIndication      MBSUPFailureIndication   `ngap:"id:444,criticality:reject,presence:optional"`
	UnknownValue                []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MulticastSessionUpdateRequestTransferIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MulticastSessionUpdateRequestTransferIEsEntry", v)
}

// DecodeMulticastSessionUpdateRequestTransferIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMulticastSessionUpdateRequestTransferIEsEntry(data []byte) (MulticastSessionUpdateRequestTransferIEsEntry, int, error) {
	var v MulticastSessionUpdateRequestTransferIEsEntry
	err := ngapcodec.DecodeParameter("MulticastSessionUpdateRequestTransferIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MulticastSessionUpdateRequestTransferIEsEntry{}, 0, err
	}
	return v, len(data), err
}
