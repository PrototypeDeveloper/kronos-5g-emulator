package gtpType

import (
	"encoding/binary"
	"fmt"
	"time"
)

const IERecoveryTimeStamp uint8 = 231

const ntpUnixOffset uint64 = 2208988800

// RecoveryTimeStamp is a TLV IE. Timestamp is the first 32 bits of the NTP
// timestamp (seconds since 1900-01-01). AdditionalData preserves optional
// octets that may follow the mandatory four-octet timestamp value.
type RecoveryTimeStamp struct {
	Header         IEHeader
	Timestamp      uint32
	AdditionalData []byte
}

func NewRecoveryTimeStamp(t time.Time) (RecoveryTimeStamp, error) {
	var r RecoveryTimeStamp
	sec := t.UTC().Unix()
	if sec < -int64(ntpUnixOffset) {
		return r, fmt.Errorf("time predates NTP epoch: %s", t.UTC())
	}
	ntp := uint64(sec + int64(ntpUnixOffset))
	if ntp > 0xffffffff {
		return r, fmt.Errorf("time exceeds 32-bit NTP seconds range: %s", t.UTC())
	}
	r.Header = IEHeader{Type: IERecoveryTimeStamp, Length: 4}
	r.Timestamp = uint32(ntp)
	return r, nil
}

func NewRecoveryTimeStampRaw(timestamp uint32) RecoveryTimeStamp {
	return RecoveryTimeStamp{
		Header:    IEHeader{Type: IERecoveryTimeStamp, Length: 4},
		Timestamp: timestamp,
	}
}

func (r RecoveryTimeStamp) Time() time.Time {
	return time.Unix(int64(r.Timestamp)-int64(ntpUnixOffset), 0).UTC()
}

func (r RecoveryTimeStamp) Encode() ([]byte, error) {
	if r.Header.Type != IERecoveryTimeStamp {
		return nil, fmt.Errorf("Recovery Time Stamp IE is not present")
	}
	valueLen := 4 + len(r.AdditionalData)
	if valueLen > 0xffff {
		return nil, fmt.Errorf("Recovery Time Stamp IE too long: %d", valueLen)
	}
	b := make([]byte, 3+valueLen)
	b[0] = IERecoveryTimeStamp
	binary.BigEndian.PutUint16(b[1:3], uint16(valueLen))
	binary.BigEndian.PutUint32(b[3:7], r.Timestamp)
	copy(b[7:], r.AdditionalData)
	return b, nil
}

func DecodeRecoveryTimeStamp(data []byte) (RecoveryTimeStamp, int, error) {
	var r RecoveryTimeStamp
	if len(data) < 3 {
		return r, 0, fmt.Errorf("Recovery Time Stamp IE is truncated")
	}
	if data[0] != IERecoveryTimeStamp {
		return r, 0, fmt.Errorf("unexpected Recovery Time Stamp IE type: %d", data[0])
	}
	length := int(binary.BigEndian.Uint16(data[1:3]))
	if length < 4 {
		return r, 0, fmt.Errorf("Recovery Time Stamp IE too short: %d", length)
	}
	total := 3 + length
	if len(data) < total {
		return r, 0, fmt.Errorf("Recovery Time Stamp IE length exceeds packet: %d", length)
	}
	r.Header = IEHeader{Type: IERecoveryTimeStamp, Length: uint16(length)}
	r.Timestamp = binary.BigEndian.Uint32(data[3:7])
	r.AdditionalData = append([]byte(nil), data[7:total]...)
	return r, total, nil
}
