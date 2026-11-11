// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PeriodicRegistrationUpdateTimer represents ASN.1 PeriodicRegistrationUpdateTimer.
type PeriodicRegistrationUpdateTimer struct {
	Value BitString `aper:"sizeLB:8,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PeriodicRegistrationUpdateTimer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PeriodicRegistrationUpdateTimer", v)
}

// DecodePeriodicRegistrationUpdateTimer decodes a standalone APER value and reports the consumed octets.
func DecodePeriodicRegistrationUpdateTimer(data []byte) (PeriodicRegistrationUpdateTimer, int, error) {
	var v PeriodicRegistrationUpdateTimer
	err := ngapcodec.DecodeParameter("PeriodicRegistrationUpdateTimer", data, &v)
	if ngapcodec.IsFatal(err) {
		return PeriodicRegistrationUpdateTimer{}, 0, err
	}
	return v, len(data), err
}
