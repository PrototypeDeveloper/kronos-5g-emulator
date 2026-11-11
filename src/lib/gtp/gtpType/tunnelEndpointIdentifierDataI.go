package gtpType

import (
	"encoding/binary"
	"fmt"
)

const IETunnelEndpointIdentifierDataI uint8 = 16

// TunnelEndpointIdentifierDataI is a TV IE containing the TEID copied from
// the G-PDU that triggered an Error Indication.
type TunnelEndpointIdentifierDataI struct {
	Header IEHeader
	TEID   uint32
}

func NewTunnelEndpointIdentifierDataI(teid uint32) TunnelEndpointIdentifierDataI {
	return TunnelEndpointIdentifierDataI{
		Header: IEHeader{Type: IETunnelEndpointIdentifierDataI, Length: 4},
		TEID:   teid,
	}
}

func (v TunnelEndpointIdentifierDataI) Encode() ([]byte, error) {
	if v.Header.Type != IETunnelEndpointIdentifierDataI {
		return nil, fmt.Errorf("Tunnel Endpoint Identifier Data I IE is not present")
	}
	b := make([]byte, 5)
	b[0] = IETunnelEndpointIdentifierDataI
	binary.BigEndian.PutUint32(b[1:5], v.TEID)
	return b, nil
}

func DecodeTunnelEndpointIdentifierDataI(data []byte) (TunnelEndpointIdentifierDataI, int, error) {
	var v TunnelEndpointIdentifierDataI
	if len(data) < 5 {
		return v, 0, fmt.Errorf("Tunnel Endpoint Identifier Data I IE is truncated")
	}
	if data[0] != IETunnelEndpointIdentifierDataI {
		return v, 0, fmt.Errorf("unexpected Tunnel Endpoint Identifier Data I IE type: %d", data[0])
	}
	v.Header = IEHeader{Type: IETunnelEndpointIdentifierDataI, Length: 4}
	v.TEID = binary.BigEndian.Uint32(data[1:5])
	return v, 5, nil
}
