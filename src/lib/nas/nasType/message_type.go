package nasType

const (
	MsgPDUSessionReleaseComplete uint8 = 0xd4
	MsgPDUSessionReleaseCommand  uint8 = 0xd3
	MsgPDUSessionReleaseReject   uint8 = 0xd2
	MsgPDUSessionReleaseRequest  uint8 = 0xd1
	MsgServiceRequest            uint8 = 0x4c
	MsgServiceAccept             uint8 = 0x4e

	MsgDeregistrationRequestUE        uint8 = 0x45
	MsgDeregistrationAcceptUE         uint8 = 0x46
	MsgRegistrationRequest            uint8 = 0x41
	MsgRegistrationAccept             uint8 = 0x42
	MsgRegistrationComplete           uint8 = 0x43
	MsgRegistrationReject             uint8 = 0x44
	MsgAuthenticationRequest          uint8 = 0x56
	MsgAuthenticationResponse         uint8 = 0x57
	MsgAuthenticationReject           uint8 = 0x58
	MsgAuthenticationFailure          uint8 = 0x59
	MsgIdentityRequest                uint8 = 0x5b
	MsgIdentityResponse               uint8 = 0x5c
	MsgSecurityModeCommand            uint8 = 0x5d
	MsgSecurityModeComplete           uint8 = 0x5e
	MsgSecurityModeReject             uint8 = 0x5f
	MsgStatus5GMM                     uint8 = 0x64
	MsgULNASTransport                 uint8 = 0x67
	MsgDLNASTransport                 uint8 = 0x68
	MsgPDUSessionEstablishmentRequest uint8 = 0xc1
	MsgPDUSessionEstablishmentAccept  uint8 = 0xc2
	MsgPDUSessionEstablishmentReject  uint8 = 0xc3
	MsgStatus5GSM                     uint8 = 0xd6

	MsgDeregistrationRequestNW                    uint8 = 0x47
	MsgDeregistrationAcceptNW                     uint8 = 0x48
	MsgServiceReject                              uint8 = 0x4d
	MsgControlPlaneServiceRequest                 uint8 = 0x4f
	MsgNetworkSliceSpecificAuthenticationCommand  uint8 = 0x50
	MsgNetworkSliceSpecificAuthenticationComplete uint8 = 0x51
	MsgNetworkSliceSpecificAuthenticationResult   uint8 = 0x52
	MsgConfigurationUpdateCommand                 uint8 = 0x54
	MsgConfigurationUpdateComplete                uint8 = 0x55
	MsgAuthenticationResult                       uint8 = 0x5a
	MsgNotification                               uint8 = 0x65
	MsgNotificationResponse                       uint8 = 0x66
	MsgPDUSessionAuthenticationCommand            uint8 = 0xc5
	MsgPDUSessionAuthenticationComplete           uint8 = 0xc6
	MsgPDUSessionAuthenticationResult             uint8 = 0xc7
	MsgPDUSessionModificationRequest              uint8 = 0xc9
	MsgPDUSessionModificationReject               uint8 = 0xca
	MsgPDUSessionModificationCommand              uint8 = 0xcb
	MsgPDUSessionModificationComplete             uint8 = 0xcc
	MsgPDUSessionModificationCommandReject        uint8 = 0xcd
)
