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
	default:
		return nil, fmt.Errorf("unsupported NAS message: EPD=0x%02x type=0x%02x", h.ExtendedProtocolDiscriminator, h.MessageType)
	}
}
