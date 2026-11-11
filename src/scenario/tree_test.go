package scenario

import (
	"strings"
	"testing"

	libnas "libnas"
	"libnas/nasMessage"
	"libnas/nasType"
	libngap "libngap"
	"libngap/ngapMessage"
	"libngap/ngapType"

	"kronos/models"
)

func findNode(n *models.TreeNode, match func(*models.TreeNode) bool) *models.TreeNode {
	if match(n) {
		return n
	}
	for _, c := range n.Kids {
		if f := findNode(c, match); f != nil {
			return f
		}
	}
	return nil
}

func named(name string) func(*models.TreeNode) bool {
	return func(n *models.TreeNode) bool { return n.Name == name }
}

// Every template of the sample decodes into a tree that names its NAS messages.
func TestTreeOfTemplates(t *testing.T) {
	sigs, ds := ReadSignals("../../config/signal.txt")
	if ds.HasErrors() {
		t.Fatal(ds)
	}
	for _, s := range sigs.List {
		tree, err := Tree(s.Bytes, nil, true)
		if err != nil {
			t.Errorf("%s: %v", s.Name, err)
			continue
		}
		if tree.Name != s.NGAP.Name {
			t.Errorf("%s: root %q", s.Name, tree.Name)
		}
		for _, p := range s.NAS {
			if p.Name != "" && findNode(tree, named(p.Name)) == nil {
				t.Errorf("%s: no %s in the tree", s.Name, p.Name)
			}
		}
	}
	s, _ := sigs.Lookup("PDUSessionResourceSetupResponse")
	tree, _ := Tree(s.Bytes, nil, true)
	if findNode(tree, named("DLQosFlowPerTNLInformation")) == nil ||
		findNode(tree, func(n *models.TreeNode) bool {
			return n.Name == "TransportLayerAddress" && !strings.Contains(n.Value, "bits")
		}) == nil {
		t.Errorf("the response transfer is not decoded: %+v", tree)
	}
	s, _ = sigs.Lookup("InitialUEMeesageRegistrationRequest")
	tree, _ = Tree(s.Bytes, nil, true)
	if n := findNode(tree, named("RANUENGAPID")); n == nil || n.Value == "" {
		t.Errorf("RANUENGAPID %+v", n)
	}
	if n := findNode(tree, func(n *models.TreeNode) bool { return n.Name == "PLMNIdentity" }); n == nil || !strings.Contains(n.Value, "/") {
		t.Errorf("PLMNIdentity %+v", n)
	}
}

// A ciphered NAS-PDU is shown through its plaintext, else as ciphered.
func TestTreeOfCipheredNAS(t *testing.T) {
	sec := &libnas.SecurityContext{IntegrityAlgorithm: 2, CipheringAlgorithm: 2, KnasInt: [16]byte{1}, KnasEnc: [16]byte{2}}
	acc := &nasMessage.RegistrationAccept{RegistrationResult: []byte{0x01}, HasGUTI5GS: true, GUTI5GS: nasType.GUTI5GS{Value: []byte{0xf2, 2, 0xf8, 0x39, 1, 0, 0x40, 1, 2, 3, 4}}}
	plain, _ := libnas.Encode(acc, nil)
	nas, err := libnas.Encode(acc, &libnas.EncodeContext{Security: sec, SecurityHeaderType: libnas.SecurityHeaderIntegrityProtectedAndCiphered, Direction: libnas.Downlink})
	if err != nil {
		t.Fatal(err)
	}
	pdu, err := libngap.Encode(&ngapMessage.DownlinkNASTransport{AMFUENGAPID: ngapType.AMFUENGAPID{Value: 1}, RANUENGAPID: ngapType.RANUENGAPID{Value: 2}, NASPDU: ngapType.NASPDU{Value: nas}})
	if err != nil {
		t.Fatal(err)
	}
	tree, err := Tree(pdu, [][]byte{plain}, false)
	if err != nil {
		t.Fatal(err)
	}
	if findNode(tree, named("plaintext (deciphered body)")) == nil || findNode(tree, named("RegistrationAccept")) == nil || findNode(tree, named("GUTI5GS")) == nil {
		t.Errorf("with the plaintext: %+v", tree)
	}
	tree, _ = Tree(pdu, nil, false)
	if findNode(tree, named("ciphered body")) == nil || findNode(tree, named("RegistrationAccept")) != nil {
		t.Errorf("without the plaintext: %+v", tree)
	}
}

func TestBitStringText(t *testing.T) {
	for _, tc := range []struct {
		bs   ngapType.BitString
		want string
	}{
		{ngapType.BitString{Bytes: []byte{0x00, 0x40}, BitLength: 10}, "0x1 (10 bits)"},
		{ngapType.BitString{Bytes: []byte{0xfe, 0x00}, BitLength: 10}, "0x3f8 (10 bits)"},
		{ngapType.BitString{Bytes: []byte{0xc0, 0x00}, BitLength: 16}, "0xc000 (16 bits)"},
		{ngapType.BitString{Bytes: []byte{0, 0, 0, 1, 0}, BitLength: 36}, "0x10 (36 bits)"},
		{ngapType.BitString{Bytes: make([]byte, 32), BitLength: 256}, strings.Repeat("00", 32) + " (256 bits)"},
	} {
		if got := bitStringText(tc.bs); got != tc.want {
			t.Errorf("%x/%d: %s, want %s", tc.bs.Bytes, tc.bs.BitLength, got, tc.want)
		}
	}
}
