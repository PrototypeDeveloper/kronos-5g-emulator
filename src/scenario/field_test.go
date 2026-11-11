package scenario

import (
	"encoding/hex"
	"strings"
	"testing"

	libnas "libnas"
	"libnas/nasMessage"
	"libnas/nasType"
	"libngap/ngapMessage"
	"libngap/ngapType"
	"libngap/ngapcodec"

	"kronos/command"
)

// fieldPath parses the path of $proto.path.
func fieldPath(t *testing.T, src string) []command.FieldSeg {
	t.Helper()
	e, err := command.ParseExpr(command.NewLexer("", "$x."+src))
	if err != nil || e.Kind != command.ExprField {
		t.Fatalf("%s: %v", src, err)
	}
	return e.Path
}

func TestFieldValueNGAP(t *testing.T) {
	guami := ngapType.GUAMI{
		PLMNIdentity: ngapType.PLMNIdentity{Value: []byte{0x02, 0xf8, 0x39}},
		AMFRegionID:  ngapType.AMFRegionID{Value: ngapType.BitString{Bytes: []byte{0xca}, BitLength: 8}},
		AMFSetID:     ngapType.AMFSetID{Value: ngapType.BitString{Bytes: []byte{0x00, 0x40}, BitLength: 10}},
		AMFPointer:   ngapType.AMFPointer{Value: ngapType.BitString{Bytes: []byte{0x04}, BitLength: 6}},
	}
	resp := &ngapMessage.NGSetupResponse{
		AMFName:             ngapType.AMFName{Value: "amf1"},
		ServedGUAMIList:     ngapType.ServedGUAMIList{List: []ngapType.ServedGUAMIItem{{GUAMI: guami}, {GUAMI: guami, HasBackupAMFName: true, BackupAMFName: ngapType.AMFName{Value: "backup"}}}},
		RelativeAMFCapacity: ngapType.RelativeAMFCapacity{Value: 255},
		PLMNSupportList:     ngapType.PLMNSupportList{List: []ngapType.PLMNSupportItem{{PLMNIdentity: ngapType.PLMNIdentity{Value: []byte{0x02, 0xf8, 0x39}}}}},
	}
	fail := &ngapMessage.NGSetupFailure{Cause: ngapType.Cause{Present: ngapType.CausePresentMisc, Misc: ngapType.CauseMisc{Value: ngapType.CauseMiscPresentControlProcessingOverload}}}

	for msg, cases := range map[any]map[string]command.Value{
		resp: {
			"AMFName":                          command.StrVal("amf1"),
			"amfname":                          command.StrVal("amf1"), // names are matched ignoring case
			"RelativeAMFCapacity":              command.IntVal(255),
			"AMFRegionID":                      command.IntVal(0xca), // a short BIT STRING is a number
			"AMFSetID":                         command.IntVal(1),
			"ServedGUAMIList[0].AMFPointer":    command.IntVal(1),
			"PLMNIdentity":                     command.StrVal("02f839"), // octets in hex
			"PLMNSupportList[0].PLMNIdentity":  command.StrVal("02f839"),
			"BackupAMFName":                    command.StrVal("backup"), // the first item that has it
			"ServedGUAMIList[0].BackupAMFName": {},                       // an optional IE that is absent
			"ServedGUAMIList[1].BackupAMFName": command.StrVal("backup"),
			"ServedGUAMIList[2].AMFSetID":      {}, // no such item
			"CriticalityDiagnostics":           {},
			"Cause":                            {},
			"NoSuchIE":                         {},
		},
		fail: {
			"Cause":      command.StrVal("Misc"), // a CHOICE: the alternative it took
			"Cause.Misc": command.StrVal("control-processing-overload"),
			"Misc":       command.StrVal("control-processing-overload"),
			"Cause.Nas":  {}, // an alternative not taken
			"TimeToWait": {},
		},
	} {
		for path, want := range cases {
			got, err := FieldValue([]any{nil, msg}, fieldPath(t, path))
			if err != nil || got != want {
				t.Errorf("%T %s = %+v, %v; want %+v", msg, path, got, err, want)
			}
		}
	}
	for path, want := range map[string]string{
		"ServedGUAMIList": "ServedGUAMIList is a list: pick an item with [index]",
		"GUAMI":           "GUAMI is a group of IEs: name one of them",
	} {
		if _, err := FieldValue([]any{resp}, fieldPath(t, path)); err == nil || err.Error() != want {
			t.Errorf("%s: want %q, got %v", path, want, err)
		}
	}
}

