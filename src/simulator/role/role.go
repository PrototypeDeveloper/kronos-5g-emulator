// Package role defines what the engine needs from a role. The
// engine and the role packages (gnb now, an AMF role later) depend on this package,
// never on each other; cmd picks the role and hands it to the engine.
package role

import (
	"kronos/models"
	"kronos/scenario"
)

// Role is the gNB role, or the AMF role of the reversed mode (not implemented yet).
type Role interface {
	Name() string // "gnb" | "amf"
	// Listen reports whether the role accepts associations (AMF) or dials one (gNB).
	Listen() bool
	// PatchNode applies the auto-fix rules to a decoded non-UE-associated
	// template before it is sent, except those of the items off (mapoff).
	PatchNode(msg any, c models.NodeBuild) error
	// NodeReceived runs the built-in processing of a received
	// non-UE-associated message. It updates st and returns a
	// line for the event log, or "" when there is nothing to report.
	NodeReceived(msg any, st *models.NodeStatus) string
	// NewUE creates the protocol context of the index-th subscriber.
	NewUE(index int, imsi string) UE
	// AutoField reports whether the role's auto-fix replaces a field of a
	// template when it is sent (the signal editor marks it), with
	// a note for a field it changes in part or only sometimes.
	AutoField(f scenario.FieldRef) (bool, string)
}

// UE is the per-subscriber protocol context: the NGAP IDs and the NAS
// mapping. The engine calls it from the subscriber's actor only, so it needs
// no locking.
type UE interface {
	// Build turns a UE-associated template into the PDU to send: auto-fix of
	// NGAP and NAS (except the items in off), then NAS protection.
	Build(sig *scenario.Signal, off models.MapOff) (models.Built, error)
	// Receive runs the built-in processing of a received UE-associated PDU.
	Receive(msg any) models.Inbound
	// HandoverContext returns the mapping context the target of a handover
	// needs (sync_send(HO)); AdoptHandover takes one over (sync_receive(HO)).
	HandoverContext() models.HOContext
	AdoptHandover(c models.HOContext) (models.Inbound, error)
	// RANID and AMFID return the NGAP IDs of the current UE association.
	RANID() (int64, bool)
	AMFID() (int64, bool)
	// ReleaseNGAP forgets the NGAP IDs (UE context release). NAS state stays.
	ReleaseNGAP()
	// Reset forgets everything (the SCTP association went down).
	Reset()
	// Status fills the role's part of the subscriber snapshot: RM/CM state,
	// identities, NAS security and PDU sessions.
	Status(st *models.UEStatus)
}
