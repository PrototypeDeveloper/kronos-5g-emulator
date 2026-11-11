// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ScheduledCommunicationTime represents ASN.1 ScheduledCommunicationTime.
type ScheduledCommunicationTime struct {
	HasDayofWeek      bool
	DayofWeek         BitString `aper:"sizeLB:7,sizeUB:7"`
	HasTimeofDayStart bool
	TimeofDayStart    int64 `aper:"valueExt,valueLB:0,valueUB:86399"`
	HasTimeofDayEnd   bool
	TimeofDayEnd      int64 `aper:"valueExt,valueLB:0,valueUB:86399"`
	HasIEExtensions   bool
	IEExtensions      []ScheduledCommunicationTimeExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ScheduledCommunicationTime) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ScheduledCommunicationTime", v)
}

// DecodeScheduledCommunicationTime decodes a standalone APER value and reports the consumed octets.
func DecodeScheduledCommunicationTime(data []byte) (ScheduledCommunicationTime, int, error) {
	var v ScheduledCommunicationTime
	err := ngapcodec.DecodeParameter("ScheduledCommunicationTime", data, &v)
	if ngapcodec.IsFatal(err) {
		return ScheduledCommunicationTime{}, 0, err
	}
	return v, len(data), err
}
