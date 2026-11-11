package scenario

import "libnas/nasMessage"

// nasMessageTypes lists the NAS message types that libnas encodes, so the
// signal editor can create them. TestNASMessageTable checks it against the
// supported entries of nasTable.
var nasMessageTypes = []any{
	(*nasMessage.AuthenticationFailure)(nil),
	(*nasMessage.AuthenticationReject)(nil),
	(*nasMessage.AuthenticationRequest)(nil),
	(*nasMessage.AuthenticationResponse)(nil),
	(*nasMessage.DeregistrationAcceptUE)(nil),
	(*nasMessage.DeregistrationRequestUE)(nil),
	(*nasMessage.DLNASTransport)(nil),
	(*nasMessage.IdentityRequest)(nil),
	(*nasMessage.IdentityResponse)(nil),
	(*nasMessage.PDUSessionEstablishmentAccept)(nil),
	(*nasMessage.PDUSessionEstablishmentReject)(nil),
	(*nasMessage.PDUSessionEstablishmentRequest)(nil),
	(*nasMessage.PDUSessionReleaseCommand)(nil),
	(*nasMessage.PDUSessionReleaseComplete)(nil),
	(*nasMessage.PDUSessionReleaseReject)(nil),
	(*nasMessage.PDUSessionReleaseRequest)(nil),
	(*nasMessage.RegistrationAccept)(nil),
	(*nasMessage.RegistrationComplete)(nil),
	(*nasMessage.RegistrationReject)(nil),
	(*nasMessage.RegistrationRequest)(nil),
	(*nasMessage.SecurityModeCommand)(nil),
	(*nasMessage.SecurityModeComplete)(nil),
	(*nasMessage.SecurityModeReject)(nil),
	(*nasMessage.ServiceAccept)(nil),
	(*nasMessage.ServiceRequest)(nil),
	(*nasMessage.Status5GMM)(nil),
	(*nasMessage.Status5GSM)(nil),
	(*nasMessage.ULNASTransport)(nil),
}
