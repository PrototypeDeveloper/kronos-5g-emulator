// Constants of package traffic.

package traffic

// IPv4 and ICMP header values.
const (
	defaultTTL      = 64 // TTL of the packets sent
	icmpEchoRequest = 8  // ICMP type of an echo request
	icmpEchoReply   = 0  // ICMP type of an echo reply
	// maxDFEcho is the largest echo packet sent with "don't fragment" (the usual MTU).
	maxDFEcho = 1500
)

// The ephemeral ports of TCP and UDP.
const (
	ephemeralLo    = 49152
	ephemeralCount = 65536 - ephemeralLo
	// maxPort is the last port (or ICMP echo identifier) Bind gives out.
	maxPort = 65535
)

// maxParsedRate is the largest rate ParseRate accepts, in bit/s.
const maxParsedRate = 1e13
