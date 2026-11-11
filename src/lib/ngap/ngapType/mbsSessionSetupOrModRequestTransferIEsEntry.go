// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSSessionSetupOrModRequestTransferIEsEntryPresent uint32

const (
	MBSSessionSetupOrModRequestTransferIEsEntryPresentNothing MBSSessionSetupOrModRequestTransferIEsEntryPresent = iota
	MBSSessionSetupOrModRequestTransferIEsEntryPresentMBSSessionTNLInfo5GC
	MBSSessionSetupOrModRequestTransferIEsEntryPresentMBSQoSFlowsToBeSetupModList
	MBSSessionSetupOrModRequestTransferIEsEntryPresentMBSSessionFSAIDList
	MBSSessionSetupOrModRequestTransferIEsEntryPresentSupportedUETypeList
	MBSSessionSetupOrModRequestTransferIEsEntryPresentMBSNGUFailureIndication
	MBSSessionSetupOrModRequestTransferIEsEntryPresentUnknown
)

type MBSSessionSetupOrModRequestTransferIEsEntry struct {
	ID                          ProtocolIEID
	Criticality                 Criticality
	Present                     MBSSessionSetupOrModRequestTransferIEsEntryPresent
	MBSSessionTNLInfo5GC        MBSSessionTNLInfo5GC     `ngap:"id:352,criticality:reject,presence:optional"`
	MBSQoSFlowsToBeSetupModList MBSQoSFlowsToBeSetupList `ngap:"id:297,criticality:reject,presence:mandatory"`
	MBSSessionFSAIDList         MBSSessionFSAIDList      `ngap:"id:357,criticality:ignore,presence:optional"`
	SupportedUETypeList         SupportedUETypeList      `ngap:"id:372,criticality:ignore,presence:optional"`
	MBSNGUFailureIndication     MBSNGUFailureIndication  `ngap:"id:434,criticality:reject,presence:optional"`
	UnknownValue                []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionSetupOrModRequestTransferIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionSetupOrModRequestTransferIEsEntry", v)
}

// DecodeMBSSessionSetupOrModRequestTransferIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionSetupOrModRequestTransferIEsEntry(data []byte) (MBSSessionSetupOrModRequestTransferIEsEntry, int, error) {
	var v MBSSessionSetupOrModRequestTransferIEsEntry
	err := ngapcodec.DecodeParameter("MBSSessionSetupOrModRequestTransferIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionSetupOrModRequestTransferIEsEntry{}, 0, err
	}
	return v, len(data), err
}