func TestFieldValueNAS(t *testing.T) {
	enc := func(m any) []byte {
		t.Helper()
		b, err := libnas.Encode(m, nil)
		if err != nil {
			t.Fatal(err)
		}
		return b
	}
	reject := enc(&nasMessage.RegistrationReject{Cause: 3, HasT3502Value: true, T3502Value: nasType.T3502Value{Value: 0x21}})
	sm := enc(&nasMessage.PDUSessionEstablishmentReject{PDUSessionID: 5, ProcedureTransactionID: 1, Cause: 27})
	dl := enc(&nasMessage.DLNASTransport{PayloadContainerType: payloadN1SMInfo, PayloadContainer: sm, HasPDUSessionID: true, PDUSessionID: nasType.PDUSessionID{Value: 5}})

	roots := NASMessages([][]byte{reject})
	for path, want := range map[string]command.Value{
		"Cause":      command.IntVal(3),
		"T3502Value": command.IntVal(0x21),
		"T3346Value": {},
	} {
		if got, err := FieldValue(roots, fieldPath(t, path)); err != nil || got != want {
			t.Errorf("RegistrationReject %s = %+v, %v; want %+v", path, got, err, want)
		}
	}
	// The 5GSM message a DL NAS Transport carries is read first.
	roots = NASMessages([][]byte{dl, {0xff}})
	if len(roots) != 2 {
		t.Fatalf("decoded %d messages", len(roots))
	}
	for path, want := range map[string]command.Value{
		"Cause":                  command.IntVal(27),
		"ProcedureTransactionID": command.IntVal(1),
		"PayloadContainerType":   command.IntVal(payloadN1SMInfo),
		"PDUSessionID":           command.IntVal(5),
		"Cause5GMM":              {},
	} {
		if got, err := FieldValue(roots, fieldPath(t, path)); err != nil || got != want {
			t.Errorf("DLNASTransport %s = %+v, %v; want %+v", path, got, err, want)
		}
	}
	if got, _ := FieldValue(roots, fieldPath(t, "PayloadContainer")); !strings.HasPrefix(got.Str, "2e05") {
		t.Errorf("PayloadContainer = %+v", got)
	}
}

// TestFieldValueTransfer: the IEs of a transfer (an NGAP value in octets) can be named.
func TestFieldValueTransfer(t *testing.T) {
	reject := ngapType.Criticality{Value: ngapType.CriticalityPresentReject}
	tr, err := ngapcodec.EncodeParameter("PDUSessionResourceSetupRequestTransfer", ngapType.PDUSessionResourceSetupRequestTransfer{
		ProtocolIEs: []ngapType.PDUSessionResourceSetupRequestTransferIEsEntry{
			{ID: ngapType.ProtocolIEID{Value: 139}, Criticality: reject,
				Present: ngapType.PDUSessionResourceSetupRequestTransferIEsEntryPresentULNGUUPTNLInformation,
				ULNGUUPTNLInformation: ngapType.UPTransportLayerInformation{Present: ngapType.UPTransportLayerInformationPresentGTPTunnel,
					GTPTunnel: ngapType.GTPTunnel{
						TransportLayerAddress: ngapType.TransportLayerAddress{Value: ngapType.BitString{Bytes: []byte{10, 10, 30, 40}, BitLength: 32}},
						GTPTEID:               ngapType.GTPTEID{Value: []byte{0, 0, 0x10, 0x01}},
					}}},
			{ID: ngapType.ProtocolIEID{Value: 134}, Criticality: reject,
				Present:        ngapType.PDUSessionResourceSetupRequestTransferIEsEntryPresentPDUSessionType,
				PDUSessionType: ngapType.PDUSessionType{Value: ngapType.PDUSessionTypePresentIpv4}},
			{ID: ngapType.ProtocolIEID{Value: 136}, Criticality: reject,
				Present: ngapType.PDUSessionResourceSetupRequestTransferIEsEntryPresentQosFlowSetupRequestList,
				QosFlowSetupRequestList: ngapType.QosFlowSetupRequestList{List: []ngapType.QosFlowSetupRequestItem{{
					QosFlowIdentifier: ngapType.QosFlowIdentifier{Value: 5},
					QosFlowLevelQosParameters: ngapType.QosFlowLevelQosParameters{
						QosCharacteristics: ngapType.QosCharacteristics{Present: ngapType.QosCharacteristicsPresentNonDynamic5QI,
							NonDynamic5QI: ngapType.NonDynamic5QIDescriptor{FiveQI: ngapType.FiveQI{Value: 9}}},
						AllocationAndRetentionPriority: ngapType.AllocationAndRetentionPriority{
							PriorityLevelARP:        ngapType.PriorityLevelARP{Value: 8},
							PreEmptionCapability:    ngapType.PreEmptionCapability{Value: ngapType.PreEmptionCapabilityPresentShallNotTriggerPreEmption},
							PreEmptionVulnerability: ngapType.PreEmptionVulnerability{Value: ngapType.PreEmptionVulnerabilityPresentNotPreEmptable},
						},
					},
				}}}},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	req := &ngapMessage.PDUSessionResourceSetupRequest{PDUSessionResourceSetupListSUReq: ngapType.PDUSessionResourceSetupListSUReq{
		List: []ngapType.PDUSessionResourceSetupItemSUReq{{PDUSessionID: ngapType.PDUSessionID{Value: 1}, PDUSessionResourceSetupRequestTransfer: tr}},
	}}
	for path, want := range map[string]command.Value{
		"TransportLayerAddress": command.StrVal("10.10.30.40"), // an address as the signal editor shows it
		"GTPTEID":               command.StrVal("00001001"),
		"PDUSessionResourceSetupListSUReq[0].PDUSessionResourceSetupRequestTransfer.QosFlowIdentifier": command.IntVal(5),
		"QosFlowSetupRequestList[0].FiveQI":      command.IntVal(9),
		"ULNGUUPTNLInformation":                  command.StrVal("GTPTunnel"),
		"PDUSessionType":                         command.StrVal("ipv4"),
		"PreEmptionCapability":                   command.StrVal("shall-not-trigger-pre-emption"),
		"PDUSessionResourceSetupRequestTransfer": command.StrVal(hex.EncodeToString(tr)), // the transfer itself is octets
		"PDUSessionAggregateMaximumBitRate":      {},                                     // an IE the entries do not have
		"Criticality":                            {},
	} {
		if got, err := FieldValue([]any{req}, fieldPath(t, path)); err != nil || got != want {
			t.Errorf("%s = %+v, %v; want %+v", path, got, err, want)
		}
	}
}
