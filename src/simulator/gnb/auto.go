package gnb

import (
	"strings"

	"kronos/models"
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
		return true, mapNote("the NGAP ID of the UE association", models.MapNGAPID)
	case under(path, "UserLocationInformation"):
		return true, mapNote("the NR cell and TA of the config (a TimeStamp: the time of sending)", models.MapPLMN, models.MapNCI, models.MapTAC, models.MapTimeStamp)
	}
	if l, ok := autoLists[msg]; ok {
		switch {
		case path == l.list:
			return false, mapNote("one item per PDU session of the procedure; the first item is the pattern", models.MapTunnel)
		case under(path, l.list+"/List/*/PDUSessionID"):
			return true, mapNote("the PDU session", models.MapTunnel)
		}
		for _, f := range l.fields {
			if under(path, l.list+"/List/*/"+l.transfer+"/>/"+f) {
				return true, mapNote("gtpu.src, the DL TEID and the QoS flows of the session", models.MapTunnel)
			}
		}
	}
	switch msg {
	case "RANConfigurationUpdate":
		switch {
		case under(path, "GlobalRANNodeID"):
			return true, mapNote("gnodeb.plmn and gnodeb.gnodebid", models.MapPLMN, models.MapGNodeBID)
		case under(path, "SupportedTAList"):
			return true, mapNote("gnodeb.tac, gnodeb.plmn and gnodeb.snssai (the first items are the pattern)", models.MapTAC, models.MapPLMN, models.MapSNSSAI)
		case under(path, "RANNodeName") && r.name != "":
			return true, mapNote("gnodeb.name", models.MapRANName)
		}
	case "NGReset":
		if under(path, "ResetType/PartOfNGInterface") {
			return true, mapNote("a partial reset names the current UE associations", models.MapNGAPID)
		}
	case "NGResetAcknowledge":
		if under(path, "UEAssociatedLogicalNGConnectionList") {
			return true, mapNote("the UE associations the AMF's NG Reset named", models.MapNGAPID)
		}
	case "PathSwitchRequest":
		switch {
		case under(path, "SourceAMFUENGAPID"):
			return true, mapNote("the AMF-UE-NGAP-ID at the source (from the handover context)", models.MapNGAPID)
		case under(path, "UESecurityCapabilities/NRencryptionAlgorithms"), under(path, "UESecurityCapabilities/NRintegrityProtectionAlgorithms"):
			return true, mapNote("ue.ciphering and ue.integrity", models.MapSecCap)
		}
	case "CellTrafficTrace":
		if under(path, "NGRANCGI") {
			return true, mapNote("the NR cell of the config", models.MapPLMN, models.MapNCI)
		}
	}
	switch msg {
	case "InitialUEMessage":
		if under(path, "FiveGSTMSI") {
			return true, mapNote("from the 5G-GUTI (left out without one)", models.MapIdentity)
		}
	case "NGSetupRequest":
		switch {
		case under(path, "GlobalRANNodeID"):
			return true, mapNote("gnodeb.plmn and gnodeb.gnodebid", models.MapPLMN, models.MapGNodeBID)
		case under(path, "SupportedTAList"):
			return true, mapNote("gnodeb.tac, gnodeb.plmn and gnodeb.snssai (the first items are the pattern)", models.MapTAC, models.MapPLMN, models.MapSNSSAI)
		case under(path, "RANNodeName") && r.name != "":
			return true, mapNote("gnodeb.name", models.MapRANName)
		}
	case "PDUSessionResourceSetupResponse", "InitialContextSetupResponse":
		list, failed := "PDUSessionResourceSetupListSURes", "PDUSessionResourceFailedToSetupListSURes"
		if msg == "InitialContextSetupResponse" {
			list, failed = "PDUSessionResourceSetupListCxtRes", "PDUSessionResourceFailedToSetupListCxtRes"
		}
		const tr = "/List/*/PDUSessionResourceSetupResponseTransfer/>/"
		switch {
		case path == list:
			return false, mapNote("one item per PDU session answered; the first item is the pattern", models.MapTunnel)
		case under(path, list+"/List/*/PDUSessionID"):
			return true, mapNote("the PDU session answered", models.MapTunnel)
		case under(path, list+tr+"DLQosFlowPerTNLInformation"):
			return true, mapNote("gtpu.src, the DL TEID and the QoS flows of the session", models.MapTunnel)
		case under(path, list+tr+"AdditionalDLQosFlowPerTNLInformation"), under(path, list+tr+"QosFlowFailedToSetupList"), under(path, failed):
			return true, mapNote("left out when sent", models.MapTunnel)
		}
	case "PDUSessionResourceReleaseResponse":
		const list = "PDUSessionResourceReleasedListRelRes"
		switch {
		case path == list:
			return false, mapNote("one item per PDU session released; the first item is the pattern", models.MapTunnel)
		case under(path, list+"/List/*/PDUSessionID"):
			return true, mapNote("the PDU session released", models.MapTunnel)
		}
	case "UEContextReleaseRequest":
		if under(path, "PDUSessionResourceListCxtRelReq") {
			return true, mapNote("the established PDU sessions", models.MapTunnel)
		}
	case "UEContextReleaseComplete":
		if under(path, "PDUSessionResourceListCxtRelCpl") {
			return true, mapNote("the established PDU sessions", models.MapTunnel)
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
			return false, mapNote("the ngKSI (high 4 bits) is replaced", models.MapNgKSI)
		case under(path, "MobileIdentity"):
			return true, mapNote("the SUCI of the IMSI, or the current 5G-GUTI (by the identity type of the template)", models.MapIdentity)
		case under(path, "UESecurityCapability"):
			return false, mapNote("the first two octets (5G-EA, 5G-IA) come from ue.ciphering and ue.integrity", models.MapSecCap)
		case under(path, "RequestedNSSAI") && r.nssai != nil:
			return true, mapNote("ue.nssai", models.MapSNSSAI)
		case status:
			return true, mapNote("the established PDU sessions", models.MapStatus)
		case under(path, "NASMessageContainer") && !inContainer:
			return true, mapNote("rebuilt from this message when sent under a security context", models.MapContainer)
		}
	case "ServiceRequest":
		switch {
		case path == "ServiceTypeAndNgKSI":
			return false, mapNote("the ngKSI (low 4 bits) is replaced", models.MapNgKSI)
		case under(path, "FiveGSTMSI"):
			return true, mapNote("the 5G-S-TMSI of the 5G-GUTI", models.MapIdentity)
		case status:
			return true, mapNote("the established PDU sessions", models.MapStatus)
		}
	case "ControlPlaneServiceRequest":
		switch {
		case path == "ControlPlaneServiceTypeAndNgKSI":
			return false, mapNote("the ngKSI (high 4 bits) is replaced", models.MapNgKSI)
		case status:
			return true, mapNote("the established PDU sessions", models.MapStatus)
		case under(path, "NASMessageContainer") && outer == "":
			return true, mapNote("rebuilt from this message when sent under a security context", models.MapContainer)
		}
	case "NotificationResponse":
		if status {
			return true, mapNote("the established PDU sessions", models.MapStatus)
		}
	case "NetworkSliceSpecificAuthenticationComplete":
		if under(path, "SNSSAI") {
			return false, mapNote("the S-NSSAI of the last Network Slice-Specific Authentication Command", models.MapSNSSAI)
		}
	case "DeregistrationRequestUE":
		switch {
		case path == "DeregistrationTypeAndNgKSI":
			return false, mapNote("the ngKSI (high 4 bits) is replaced", models.MapNgKSI)
		case under(path, "MobileIdentity"):
			return true, mapNote("the current 5G-GUTI, or the SUCI of the IMSI (by the identity type of the template)", models.MapIdentity)
		}
	case "IdentityResponse":
		if under(path, "MobileIdentity") {
			return true, mapNote("the identity the Identity Request asked for: SUCI, 5G-GUTI, IMEI, 5G-S-TMSI or IMEISV", models.MapIdentity)
		}
	case "AuthenticationFailure":
		if under(path, "AuthenticationFailureParameter") {
			return true, mapNote("the AUTS of a synch failure from the last Authentication Request (left out for other causes)", models.MapRES)
		}
	case "Status5GSM":
		if path == "PDUSessionID" {
			return true, mapNote("the PDU session of ue.session", models.MapSession)
		}
	case "AuthenticationResponse":
		if under(path, "AuthenticationResponseParameter") {
			return true, mapNote("the RES* computed from the AuthenticationRequest", models.MapRES)
		}
	case "SecurityModeComplete":
		if under(path, "IMEISV") {
			return false, mapNote("replaced by ue.imeisv when the Security Mode Command asks for it", models.MapIMEISV)
		}
	case "ULNASTransport":
		switch {
		case under(path, "PDUSessionID"):
			return true, mapNote("the PDU session of ue.session", models.MapSession)
		case under(path, "SNSSAI"):
			return false, mapNote("replaced when ue.session gives one", models.MapSNSSAI)
		case under(path, "DNN"):
			return false, mapNote("replaced when ue.session gives one", models.MapSession)
		}
	case "PDUSessionEstablishmentRequest", "PDUSessionReleaseRequest", "PDUSessionModificationRequest",
		"PDUSessionReleaseComplete", "PDUSessionModificationComplete", "PDUSessionModificationCommandReject", "PDUSessionAuthenticationComplete":
		switch {
		case path == "PDUSessionID":
			return true, mapNote("the PDU session of ue.session", models.MapSession)
		case path == "ProcedureTransactionID" && strings.HasSuffix(nas, "Request"):
			return true, mapNote("allocated when sent", models.MapSession)
		case path == "ProcedureTransactionID":
			return true, mapNote("the PTI of the command it answers", models.MapSession)
		case under(path, "PDUSessionType"):
			return false, mapNote("replaced when ue.session gives a type", models.MapSession)
		}
	}
	return false, ""
}

