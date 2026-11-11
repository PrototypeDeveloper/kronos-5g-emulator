// Constants of the engine.

package engine

import "time"

// Event bus.
const (
	// eventHistory is how many events the bus keeps for late subscribers.
	eventHistory = 10000
)

// Actors.
const (
	// nodeInbox is the queue size of the node actor's inbox.
	nodeInbox = 1024
	// subscriberInbox is the queue size of a subscriber actor's inbox.
	subscriberInbox = 256
	// maxEntryChain bounds entry blocks that move again.
	maxEntryChain = 16
)

// Commands and batches.
const (
	// maxBatchDepth bounds batches that include batches.
	maxBatchDepth = 8
	// subscriberReplyTimeout is how long a command for one subscriber waits for its answer.
	subscriberReplyTimeout = 30 * time.Second
)

// N2 (SCTP).
const (
	// sctpDialTimeout bounds the setup of the SCTP association (linkup).
	sctpDialTimeout = 10 * time.Second
)

// N3 (GTP-U) and the IP packets it carries.
const (
	// socketBuffer is the kernel buffer asked for each way: iperf at high rates
	// needs more than the default.
	socketBuffer = 8 << 20
	// packetBuffer is the read buffer of the GTP-U socket and the TUN devices:
	// room for the largest UDP datagram or IP packet.
	packetBuffer = 65536
	// ipv4HeaderLen is the length of an IPv4 header without options (RFC 791).
	ipv4HeaderLen = 20
	// ipv4Version is the version field of an IPv4 header.
	ipv4Version = 4
	// icmpEchoRequest and icmpEchoReply are the ICMP types of echo (RFC 792).
	icmpEchoRequest = 8
	icmpEchoReply   = 0
)

// TUN devices (ue.tun).
const (
	// tunMTU leaves room for the G-PDU around the packet in a 1500-octet
	// frame: IPv4 20, UDP 8, GTP-U 8 and its PDU Session Container 8.
	tunMTU = 1500 - 44
	// tunTable and tunPriority place the routing of subscriber i (from 0)
	// in table tunTable+1+i, looked up by a source rule of tunPriority.
	tunTable    = 10000
	tunPriority = 1000
)

// timeLayout is how the builtin $time shows the current time.
const timeLayout = "15:04:05.000"
