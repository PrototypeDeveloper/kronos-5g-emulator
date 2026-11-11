package gnb

import (
	"strings"

	"kronos/scenario"
)

// AutoField implements role.Role: the fields of a template that the gNB's
// auto-fix (the mapping) replaces when it is sent, which the signal editor
// marks, and notes on fields it changes in part or only sometimes. It
// follows build, patchNAS and the session patches; raw sends the template
// as it is, and nas=keep keeps the NAS but its security header.
func (r *Role) AutoField(f scenario.FieldRef) (bool, string) {
	switch {
	case f.Mode == scenario.ModeRaw:
		return false, ""
	case f.NAS == "":
		return r.autoNGAP(f.NGAP, f.Path)
	case strings.HasPrefix(f.Path, "@"):
		return true, "set when the NAS message is protected"
	case f.Mode == scenario.ModeNASKeep:
		return false, ""
	}
	return r.autoNAS(f.NAS, f.Outer, f.Path)
}

// under reports whether path is prefix or a field under it; "*" in prefix
// matches one segment.
func under(path, prefix string) bool {
	ps, qs := strings.Split(path, "/"), strings.Split(prefix, "/")
	if len(ps) < len(qs) {
		return false
	}
	for i, q := range qs {
		if q != "*" && q != ps[i] {
			return false
		}
	}
	return true
}

// autoNGAP is AutoField for a field of the NGAP message msg.
func (r *Role) autoNGAP(msg, path string) (bool, string) {
	// Fields of any UE-associated message, then those of particular messages.
	switch {
	case under(path, "RANUENGAPID"), under(path, "AMFUENGAPID"):
		return true, "the NGAP ID of the UE association"
	case under(path, "UserLocationInformation"):
		return true, "the NR cell and TA of the config (a TimeStamp: the time of sending)"
	}
	switch msg {
	case "InitialUEMessage":
		if under(path, "FiveGSTMSI") {
			return true, "from the 5G-GUTI (left out without one)"
		}
	case "NGSetupRequest":
		switch {
		case under(path, "GlobalRANNodeID"):
			return true, "gnodeb.plmn and gnodeb.gnodebid"
		case under(path, "SupportedTAList"):
			return true, "gnodeb.tac, gnodeb.plmn and gnodeb.snssai (the first items are the pattern)"
		case under(path, "RANNodeName") && r.name != "":
			return true, "gnodeb.name"
		}
	case "PDUSessionResourceSetupResponse", "InitialContextSetupResponse":
		list, failed := "PDUSessionResourceSetupListSURes", "PDUSessionResourceFailedToSetupListSURes"
		if msg == "InitialContextSetupResponse" {
			list, failed = "PDUSessionResourceSetupListCxtRes", "PDUSessionResourceFailedToSetupListCxtRes"
		}
		const tr = "/List/*/PDUSessionResourceSetupResponseTransfer/>/"
		switch {
		case path == list:
			return false, "one item per PDU session answered; the first item is the pattern"
		case under(path, list+"/List/*/PDUSessionID"):
			return true, "the PDU session answered"
		case under(path, list+tr+"DLQosFlowPerTNLInformation"):
			return true, "gtpu.src, the DL TEID and the QoS flows of the session"
		case under(path, list+tr+"AdditionalDLQosFlowPerTNLInformation"), under(path, list+tr+"QosFlowFailedToSetupList"), under(path, failed):
			return true, "left out when sent"
		}
	case "PDUSessionResourceReleaseResponse":
		const list = "PDUSessionResourceReleasedListRelRes"
		switch {
		case path == list:
			return false, "one item per PDU session released; the first item is the pattern"
		case under(path, list+"/List/*/PDUSessionID"):
			return true, "the PDU session released"
		}
	case "UEContextReleaseRequest":
		if under(path, "PDUSessionResourceListCxtRelReq") {
			return true, "the established PDU sessions"
		}
	case "UEContextReleaseComplete":
		if under(path, "PDUSessionResourceListCxtRelCpl") {
			return true, "the established PDU sessions"
		}
	}
	return false, ""
}

// autoNAS is AutoField for a field of the NAS message nas; outer is the
// message whose container holds it ("" if none).
func (r *Role) autoNAS(nas, outer, path string) (bool, string) {
	status := under(path, "PDUSessionStatus") || under(path, "UplinkDataStatus") || under(path, "AllowedPDUSessionStatus")
	switch nas {
	case "RegistrationRequest":
		inContainer := outer != ""
		switch {
		case path == "RegistrationTypeAndNgKSI" && !inContainer:
			return false, "the ngKSI (high 4 bits) is replaced"
		case under(path, "MobileIdentity"):
			return true, "the SUCI of the IMSI, or the current 5G-GUTI (by the identity type of the template)"
		case under(path, "UESecurityCapability"):
			return false, "the first two octets (5G-EA, 5G-IA) come from ue.ciphering and ue.integrity"
		case under(path, "RequestedNSSAI") && r.nssai != nil:
			return true, "ue.nssai"
		case status:
			return true, "the established PDU sessions"
		case under(path, "NASMessageContainer") && !inContainer:
			return true, "rebuilt from this message when sent under a security context"
		}
	case "ServiceRequest":
		switch {
		case path == "ServiceTypeAndNgKSI":
			return false, "the ngKSI (low 4 bits) is replaced"
		case under(path, "FiveGSTMSI"):
			return true, "the 5G-S-TMSI of the 5G-GUTI"
		case status:
			return true, "the established PDU sessions"
		}
	case "DeregistrationRequestUE":
		switch {
		case path == "DeregistrationTypeAndNgKSI":
			return false, "the ngKSI (high 4 bits) is replaced"
		case under(path, "MobileIdentity"):
			return true, "the current 5G-GUTI, or the SUCI of the IMSI (by the identity type of the template)"
		}
	case "AuthenticationResponse":
		if under(path, "AuthenticationResponseParameter") {
			return true, "the RES* computed from the AuthenticationRequest"
		}
	case "SecurityModeComplete":
		if under(path, "IMEISV") {
			return false, "replaced by ue.imeisv when the Security Mode Command asks for it"
		}
	case "ULNASTransport":
		switch {
		case under(path, "PDUSessionID"):
			return true, "the PDU session of ue.session"
		case under(path, "SNSSAI"), under(path, "DNN"):
			return false, "replaced when ue.session gives one"
		}
	case "PDUSessionEstablishmentRequest", "PDUSessionReleaseRequest", "PDUSessionReleaseComplete":
		switch {
		case path == "PDUSessionID":
			return true, "the PDU session of ue.session"
		case path == "ProcedureTransactionID":
			return true, "allocated when sent (the PTI of the command for a Release Complete)"
		case under(path, "PDUSessionType"):
			return false, "replaced when ue.session gives a type"
		}
	}
	return false, ""
}
