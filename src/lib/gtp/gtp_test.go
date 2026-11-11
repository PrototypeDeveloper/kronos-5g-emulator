package libgtp

import (
	"bytes"
	"encoding/binary"
	"net"
	"testing"

	"libgtp/gtpMessage"
	"libgtp/gtpType"
)

func TestGPDURoundTripWithoutOptionalHeader(t *testing.T) {
	tpdu := []byte{0x45, 0x00, 0x00, 0x14, 0xde, 0xad, 0xbe, 0xef}
	in := gtpMessage.NewGPDU(0x11223344, tpdu)

	encoded, err := Encode(in)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	if got, want := encoded[0], byte(0x30); got != want {
		t.Fatalf("flags = 0x%02x, want 0x%02x", got, want)
	}
	if got, want := encoded[1], byte(gtpType.MsgGPDU); got != want {
		t.Fatalf("message type = %d, want %d", got, want)
	}

	decodedAny, err := Decode(encoded)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	decoded := decodedAny.(*gtpMessage.GPDU)

	if decoded.Header.TEID != 0x11223344 {
		t.Fatalf("TEID = 0x%08x", decoded.Header.TEID)
	}
	if !bytes.Equal(decoded.TPDU, tpdu) {
		t.Fatalf("TPDU mismatch: %x != %x", decoded.TPDU, tpdu)
	}
	if decoded.Header.E || decoded.Header.S || decoded.Header.PN {
		t.Fatalf("unexpected optional flags: E=%v S=%v PN=%v", decoded.Header.E, decoded.Header.S, decoded.Header.PN)
	}

	reencoded, err := Encode(decoded)
	if err != nil {
		t.Fatalf("re-Encode failed: %v", err)
	}
	if !bytes.Equal(encoded, reencoded) {
		t.Fatalf("round trip mismatch:\n%x\n%x", encoded, reencoded)
	}
}

func TestGPDUWithPDUSessionContainer(t *testing.T) {
	tpdu := []byte{0x60, 0x00, 0x00, 0x00, 0xca, 0xfe}
	in, err := gtpMessage.NewGPDUWithPDUSessionContainer(
		0xaabbccdd,
		tpdu,
		gtpType.PDUSessionInformationDL,
		9,
	)
	if err != nil {
		t.Fatalf("constructor failed: %v", err)
	}

	encoded, err := Encode(in)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	// E=1 means the common 4-octet optional block is also present.
	if got, want := encoded[0], byte(0x34); got != want {
		t.Fatalf("flags = 0x%02x, want 0x%02x", got, want)
	}
	if got, want := encoded[11], byte(gtpType.ExtensionHeaderPDUSessionContainer); got != want {
		t.Fatalf("next extension type = 0x%02x, want 0x%02x", got, want)
	}
	if got, want := encoded[12], byte(1); got != want {
		t.Fatalf("PSC ext length = %d, want %d", got, want)
	}
	if got := encoded[14] & 0x3f; got != 9 {
		t.Fatalf("QFI = %d, want 9", got)
	}

	decodedAny, err := Decode(encoded)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	decoded := decodedAny.(*gtpMessage.GPDU)

	if !decoded.PDUSessionContainer.IsPresent() {
		t.Fatal("PDU Session Container not decoded")
	}
	if decoded.PDUSessionContainer.PDUType != gtpType.PDUSessionInformationDL {
		t.Fatalf("PDU type = %d", decoded.PDUSessionContainer.PDUType)
	}
	if decoded.PDUSessionContainer.QFI != 9 {
		t.Fatalf("QFI = %d", decoded.PDUSessionContainer.QFI)
	}
	if decoded.PDUSessionContainer.Header.Length != 1 {
		t.Fatalf("PSC length = %d", decoded.PDUSessionContainer.Header.Length)
	}
	if !bytes.Equal(decoded.TPDU, tpdu) {
		t.Fatalf("TPDU mismatch: %x != %x", decoded.TPDU, tpdu)
	}

	// Direct typed modification must be reflected into the extension-header chain.
	decoded.PDUSessionContainer.QFI = 17
	modified, err := Encode(decoded)
	if err != nil {
		t.Fatalf("modified Encode failed: %v", err)
	}
	modifiedAny, err := Decode(modified)
	if err != nil {
		t.Fatalf("modified Decode failed: %v", err)
	}
	if got := modifiedAny.(*gtpMessage.GPDU).PDUSessionContainer.QFI; got != 17 {
		t.Fatalf("modified QFI = %d, want 17", got)
	}
}