// mapNote adds to a note the mapping items (mapon / mapoff) that control the field.
func mapNote(note string, items ...models.MapItem) string {
	names := make([]string, len(items))
	for i, it := range items {
		names[i] = it.String()
	}
	return note + " [mapping: " + strings.Join(names, ", ") + "]"
}

// autoLists are the PDU session lists of the procedures besides session
// setup and release, by NGAP message.
var autoLists = map[string]autoList{
	"HandoverRequired":                   {list: "PDUSessionResourceListHORqd"},
	"HandoverRequestAcknowledge":         {"PDUSessionResourceAdmittedList", "HandoverRequestAcknowledgeTransfer", []string{"DLNGUUPTNLInformation", "QosFlowSetupResponseList"}},
	"PathSwitchRequest":                  {"PDUSessionResourceToBeSwitchedDLList", "PathSwitchRequestTransfer", []string{"DLNGUUPTNLInformation", "QosFlowAcceptedList"}},
	"PDUSessionResourceModifyResponse":   {"PDUSessionResourceModifyListModRes", "PDUSessionResourceModifyResponseTransfer", []string{"DLNGUUPTNLInformation", "QosFlowAddOrModifyResponseList"}},
	"PDUSessionResourceModifyIndication": {"PDUSessionResourceModifyListModInd", "PDUSessionResourceModifyIndicationTransfer", []string{"DLQosFlowPerTNLInformation"}},
	"PDUSessionResourceNotify":           {list: "PDUSessionResourceNotifyList"},
	"InitialContextSetupFailure":         {list: "PDUSessionResourceFailedToSetupListCxtFail"},
}
