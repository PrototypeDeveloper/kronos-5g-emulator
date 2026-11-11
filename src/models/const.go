// Constants of package models.

package models

import "time"

// Event kinds.
const (
	KindMessage = "message" // an NGAP PDU was sent (tx) or received (rx)
	KindState   = "state"   // a subscriber changed state
	KindSnap    = "snap"    // snap() output
	KindLink    = "link"    // SCTP up/down, NG Setup result
	KindInfo    = "info"    // results of built-in processing: authentication, security mode, registration
	KindCommand = "command" // summary of a command run on many subscribers
	KindWarn    = "warn"    // something unexpected that did not stop processing
	KindError   = "error"   // a failure, e.g. a send that could not be built
)

// SourceNode is Event.Source for the node; subscribers use their IMSI.
const SourceNode = "node"

// SourceWeb is Event.Source for a command entered in the web console.
const SourceWeb = "web"

// NG Setup states.
const (
	SetupNone   = ""
	SetupOK     = "ok"
	SetupFailed = "failed"
)

// Registration and connection states of a subscriber, as the signals tell
// (for display only, they do not drive the scenario).
const (
	RMRegistered   = "REGISTERED"
	RMDeregistered = "DEREGISTERED"
	CMConnected    = "CONNECTED"
	CMIdle         = "IDLE"
)

// Defaults and limits of iperf (package traffic/iperf).
const (
	IperfDefaultPort     = 5201
	IperfDefaultRate     = 1_000_000 // bit/s, as iperf3 for UDP
	IperfDefaultDuration = 10 * time.Second
	// IperfDefaultLength keeps the inner packet (and so the G-PDU) within 1500 octets.
	IperfDefaultLength = 1400
	IperfMinLength     = 16
	IperfMaxLength     = 65000
	IperfMaxRate       = 100_000_000_000
	IperfMaxDuration   = 24 * time.Hour
)

// Defaults and limits of iperf TCP tests (iperftcp).
const (
	IperfDefaultTCPLength = 128 << 10 // octets per write, as iperf3 for TCP
	IperfMaxTCPLength     = 1 << 20
)

// IP protocol numbers.
const (
	ProtoICMP uint8 = 1
	ProtoTCP  uint8 = 6
	ProtoUDP  uint8 = 17
)

// PagingName is non-UE-associated in NGAP but is routed to a subscriber by
// its 5G-S-TMSI, so its trigger is written in the subscriber blocks.
const PagingName = "Paging"

// NAS extended protocol discriminators (TS 24.007).
const (
	EPD5GMM uint8 = 0x7e
	EPD5GSM uint8 = 0x2e
)

// PPIDNGAP is the payload protocol identifier of NGAP (TS 38.412 §7).
const PPIDNGAP uint32 = 60
