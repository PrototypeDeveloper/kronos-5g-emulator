// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EventType represents ASN.1 EventType.
const (
	EventTypePresentDirect                                              Enumerated = 0
	EventTypePresentChangeOfServeCell                                   Enumerated = 1
	EventTypePresentUePresenceInAreaOfInterest                          Enumerated = 2
	EventTypePresentStopChangeOfServeCell                               Enumerated = 3
	EventTypePresentStopUePresenceInAreaOfInterest                      Enumerated = 4
	EventTypePresentCancelLocationReportingForTheUe                     Enumerated = 5
	EventTypePresentChangeOfServingCellAndUEPresenceInTheAreaOfInterest Enumerated = 6
	EventTypePresentReportTheAerialUEFlightInformation                  Enumerated = 7
	EventTypePresentCancelTheAerialUEFlightInformationReporting         Enumerated = 8
)

type EventType struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:5"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EventType) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("EventType", v) }

// DecodeEventType decodes a standalone APER value and reports the consumed octets.
func DecodeEventType(data []byte) (EventType, int, error) {
	var v EventType
	err := ngapcodec.DecodeParameter("EventType", data, &v)
	if ngapcodec.IsFatal(err) {
		return EventType{}, 0, err
	}
	return v, len(data), err
}
