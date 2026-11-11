// Constants of package iperf.

package iperf

import "time"

// The version the client announces, and the read buffer of TCP tests.
const (
	clientVersion = "3.9"
	tcpReadBufLen = 256 << 10 // smallest read buffer of a downlink TCP test
)

// Control states (iperf_api.h), sent as one signed octet.
const (
	testStart       = 1
	testRunning     = 2
	testEnd         = 4
	paramExchange   = 9
	createStreams   = 10
	serverTerminate = 11
	clientTerminate = 12
	exchangeResults = 13
	displayResults  = 14
	iperfDone       = 16
	accessDenied    = -1
	serverError     = -2
)

// The protocol.
const (
	cookieLen = 37 // 36 characters and a NUL
	// base32 is the alphabet of the test cookie.
	base32      = "abcdefghijklmnopqrstuvwxyz234567"
	maxJSONLen  = 1 << 20 // largest JSON text accepted on the control connection
	pacingTimer = 1000    // the pacing_timer parameter (microseconds), as iperf3 by default
	// udpConnectLen is the length of the UDP connect message and its reply.
	udpConnectLen = 4
)

// udpHeader is the start of each data datagram: sender time and a packet
// counter (32 bits; iperf3 uses 64 with --udp-counters-64bit).
const udpHeader = 12

// Timeouts of the control connection.
const (
	connectTimeout = 5 * time.Second  // to open the control (and data) connection
	stepTimeout    = 5 * time.Second  // for each step of the setup, and of the server's end of a test
	resultsTimeout = 10 * time.Second // for each step of the client's end of a test
	// promptReadTimeout bounds a read of what has arrived or follows at once.
	promptReadTimeout = time.Second
	// testEndSlack is how many seconds past the test time the server waits for TEST_END.
	testEndSlack = 10
)

// Pacing of the UDP data.
const (
	tickInterval = time.Millisecond // the tick of the sending (and client test) loop
	maxBurst     = 4096             // most datagrams sent at one tick, after a stall
)