func TestGPDUOptionalSequenceNumber(t *testing.T) {
	in := gtpMessage.NewGPDU(0x01020304, []byte{1, 2, 3, 4})
	in.Header.S = true
	in.Header.SequenceNumber = 0x1234

	encoded, err := Encode(in)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	if got, want := encoded[0], byte(0x32); got != want {
		t.Fatalf("flags = 0x%02x, want 0x%02x", got, want)
	}

	decodedAny, err := Decode(encoded)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	decoded := decodedAny.(*gtpMessage.GPDU)
	if !decoded.Header.S || decoded.Header.SequenceNumber != 0x1234 {
		t.Fatalf("sequence not preserved: S=%v seq=0x%04x", decoded.Header.S, decoded.Header.SequenceNumber)
	}
	if decoded.Header.NPDUNumber != 0 || decoded.Header.NextExtensionHeaderType != 0 {
		t.Fatalf("unused optional fields must be zero: npdu=%d next=%d", decoded.Header.NPDUNumber, decoded.Header.NextExtensionHeaderType)
	}
}

func TestGPDUUnknownExtensionHeaderPreserved(t *testing.T) {
	in := gtpMessage.NewGPDU(0x10203040, []byte{0xaa, 0xbb})
	in.ExtensionHeaders = []gtpType.ExtensionHeader{
		gtpType.NewExtensionHeader(0x40, []byte{0x12, 0x34}),
	}

	encoded, err := Encode(in)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	decodedAny, err := Decode(encoded)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	decoded := decodedAny.(*gtpMessage.GPDU)
	if len(decoded.ExtensionHeaders) != 1 {
		t.Fatalf("extension count = %d", len(decoded.ExtensionHeaders))
	}
	if decoded.ExtensionHeaders[0].Header.Type != 0x40 {
		t.Fatalf("extension type = 0x%02x", decoded.ExtensionHeaders[0].Header.Type)
	}

	reencoded, err := Encode(decoded)
	if err != nil {
		t.Fatalf("re-Encode failed: %v", err)
	}
	if !bytes.Equal(encoded, reencoded) {
		t.Fatalf("unknown extension header not preserved:\n%x\n%x", encoded, reencoded)
	}
}

func TestGTPDecodeRejectsLengthMismatch(t *testing.T) {
	data := []byte{
		0x30, gtpType.MsgGPDU,
		0x00, 0x02,
		0x11, 0x22, 0x33, 0x44,
		0xaa,
	}
	if _, err := Decode(data); err == nil {
		t.Fatal("expected length mismatch error")
	}
}

