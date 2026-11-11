// Constants of package tun.

package tun

// tunSetIff is TUNSETIFF, _IOW('T', 202, int), as on x86, ARM and RISC-V.
const tunSetIff = 0x400454ca

// capNetAdmin is the bit of CAP_NET_ADMIN in the capability sets.
const capNetAdmin = 12

// Limits of the configuration.
const (
	minMTU             = 576   // smallest TUN MTU accepted (the datagram every IPv4 host takes)
	maxMTU             = 65535 // largest TUN MTU accepted (the largest IPv4 packet)
	firstReservedTable = 252   // the routing tables 252 to 255 are the kernel's
	lastReservedTable  = 255
	maxRuleDeletes     = 16 // most copies of the source rule removed
)

// rtnetlink values.
const (
	nlmFReplace = 0x100 // NLM_F_REPLACE

	// fib rule attributes and actions (linux/fib_rules.h)
	fraPriority = 6
	fraTable    = 15
	fraSrc      = 2
	frActToTbl  = 1

	rtprotStatic = 4
	rtScopeLink  = 253
	rtnUnicast   = 1
)

// rtnetlink messages.
const (
	nlBufLen      = 8192 // receive buffer of the acknowledgements
	fibRuleHdrLen = 12   // struct fib_rule_hdr
	hostPrefixLen = 32   // prefix length of one IPv4 address
)
