// Constants of package sctp.

package sctp

import (
	"syscall"
	"time"
)

// Linux constants from <linux/sctp.h> that package syscall does not define.
const (
	solSCTP         = syscall.IPPROTO_SCTP // SOL_SCTP
	sctpInitMsg     = 2
	sctpNoDelay     = 3
	sctpStatus      = 14
	sctpRecvRcvInfo = 32
	cmsgSndInfo     = 2 // SCTP_SNDINFO
	cmsgRcvInfo     = 3 // SCTP_RCVINFO
	msgNotification = 0x8000
	msgEOR          = 0x80

	sndInfoLen = 16  // struct sctp_sndinfo
	rcvInfoLen = 28  // struct sctp_rcvinfo
	initMsgLen = 8   // struct sctp_initmsg
	statusLen  = 176 // struct sctp_status
)

// Defaults of Options.
const (
	defaultStreams = 16               // outbound streams requested and inbound streams accepted
	defaultTimeout = 10 * time.Second // Dial: time allowed to set up the association
)

// Buffers and queues of the sockets.
const (
	recvBufLen    = 64 * 1024 // read buffer of a Conn; a larger message takes several reads
	oobExtra      = 64        // room past SCTP_RCVINFO in the control message buffer
	listenBacklog = 128       // associations waiting for Accept
)