func TestEchoRequestRoundTrip(t *testing.T) {
	in := gtpMessage.NewEchoRequest(0x1234)
	in.RecoveryTimeStamp = gtpType.NewRecoveryTimeStampRaw(0x01020304)
	priv, err := gtpType.NewPrivateExtension(0x3344, []byte{0xaa, 0xbb})
	if err != nil {
		t.Fatalf("NewPrivateExtension failed: %v", err)
	}
	in.PrivateExtension = []gtpType.PrivateExtension{priv}
	unknown, err := gtpType.NewInformationElement(200, []byte{0xde, 0xad})
	if err != nil {
		t.Fatalf("NewInformationElement failed: %v", err)
	}
	in.AdditionalIEs = []gtpType.InformationElement{unknown}

	encoded, err := Encode(in)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	if got, want := encoded[0], byte(0x32); got != want {
		t.Fatalf("flags = 0x%02x, want 0x%02x", got, want)
	}
	if got, want := encoded[1], byte(gtpType.MsgEchoRequest); got != want {
		t.Fatalf("message type = %d, want %d", got, want)
	}
	if in.Header.TEID != 0 || in.Header.SequenceNumber != 0x1234 || !in.Header.S {
		t.Fatalf("unexpected header: %+v", in.Header)
	}
	// IEs shall be sorted by ascending Type: 200, 231, 255.
	if got := encoded[12]; got != 200 {
		t.Fatalf("first IE type = %d, want 200", got)
	}

	decodedAny, err := Decode(encoded)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	decoded := decodedAny.(*gtpMessage.EchoRequest)
	if decoded.Header.TEID != 0 || decoded.Header.SequenceNumber != 0x1234 || !decoded.Header.S {
		t.Fatalf("decoded header mismatch: %+v", decoded.Header)
	}
	if decoded.RecoveryTimeStamp.Timestamp != 0x01020304 {
		t.Fatalf("timestamp = 0x%08x", decoded.RecoveryTimeStamp.Timestamp)
	}
	if decoded.RecoveryTimeStamp.Header.Type != gtpType.IERecoveryTimeStamp || decoded.RecoveryTimeStamp.Header.Length != 4 {
		t.Fatalf("timestamp header = %+v", decoded.RecoveryTimeStamp.Header)
	}
	if len(decoded.PrivateExtension) != 1 || decoded.PrivateExtension[0].ExtensionIdentifier != 0x3344 {
		t.Fatalf("private extension mismatch: %+v", decoded.PrivateExtension)
	}
	if len(decoded.AdditionalIEs) != 1 || decoded.AdditionalIEs[0].Header.Type != 200 {
		t.Fatalf("additional IE mismatch: %+v", decoded.AdditionalIEs)
	}

	reencoded, err := Encode(decoded)
	if err != nil {
		t.Fatalf("re-Encode failed: %v", err)
	}
	if !bytes.Equal(encoded, reencoded) {
		t.Fatalf("round trip mismatch:\n%x\n%x", encoded, reencoded)
	}
}

func TestEchoRequestMinimalWire(t *testing.T) {
	in := gtpMessage.NewEchoRequest(0xbeef)
	encoded, err := Encode(in)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	want := []byte{
		0x32, gtpType.MsgEchoRequest,
		0x00, 0x04,
		0x00, 0x00, 0x00, 0x00,
		0xbe, 0xef, 0x00, 0x00,
	}
	if !bytes.Equal(encoded, want) {
		t.Fatalf("wire mismatch:\n got %x\nwant %x", encoded, want)
	}
}

func TestEchoResponseRoundTrip(t *testing.T) {
	in := gtpMessage.NewEchoResponse(0x4321, gtpType.NewRecovery())
	in.RecoveryTimeStamp = gtpType.NewRecoveryTimeStampRaw(0x11223344)
	priv, err := gtpType.NewPrivateExtension(0x5566, []byte{1, 2, 3})
	if err != nil {
		t.Fatalf("NewPrivateExtension failed: %v", err)
	}
	in.PrivateExtension = []gtpType.PrivateExtension{priv}

	encoded, err := Encode(in)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	if got, want := encoded[12], byte(gtpType.IERecovery); got != want {
		t.Fatalf("first IE = %d, want Recovery(%d)", got, want)
	}

	decodedAny, err := Decode(encoded)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	decoded := decodedAny.(*gtpMessage.EchoResponse)
	if decoded.Header.SequenceNumber != 0x4321 || decoded.Header.TEID != 0 || !decoded.Header.S {
		t.Fatalf("decoded header mismatch: %+v", decoded.Header)
	}
	if decoded.Recovery.Header.Type != gtpType.IERecovery || decoded.Recovery.RestartCounter != 0 {
		t.Fatalf("Recovery mismatch: %+v", decoded.Recovery)
	}
	if decoded.RecoveryTimeStamp.Timestamp != 0x11223344 {
		t.Fatalf("timestamp = 0x%08x", decoded.RecoveryTimeStamp.Timestamp)
	}

	reencoded, err := Encode(decoded)
	if err != nil {
		t.Fatalf("re-Encode failed: %v", err)
	}
	if !bytes.Equal(encoded, reencoded) {
		t.Fatalf("round trip mismatch:\n%x\n%x", encoded, reencoded)
	}
}

