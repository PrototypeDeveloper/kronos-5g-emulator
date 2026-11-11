// Data types of package gnb.

package gnb

import (
	"net/netip"
	"sync/atomic"

	libnas "libnas"
	"libngap/ngapType"

	"kronos/common/security"
	"kronos/models"
)

// The node and its subscribers.

// Role is the gNB role.
type Role struct {
	// node (NG Setup)
	name   string // RANNodeName; "" keeps the template value
	plmn   []byte
	gnbID  ngapType.BitString
	tac    []byte
	slices []ngapType.SliceSupportItem

	// subscribers (NAS mapping)
	mcc, mnc string
	nci      ngapType.BitString // 36-bit NR Cell Identity
	snn      string             // serving network name
	k, opc   []byte
	suci     security.SUCIParams // how the SUCI conceals the IMSI
	imeisv   []byte              // encoded IMEISV; nil keeps the template value
	imei     []byte              // encoded IMEI of the IMEISV (IdentityResponse)
	ea, ia   byte                // UE security capability: 5G-EA and 5G-IA octets
	nssai    []byte              // Requested NSSAI value; nil keeps the template value
	nextRAN  atomic.Int64

	// PDU sessions (N3)
	sessions    []sessionConf
	n3          netip.Addr     // gNB N3 address (gtpu.src)
	upfOverride netip.AddrPort // gtpu.dst: overrides the UPF the network names
	nextTEID    atomic.Int64

	// lastReset is the UE association list of the last NGReset the AMF sent
	// (nil: the whole interface), which NGResetAcknowledge echoes. Only the
	// node actor uses it.
	lastReset []ngapType.UEAssociatedLogicalNGConnectionItem
}

// autoList is a PDU session list the mapping fills: the list IE, the
// transfer of its items ("" without one) and the transfer fields it writes.
type autoList struct {
	list, transfer string
	fields         []string
}

// UE is the gNB-side context of one subscriber: the NGAP IDs and what the NAS
// mapping needs. It is not a UE simulator: it keeps only the
// identifiers, keys and COUNTs needed to build and check NAS messages.
type UE struct {
	r     *Role
	index int
	imsi  string
	msin  string

	ranID, amfID   int64
	hasRAN, hasAMF bool

	resStar         []byte // answer to the last AuthenticationRequest
	kamf            []byte
	ngKSI           uint8 // from the network (TSC and KSI)
	sec             *libnas.SecurityContext
	imeisvRequested bool
	guti            []byte // 5GS mobile identity value of the 5G-GUTI
	registered      bool   // RM state as the signals tell (display only)

	sessions       map[uint8]*pdu
	pendingSetup   []uint8 // sessions of the last setup request, answered by the next response
	pendingRelease []uint8 // sessions of the last release command, answered by the next response
	nextPTI        uint8

	off models.MapOff // the mapping items that are off for the signal being built

	rand, sqn     []byte                  // RAND and SQN of the last AuthenticationRequest (the AUTS of a synch failure)
	identityType  uint8                   // the identity the last IdentityRequest asked for; 0 none
	prevSec       *libnas.SecurityContext // before the last Security Mode Command (SecurityModeReject puts it back)
	pendingHO     []uint8                 // sessions of the last HandoverRequest, answered by HandoverRequestAcknowledge
	pendingModify []uint8                 // sessions of the last PDUSessionResourceModifyRequest
	sourceAMF     int64                   // AMF-UE-NGAP-ID at the handover source (PathSwitchRequest)
	hasSourceAMF  bool
}

// PDU sessions.

// pdu is the gNB-side state of one PDU session.
type pdu struct {
	id     uint8
	pti    uint8 // PTI of the procedure the UE started last
	relPTI uint8 // PTI of the last PDU Session Release Command
	ueIP   netip.Addr
	upf    netip.Addr
	ulTEID uint32
	dlTEID uint32
	qfis   []uint8
}

// sessionConf is a PDU session of config (ue.session).
type sessionConf struct {
	id      uint8
	pduType uint8  // 1 IPv4, 2 IPv6, 3 IPv4v6 (TS 24.501 §9.11.4.11)
	dnn     []byte // DNN IE value: length-prefixed labels
	snssai  []byte // S-NSSAI IE value: SST [SD]
}
