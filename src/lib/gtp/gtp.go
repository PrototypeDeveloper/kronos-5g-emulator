package libgtp

import (
	"fmt"

	"libgtp/gtpMessage"
	"libgtp/gtpType"
)

func Encode(i any) ([]byte, error) {
	switch msg := i.(type) {
	case *gtpMessage.EchoRequest:
		return msg.Encode()
	case *gtpMessage.EchoResponse:
		return msg.Encode()
	case *gtpMessage.ErrorIndication:
		return msg.Encode()
	case *gtpMessage.EndMarker:
		return msg.Encode()
	case *gtpMessage.SupportedExtensionHeadersNotification:
		return msg.Encode()
	case *gtpMessage.GPDU:
		return msg.Encode()
	default:
		return nil, fmt.Errorf("unsupported GTPv1-U message type: %T", i)
	}
}

func Decode(data []byte) (any, error) {
	header, err := gtpType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}

	switch header.MessageType {
	case gtpType.MsgEchoRequest:
		return gtpMessage.DecodeEchoRequest(data)
	case gtpType.MsgEchoResponse:
		return gtpMessage.DecodeEchoResponse(data)
	case gtpType.MsgErrorIndication:
		return gtpMessage.DecodeErrorIndication(data)
	case gtpType.MsgEndMarker:
		return gtpMessage.DecodeEndMarker(data)
	case gtpType.MsgSupportedExtensionHeadersNotification:
		return gtpMessage.DecodeSupportedExtensionHeadersNotification(data)
	case gtpType.MsgGPDU:
		return gtpMessage.DecodeGPDU(data)
	default:
		return nil, fmt.Errorf("unsupported GTPv1-U message type: %d", header.MessageType)
	}
}