func TestEchoResponseMinimalWire(t *testing.T) {
	in := gtpMessage.NewEchoResponse(0x0102, gtpType.NewRecovery())
	encoded, err := Encode(in)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	want := []byte{
		0x32, gtpType.MsgEchoResponse,
		0x00, 0x06,
		0x00, 0x00, 0x00, 0x00,
		0x01, 0x02, 0x00, 0x00,
		gtpType.IERecovery, 0x00,
	}
	if !bytes.Equal(encoded, want) {
		t.Fatalf("wire mismatch:\n got %x\nwant %x", encoded, want)
	}
}

func TestEchoResponseRequiresRecovery(t *testing.T) {
	in := gtpMessage.NewEchoResponse(1, gtpType.Recovery{})
	if _, err := Encode(in); err == nil {
		t.Fatal("expected missing Recovery error")
	}
}

func TestEchoDecodeRejectsInvalidHeader(t *testing.T) {
	// S=0 is invalid for Echo Request even when the packet length itself is valid.
	data := []byte{
		0x30, gtpType.MsgEchoRequest,
		0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}
	if _, err := Decode(data); err == nil {
		t.Fatal("expected S flag validation error")
	}
}

func TestErrorIndicationRoundTripIPv4(t *testing.T) {
	peer, err := gtpType.NewGTPUPeerAddress(net.ParseIP("192.0.2.10"))
	if err != nil {
		t.Fatalf("NewGTPUPeerAddress failed: %v", err)
	}
	in := gtpMessage.NewErrorIndication(0x2222, 0x11223344, peer)
	in.RecoveryTimeStamp = gtpType.NewRecoveryTimeStampRaw(0x01020304)
	priv, err := gtpType.NewPrivateExtension(0x3344, []byte{0xaa})
	if err != nil {
		t.Fatalf("NewPrivateExtension failed: %v", err)
	}
	in.PrivateExtension = []gtpType.PrivateExtension{priv}
	unknown, err := gtpType.NewInformationElement(200, []byte{0xde, 0xad})
	if err != nil {
		t.Fatalf("NewInformationElement failed: %v", err)
	}
	in.AdditionalIEs = []gtpType.InformationElement{unknown}

	encoded, err := Encode(in)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	if got, want := encoded[0], byte(0x32); got != want {
		t.Fatalf("flags = 0x%02x, want 0x%02x", got, want)
	}
	if got, want := encoded[1], byte(gtpType.MsgErrorIndication); got != want {
		t.Fatalf("message type = %d, want %d", got, want)
	}
	if got := binary.BigEndian.Uint32(encoded[4:8]); got != 0 {
		t.Fatalf("header TEID = 0x%08x, want zero", got)
	}
	if got, want := encoded[12], byte(gtpType.IETunnelEndpointIdentifierDataI); got != want {
		t.Fatalf("first IE = %d, want %d", got, want)
	}

	decodedAny, err := Decode(encoded)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	decoded := decodedAny.(*gtpMessage.ErrorIndication)
	if decoded.Header.TEID != 0 || !decoded.Header.S || decoded.Header.SequenceNumber != 0x2222 {
		t.Fatalf("decoded header mismatch: %+v", decoded.Header)
	}
	if got := decoded.TunnelEndpointIdentifierDataI.TEID; got != 0x11223344 {
		t.Fatalf("offending TEID = 0x%08x", got)
	}
	if !decoded.GTPUPeerAddress.IsIPv4() || !decoded.GTPUPeerAddress.IP().Equal(net.ParseIP("192.0.2.10")) {
		t.Fatalf("peer address mismatch: %+v", decoded.GTPUPeerAddress)
	}
	if decoded.GTPUPeerAddress.Header.Type != gtpType.IEGTPUPeerAddress || decoded.GTPUPeerAddress.Header.Length != 4 {
		t.Fatalf("peer address header = %+v", decoded.GTPUPeerAddress.Header)
	}
	if decoded.RecoveryTimeStamp.Timestamp != 0x01020304 {
		t.Fatalf("timestamp = 0x%08x", decoded.RecoveryTimeStamp.Timestamp)
	}
	if len(decoded.PrivateExtension) != 1 || decoded.PrivateExtension[0].ExtensionIdentifier != 0x3344 {
		t.Fatalf("private extension mismatch: %+v", decoded.PrivateExtension)
	}
	if len(decoded.AdditionalIEs) != 1 || decoded.AdditionalIEs[0].Header.Type != 200 {
		t.Fatalf("additional IE mismatch: %+v", decoded.AdditionalIEs)
	}

	reencoded, err := Encode(decoded)
	if err != nil {
		t.Fatalf("re-Encode failed: %v", err)
	}
	if !bytes.Equal(encoded, reencoded) {
		t.Fatalf("round trip mismatch:\n%x\n%x", encoded, reencoded)
	}
}

