package engine_test

import (
	"strings"
	"testing"

	libnas "libnas"
	"libnas/nasMessage"
	"libnas/nasType"
	"libngap/ngapMessage"
	"libngap/ngapType"

	"kronos/models"
)

const fieldScenario = `
var { private c = 0 }
node {
	entry { send(NGSetupRequest) }
	receive(NGSetupResponse) {
		snap("amf %s capacity %d plmn %s", $ngap.AMFName, $ngap.RelativeAMFCapacity, $ngap.PLMNSupportList[0].PLMNIdentity)
	}
}
start {
	entry { send(InitialUEMeesageRegistrationRequest) }
	receive(DownLinkNasTransport:RegistrationReject) {
		if ($nas.Cause == 3 && $ngap.AMFUENGAPID == 100) {
			snap("illegal ue, t3502 %d", $nas.T3502Value)
		} else {
			snap("other cause %d", $nas.Cause)
		}
		set(c = $nas.T3346Value)
		snap("after c=%d", c)
	}
	receive(DownlinkNASTransport:ConfigurationUpdateCommand) {
		snap("new guti %s ack %d", $nas.GUTI5GS, $nas.ConfigurationUpdateIndication)
	}
}
`

// TestFieldOfTheReceivedMessage: handlers read IEs of the message that started them.
func TestFieldOfTheReceivedMessage(t *testing.T) {
	hs := start(t, true, 1, fieldScenario)
	if err := hs.eng.LinkUp(); err != nil {
		t.Fatal(err)
	}
	hs.amf.accept(t)
	recvAs[*ngapMessage.NGSetupRequest](t, hs.amf)
	hs.amf.send(t, setupResponse())
	waitAll(t, hs, "amf stub-amf capacity 255 plmn 02f839")

	ium := recvAs[*ngapMessage.InitialUEMessage](t, hs.amf)
	nas, err := libnas.Encode(&nasMessage.RegistrationReject{Cause: 3, HasT3502Value: true, T3502Value: nasType.T3502Value{Value: 0x22}}, nil)
	if err != nil {
		t.Fatal(err)
	}
	hs.amf.send(t, &ngapMessage.DownlinkNASTransport{AMFUENGAPID: ngapType.AMFUENGAPID{Value: 100}, RANUENGAPID: ium.RANUENGAPID, NASPDU: ngapType.NASPDU{Value: nas}})
	waitAll(t, hs, "illegal ue, t3502 34")
	// An IE the message does not have stops the set (an error) and the body goes on.
	ev := hs.wait(t, "the error of a missing IE", func(e models.Event) bool { return e.Kind == models.KindError })
	if !strings.Contains(ev.Text, "set(c = $nas.T3346Value)") || !strings.Contains(ev.Text, "no such IE") {
		t.Fatalf("error %+v", ev)
	}
	waitAll(t, hs, "after c=0")

	// A message libnas has only lately: named, matched and read.
	nas, err = libnas.Encode(&nasMessage.ConfigurationUpdateCommand{HasConfigurationUpdateIndication: true, ConfigurationUpdateIndication: nasType.ConfigurationUpdateIndication{Value: 1},
		HasGUTI5GS: true, GUTI5GS: nasType.GUTI5GS{Value: []byte{0xf2, 0x02, 0xf8, 0x39, 0xca, 0xfe, 0, 0, 0, 0, 0x2a}}}, nil)
	if err != nil {
		t.Fatal(err)
	}
	hs.amf.send(t, &ngapMessage.DownlinkNASTransport{AMFUENGAPID: ngapType.AMFUENGAPID{Value: 100}, RANUENGAPID: ium.RANUENGAPID, NASPDU: ngapType.NASPDU{Value: nas}})
	waitAll(t, hs, "new guti f202f839cafe000000002a ack 1")
}
