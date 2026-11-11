package scenario

import (
	"encoding/binary"

	"libngap/ngapMessage"
	"libngap/ngapType"
	"libngap/ngapcodec"

	"kronos/models"
)

// HandoverULTEIDs returns the uplink TEIDs of the sessions of a
// HandoverRequest (from each Handover Request Transfer), which tell the
// subscriber it is for.
func HandoverULTEIDs(msg any) []uint32 {
	m, ok := msg.(*ngapMessage.HandoverRequest)
	if !ok {
		return nil
	}
	var out []uint32
	for _, it := range m.PDUSessionResourceSetupListHOReq.List {
		var tr ngapType.PDUSessionResourceSetupRequestTransfer
		if err := ngapcodec.DecodeParameter("PDUSessionResourceSetupRequestTransfer", it.HandoverRequestTransfer, &tr); ngapcodec.IsFatal(err) {
			continue
		}
		for _, ie := range tr.ProtocolIEs {
			if ie.Present != ngapType.PDUSessionResourceSetupRequestTransferIEsEntryPresentULNGUUPTNLInformation ||
				ie.ULNGUUPTNLInformation.Present != ngapType.UPTransportLayerInformationPresentGTPTunnel {
				continue
			}
			if teid := ie.ULNGUUPTNLInformation.GTPTunnel.GTPTEID.Value; len(teid) == models.TEIDLen {
				out = append(out, binary.BigEndian.Uint32(teid))
			}
		}
	}
	return out
}