func TestErrorIndicationRoundTripIPv6(t *testing.T) {
	peer, err := gtpType.NewGTPUPeerAddress(net.ParseIP("2001:db8::10"))
	if err != nil {
		t.Fatalf("NewGTPUPeerAddress failed: %v", err)
	}
	in := gtpMessage.NewErrorIndication(0x1234, 0xaabbccdd, peer)
	encoded, err := Encode(in)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	decodedAny, err := Decode(encoded)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	decoded := decodedAny.(*gtpMessage.ErrorIndication)
	if !decoded.GTPUPeerAddress.IsIPv6() || !decoded.GTPUPeerAddress.IP().Equal(net.ParseIP("2001:db8::10")) {
		t.Fatalf("IPv6 peer address mismatch: %+v", decoded.GTPUPeerAddress)
	}
	if decoded.GTPUPeerAddress.Header.Length != 16 {
		t.Fatalf("peer address length = %d, want 16", decoded.GTPUPeerAddress.Header.Length)
	}
}

func TestErrorIndicationRequiresMandatoryIEs(t *testing.T) {
	peer, err := gtpType.NewGTPUPeerAddress(net.ParseIP("192.0.2.1"))
	if err != nil {
		t.Fatal(err)
	}
	missingTEID := gtpMessage.NewErrorIndication(1, 1, peer)
	missingTEID.TunnelEndpointIdentifierDataI = gtpType.TunnelEndpointIdentifierDataI{}
	if _, err := Encode(missingTEID); err == nil {
		t.Fatal("expected missing Tunnel Endpoint Identifier Data I error")
	}

	missingPeer := gtpMessage.NewErrorIndication(1, 1, peer)
	missingPeer.GTPUPeerAddress = gtpType.GTPUPeerAddress{}
	if _, err := Encode(missingPeer); err == nil {
		t.Fatal("expected missing GTP-U Peer Address error")
	}
}

func TestErrorIndicationDecodeRejectsInvalidHeader(t *testing.T) {
	peer, err := gtpType.NewGTPUPeerAddress(net.ParseIP("192.0.2.1"))
	if err != nil {
		t.Fatal(err)
	}
	in := gtpMessage.NewErrorIndication(1, 0x01020304, peer)
	encoded, err := Encode(in)
	if err != nil {
		t.Fatal(err)
	}
	encoded[0] = 0x30 // clear S flag, retaining a self-consistent packet length.
	if _, err := Decode(encoded); err == nil {
		t.Fatal("expected S flag validation error")
	}
}

