package gtpType

import "fmt"

const IERecovery uint8 = 14

// Recovery is a TV IE. The RestartCounter shall be set to zero for GTP-U Echo
// Response and ignored by the receiver; it is retained in the structure to
// represent the received wire value exactly.
type Recovery struct {
	Header         IEHeader
	RestartCounter uint8
}

func NewRecovery() Recovery {
	return Recovery{
		Header: IEHeader{Type: IERecovery, Length: 1},
	}
}

func (r Recovery) Encode() ([]byte, error) {
	if r.Header.Type != IERecovery {
		return nil, fmt.Errorf("Recovery IE is not present")
	}
	if r.RestartCounter != 0 {
		return nil, fmt.Errorf("GTP-U Recovery restart counter must be zero: %d", r.RestartCounter)
	}
	return []byte{IERecovery, 0}, nil
}

func DecodeRecovery(data []byte) (Recovery, int, error) {
	var r Recovery
	if len(data) < 2 {
		return r, 0, fmt.Errorf("Recovery IE is truncated")
	}
	if data[0] != IERecovery {
		return r, 0, fmt.Errorf("unexpected Recovery IE type: %d", data[0])
	}
	r.Header = IEHeader{Type: IERecovery, Length: 1}
	r.RestartCounter = data[1]
	return r, 2, nil
}
