package libnas

import (
	"fmt"
	"libnas/nasMessage"
	"libnas/nasType"
)

func encodePlain(value any) ([]byte, error) {
	switch m := value.(type) {
	case *nasMessage.PDUSessionReleaseComplete:
		return m.Encode()
	case *nasMessage.PDUSessionReleaseCommand:
		return m.Encode()
	case *nasMessage.PDUSessionReleaseReject:
		return m.Encode()
	case *nasMessage.PDUSessionReleaseRequest:
		return m.Encode()
	case *nasMessage.ServiceRequest:
		return m.Encode()
	case *nasMessage.ServiceAccept:
		return m.Encode()
	case *nasMessage.DeregistrationRequestUE:
		return m.Encode()
	case *nasMessage.DeregistrationAcceptUE:
		return m.Encode()
	case *nasMessage.RegistrationRequest:
		return m.Encode()
	case *nasMessage.RegistrationAccept:
		return m.Encode()
	case *nasMessage.RegistrationComplete:
		return m.Encode()
	case *nasMessage.RegistrationReject:
		return m.Encode()
	case *nasMessage.AuthenticationRequest:
		return m.Encode()
	case *nasMessage.AuthenticationResponse:
		return m.Encode()
	case *nasMessage.AuthenticationReject:
		return m.Encode()
	case *nasMessage.AuthenticationFailure:
		return m.Encode()
	case *nasMessage.IdentityRequest:
		return m.Encode()
	case *nasMessage.IdentityResponse:
		return m.Encode()
	case *nasMessage.SecurityModeCommand:
		return m.Encode()
	case *nasMessage.SecurityModeComplete:
		return m.Encode()
	case *nasMessage.SecurityModeReject:
		return m.Encode()
	case *nasMessage.Status5GMM:
		return m.Encode()
	case *nasMessage.ULNASTransport:
		return m.Encode()
	case *nasMessage.DLNASTransport:
		return m.Encode()
	case *nasMessage.PDUSessionEstablishmentRequest:
		return m.Encode()
	case *nasMessage.PDUSessionEstablishmentAccept:
		return m.Encode()
	case *nasMessage.PDUSessionEstablishmentReject:
		return m.Encode()
	case *nasMessage.Status5GSM:
		return m.Encode()
	case *nasMessage.DeregistrationRequestNW:
		return m.Encode()
	case *nasMessage.DeregistrationAcceptNW:
		return m.Encode()
	case *nasMessage.ServiceReject:
		return m.Encode()
	case *nasMessage.ControlPlaneServiceRequest:
		return m.Encode()
	case *nasMessage.NetworkSliceSpecificAuthenticationCommand:
		return m.Encode()
	case *nasMessage.NetworkSliceSpecificAuthenticationComplete:
		return m.Encode()
	case *nasMessage.NetworkSliceSpecificAuthenticationResult:
		return m.Encode()
	case *nasMessage.ConfigurationUpdateCommand:
		return m.Encode()
	case *nasMessage.ConfigurationUpdateComplete:
		return m.Encode()
	case *nasMessage.AuthenticationResult:
		return m.Encode()
	case *nasMessage.Notification:
		return m.Encode()
	case *nasMessage.NotificationResponse:
		return m.Encode()
	case *nasMessage.PDUSessionAuthenticationCommand:
		return m.Encode()
	case *nasMessage.PDUSessionAuthenticationComplete:
		return m.Encode()
	case *nasMessage.PDUSessionAuthenticationResult:
		return m.Encode()
	case *nasMessage.PDUSessionModificationRequest:
		return m.Encode()
	case *nasMessage.PDUSessionModificationReject:
		return m.Encode()
	case *nasMessage.PDUSessionModificationCommand:
		return m.Encode()
	case *nasMessage.PDUSessionModificationComplete:
		return m.Encode()
	case *nasMessage.PDUSessionModificationCommandReject:
		return m.Encode()
	default:
		return nil, fmt.Errorf("unsupported NAS message type: %T", value)
	}
}
func decodePlain(data []byte) (any, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	switch [2]uint8{h.ExtendedProtocolDiscriminator, h.MessageType} {
	case [2]uint8{nasType.EPD5GSM, nasType.MsgPDUSessionReleaseComplete}:
		return nasMessage.DecodePDUSessionReleaseComplete(data)
	case [2]uint8{nasType.EPD5GSM, nasType.MsgPDUSessionReleaseCommand}:
		return nasMessage.DecodePDUSessionReleaseCommand(data)
	case [2]uint8{nasType.EPD5GSM, nasType.MsgPDUSessionReleaseReject}:
		return nasMessage.DecodePDUSessionReleaseReject(data)
	case [2]uint8{nasType.EPD5GSM, nasType.MsgPDUSessionReleaseRequest}:
		return nasMessage.DecodePDUSessionReleaseRequest(data)
	case [2]uint8{nasType.EPD5GMM, nasType.MsgServiceRequest}:
		return nasMessage.DecodeServiceRequest(data)
	case [2]uint8{nasType.EPD5GMM, nasType.MsgServiceAccept}:
		return nasMessage.DecodeServiceAccept(data)
	case [2]uint8{nasType.EPD5GMM, nasType.MsgDeregistrationRequestUE}:
		m, e := nasMessage.DecodeDeregistrationRequestUE(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GMM, nasType.MsgDeregistrationAcceptUE}:
		m, e := nasMessage.DecodeDeregistrationAcceptUE(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GMM, nasType.MsgRegistrationRequest}:
		m, e := nasMessage.DecodeRegistrationRequest(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GMM, nasType.MsgRegistrationAccept}:
		m, e := nasMessage.DecodeRegistrationAccept(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GMM, nasType.MsgRegistrationComplete}:
		m, e := nasMessage.DecodeRegistrationComplete(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GMM, nasType.MsgRegistrationReject}:
		m, e := nasMessage.DecodeRegistrationReject(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GMM, nasType.MsgAuthenticationRequest}:
		m, e := nasMessage.DecodeAuthenticationRequest(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GMM, nasType.MsgAuthenticationResponse}:
		m, e := nasMessage.DecodeAuthenticationResponse(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GMM, nasType.MsgAuthenticationReject}:
		m, e := nasMessage.DecodeAuthenticationReject(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GMM, nasType.MsgAuthenticationFailure}:
		m, e := nasMessage.DecodeAuthenticationFailure(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GMM, nasType.MsgIdentityRequest}:
		m, e := nasMessage.DecodeIdentityRequest(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GMM, nasType.MsgIdentityResponse}:
		m, e := nasMessage.DecodeIdentityResponse(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GMM, nasType.MsgSecurityModeCommand}:
		m, e := nasMessage.DecodeSecurityModeCommand(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GMM, nasType.MsgSecurityModeComplete}:
		m, e := nasMessage.DecodeSecurityModeComplete(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GMM, nasType.MsgSecurityModeReject}:
		m, e := nasMessage.DecodeSecurityModeReject(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GMM, nasType.MsgStatus5GMM}:
		m, e := nasMessage.DecodeStatus5GMM(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GMM, nasType.MsgULNASTransport}:
		m, e := nasMessage.DecodeULNASTransport(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GMM, nasType.MsgDLNASTransport}:
		m, e := nasMessage.DecodeDLNASTransport(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GSM, nasType.MsgPDUSessionEstablishmentRequest}:
		m, e := nasMessage.DecodePDUSessionEstablishmentRequest(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GSM, nasType.MsgPDUSessionEstablishmentAccept}:
		m, e := nasMessage.DecodePDUSessionEstablishmentAccept(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GSM, nasType.MsgPDUSessionEstablishmentReject}:
		m, e := nasMessage.DecodePDUSessionEstablishmentReject(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GSM, nasType.MsgStatus5GSM}:
		m, e := nasMessage.DecodeStatus5GSM(data)
		if e != nil {
			return nil, e
		}
		return m, nil
	case [2]uint8{nasType.EPD5GMM, nasType.MsgDeregistrationRequestNW}:
		return decoded(nasMessage.DecodeDeregistrationRequestNW(data))
	case [2]uint8{nasType.EPD5GMM, nasType.MsgDeregistrationAcceptNW}:
		return decoded(nasMessage.DecodeDeregistrationAcceptNW(data))
	case [2]uint8{nasType.EPD5GMM, nasType.MsgServiceReject}:
		return decoded(nasMessage.DecodeServiceReject(data))
	case [2]uint8{nasType.EPD5GMM, nasType.MsgControlPlaneServiceRequest}:
		return decoded(nasMessage.DecodeControlPlaneServiceRequest(data))
	case [2]uint8{nasType.EPD5GMM, nasType.MsgNetworkSliceSpecificAuthenticationCommand}:
		return decoded(nasMessage.DecodeNetworkSliceSpecificAuthenticationCommand(data))
	case [2]uint8{nasType.EPD5GMM, nasType.MsgNetworkSliceSpecificAuthenticationComplete}:
		return decoded(nasMessage.DecodeNetworkSliceSpecificAuthenticationComplete(data))
	case [2]uint8{nasType.EPD5GMM, nasType.MsgNetworkSliceSpecificAuthenticationResult}:
		return decoded(nasMessage.DecodeNetworkSliceSpecificAuthenticationResult(data))
	case [2]uint8{nasType.EPD5GMM, nasType.MsgConfigurationUpdateCommand}:
		return decoded(nasMessage.DecodeConfigurationUpdateCommand(data))
	case [2]uint8{nasType.EPD5GMM, nasType.MsgConfigurationUpdateComplete}:
		return decoded(nasMessage.DecodeConfigurationUpdateComplete(data))
	case [2]uint8{nasType.EPD5GMM, nasType.MsgAuthenticationResult}:
		return decoded(nasMessage.DecodeAuthenticationResult(data))
	case [2]uint8{nasType.EPD5GMM, nasType.MsgNotification}:
		return decoded(nasMessage.DecodeNotification(data))
	case [2]uint8{nasType.EPD5GMM, nasType.MsgNotificationResponse}:
		return decoded(nasMessage.DecodeNotificationResponse(data))
	case [2]uint8{nasType.EPD5GSM, nasType.MsgPDUSessionAuthenticationCommand}:
		return decoded(nasMessage.DecodePDUSessionAuthenticationCommand(data))
	case [2]uint8{nasType.EPD5GSM, nasType.MsgPDUSessionAuthenticationComplete}:
		return decoded(nasMessage.DecodePDUSessionAuthenticationComplete(data))
	case [2]uint8{nasType.EPD5GSM, nasType.MsgPDUSessionAuthenticationResult}:
		return decoded(nasMessage.DecodePDUSessionAuthenticationResult(data))
	case [2]uint8{nasType.EPD5GSM, nasType.MsgPDUSessionModificationRequest}:
		return decoded(nasMessage.DecodePDUSessionModificationRequest(data))
	case [2]uint8{nasType.EPD5GSM, nasType.MsgPDUSessionModificationReject}:
		return decoded(nasMessage.DecodePDUSessionModificationReject(data))
	case [2]uint8{nasType.EPD5GSM, nasType.MsgPDUSessionModificationCommand}:
		return decoded(nasMessage.DecodePDUSessionModificationCommand(data))
	case [2]uint8{nasType.EPD5GSM, nasType.MsgPDUSessionModificationComplete}:
		return decoded(nasMessage.DecodePDUSessionModificationComplete(data))
	case [2]uint8{nasType.EPD5GSM, nasType.MsgPDUSessionModificationCommandReject}:
		return decoded(nasMessage.DecodePDUSessionModificationCommandReject(data))
	default:
		return nil, fmt.Errorf("unsupported NAS message: EPD=0x%02x type=0x%02x", h.ExtendedProtocolDiscriminator, h.MessageType)
	}
}

// decoded returns a decoded message, or a nil interface on error.
func decoded[T any](m *T, err error) (any, error) {
	if err != nil {
		return nil, err
	}
	return m, nil
}