func TestEndMarkerMinimalWire(t *testing.T) {
	in := gtpMessage.NewEndMarker(0x11223344)
	encoded, err := Encode(in)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	want := []byte{
		0x30, gtpType.MsgEndMarker,
		0x00, 0x00,
		0x11, 0x22, 0x33, 0x44,
	}
	if !bytes.Equal(encoded, want) {
		t.Fatalf("wire mismatch:\n got %x\nwant %x", encoded, want)
	}

	decodedAny, err := Decode(encoded)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	decoded := decodedAny.(*gtpMessage.EndMarker)
	if decoded.Header.TEID != 0x11223344 || decoded.Header.MessageType != gtpType.MsgEndMarker {
		t.Fatalf("decoded header mismatch: %+v", decoded.Header)
	}
}

func TestEndMarkerWithPDUSessionContainerRoundTrip(t *testing.T) {
	in, err := gtpMessage.NewEndMarkerWithPDUSessionContainer(
		0xaabbccdd,
		gtpType.PDUSessionInformationDL,
		9,
	)
	if err != nil {
		t.Fatalf("NewEndMarkerWithPDUSessionContainer failed: %v", err)
	}

	encoded, err := Encode(in)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	if encoded[0]&0x04 == 0 {
		t.Fatal("E flag is not set")
	}
	if got := encoded[11]; got != gtpType.ExtensionHeaderPDUSessionContainer {
		t.Fatalf("next extension header = 0x%02x", got)
	}

	decodedAny, err := Decode(encoded)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	decoded := decodedAny.(*gtpMessage.EndMarker)
	if decoded.Header.TEID != 0xaabbccdd {
		t.Fatalf("TEID = 0x%08x", decoded.Header.TEID)
	}
	if !decoded.PDUSessionContainer.IsPresent() {
		t.Fatal("PDU Session Container not decoded")
	}
	if decoded.PDUSessionContainer.QFI != 9 {
		t.Fatalf("QFI = %d", decoded.PDUSessionContainer.QFI)
	}
	if decoded.PDUSessionContainer.PDUType != gtpType.PDUSessionInformationDL {
		t.Fatalf("PDU Type = %d", decoded.PDUSessionContainer.PDUType)
	}

	reencoded, err := Encode(decoded)
	if err != nil {
		t.Fatalf("re-Encode failed: %v", err)
	}
	if !bytes.Equal(encoded, reencoded) {
		t.Fatalf("round trip mismatch:\n%x\n%x", encoded, reencoded)
	}
}

func TestEndMarkerPrivateExtensionRoundTrip(t *testing.T) {
	in := gtpMessage.NewEndMarker(0x01020304)
	priv, err := gtpType.NewPrivateExtension(0x3344, []byte{0xaa, 0xbb})
	if err != nil {
		t.Fatal(err)
	}
	in.PrivateExtension = []gtpType.PrivateExtension{priv}

	encoded, err := Encode(in)
	if err != nil {
		t.Fatal(err)
	}
	decodedAny, err := Decode(encoded)
	if err != nil {
		t.Fatal(err)
	}
	decoded := decodedAny.(*gtpMessage.EndMarker)
	if len(decoded.PrivateExtension) != 1 || decoded.PrivateExtension[0].ExtensionIdentifier != 0x3344 {
		t.Fatalf("Private Extension mismatch: %+v", decoded.PrivateExtension)
	}
}

func TestEndMarkerRejectsZeroTEID(t *testing.T) {
	in := gtpMessage.NewEndMarker(0)
	if _, err := Encode(in); err == nil {
		t.Fatal("expected zero TEID error")
	}
}
