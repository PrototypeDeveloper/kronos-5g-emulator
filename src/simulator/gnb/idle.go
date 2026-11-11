package gnb

import (
	"errors"

	libnas "libnas"
	"libnas/nasMessage"
	"libnas/nasType"
	nassec "libnas/security"
	"libngap/ngapType"

	"kronos/common/security"
	"kronos/models"
	"kronos/scenario"
)

// errNoGUTI is returned when a message needs the 5G-GUTI before registration.
var errNoGUTI = errors.New("no 5G-GUTI (not registered yet)")

// sTMSI splits the 5G-GUTI (5GS mobile identity value: 0xf2, PLMN, AMF region
// ID, AMF set ID and pointer, 5G-TMSI) into the parts of the 5G-S-TMSI.
func (u *UE) sTMSI() (setID, pointer uint64, tmsi []byte, ok bool) {
	if len(u.guti) != gutiLen {
		return 0, 0, nil, false
	}
	setID = uint64(u.guti[5])<<2 | uint64(u.guti[6]>>6)
	pointer = uint64(u.guti[6] & 0x3f)
	return setID, pointer, u.guti[7:11], true
}

// pagingKey is how the engine finds this UE when the network pages it.
func (u *UE) pagingKey() string {
	setID, pointer, tmsi, ok := u.sTMSI()
	if !ok {
		return ""
	}
	return scenario.PagingKey(setID, pointer, tmsi)
}

// sTMSIIdentity is the 5GS mobile identity value of the 5G-S-TMSI (type 4).
func (u *UE) sTMSIIdentity() ([]byte, error) {
	if len(u.guti) != gutiLen {
		return nil, errNoGUTI
	}
	return append([]byte{sTMSIIdentityOctet}, u.guti[5:11]...), nil
}

// patchFiveGSTMSI writes the 5G-S-TMSI into the InitialUEMessage IE. It
// reports false when the UE has no 5G-GUTI.
func (u *UE) patchFiveGSTMSI(f *ngapType.FiveGSTMSI) bool {
	setID, pointer, tmsi, ok := u.sTMSI()
	if !ok {
		return false
	}
	f.AMFSetID.Value = BitString(setID, amfSetIDBits)
	f.AMFPointer.Value = BitString(pointer, amfPointerBits)
	f.FiveGTMSI.Value = clone(tmsi)
	return true
}

// ksi is the ngKSI the UE puts in its messages: the network's, or "no key".
func (u *UE) ksi() uint8 {
	if u.sec == nil {
		return noKey
	}
	return u.ngKSI & ksiMask
}

// sessionStatus is the PDU session status bitmap of the established sessions
// (TS 24.501 §9.11.3.44: PSI(n) is bit n of the two octets).
func (u *UE) sessionStatus() []byte {
	b := make([]byte, 2)
	for _, id := range u.established() {
		if id < 8 {
			b[0] |= 1 << id
		} else if id < 16 {
			b[1] |= 1 << (id - 8)
		}
	}
	return b
}

// patchServiceRequest writes the ngKSI (low half octet), the 5G-S-TMSI and the
// session status IEs. The NAS message container repeats the whole message.
func (u *UE) patchServiceRequest(t *nasMessage.ServiceRequest) error {
	if u.off.On(models.MapNgKSI) {
		t.ServiceTypeAndNgKSI = t.ServiceTypeAndNgKSI&highHalfOctet | u.ksi()
	}
	if u.off.On(models.MapIdentity) {
		id, err := u.sTMSIIdentity()
		if err != nil {
			return err
		}
		t.FiveGSTMSI = id
	}
	if u.off.On(models.MapStatus) {
		status := u.sessionStatus()
		if t.HasPDUSessionStatus {
			t.PDUSessionStatus = clone(status)
		}
		if t.HasUplinkDataStatus {
			t.UplinkDataStatus = clone(status)
		}
		if t.HasAllowedPDUSessionStatus {
			t.AllowedPDUSessionStatus = clone(status)
		}
	}
	if !t.HasNASMessageContainer || !u.off.On(models.MapContainer) {
		return nil
	}
	if u.sec == nil {
		t.HasNASMessageContainer = false // only sent under a security context
		return nil
	}
	inner, err := libnas.Decode(t.NASMessageContainer, nil)
	if err != nil {
		return err
	}
	if in, ok := inner.(*nasMessage.ServiceRequest); ok {
		in.HasNASMessageContainer = false
		if err := u.patchServiceRequest(in); err != nil {
			return err
		}
	}
	t.NASMessageContainer, err = libnas.Encode(inner, nil)
	return err
}

// patchDeregistration writes the ngKSI (high half octet) and the identity.
func (u *UE) patchDeregistration(t *nasMessage.DeregistrationRequestUE) error {
	if u.off.On(models.MapNgKSI) {
		t.DeregistrationTypeAndNgKSI = u.ksi()<<ngKSIShift | t.DeregistrationTypeAndNgKSI&lowHalfOctet
	}
	if len(t.MobileIdentity) == 0 || !u.off.On(models.MapIdentity) {
		return nil
	}
	switch t.MobileIdentity[0] & identityTypeMask {
	case nasType.MobileIdentity5GGUTI:
		if u.guti == nil {
			return errNoGUTI
		}
		t.MobileIdentity = clone(u.guti)
	case nasType.MobileIdentitySUCI:
		suci, err := security.NullSUCI(u.r.mcc, u.r.mnc, u.r.ri, u.msin)
		if err != nil {
			return err
		}
		t.MobileIdentity = suci
	}
	return nil
}

// patchRegistrationContainer rebuilds the complete RegistrationRequest that
// a registration under a security context carries in its container (the
// same values as the outer message, ngKSI included).
func (u *UE) patchRegistrationContainer(t *nasMessage.RegistrationRequest) error {
	if !t.HasNASMessageContainer {
		return nil
	}
	if u.sec == nil {
		t.HasNASMessageContainer = false
		return nil
	}
	inner, err := libnas.Decode(t.NASMessageContainer.Value, nil)
	if err != nil {
		return err
	}
	if in, ok := inner.(*nasMessage.RegistrationRequest); ok {
		in.HasNASMessageContainer = false
		u.patchRegistrationRequest(in, false)
	}
	t.NASMessageContainer.Value, err = libnas.Encode(inner, nil)
	return err
}

// cipherContainer ciphers the NAS message container of an initial NAS
// message sent under a security context. The container uses the COUNT of
// the integrity-protected message that carries it (TS 24.501 §4.4.6), i.e.
// the current uplink COUNT, which libnas.Encode then uses and advances.
func (u *UE) cipherContainer(m any) error {
	var c *[]byte
	switch t := m.(type) {
	case *nasMessage.ServiceRequest:
		if t.HasNASMessageContainer {
			c = &t.NASMessageContainer
		}
	case *nasMessage.RegistrationRequest:
		if t.HasNASMessageContainer {
			c = &t.NASMessageContainer.Value
		}
	}
	if c == nil {
		return nil
	}
	bearer := u.sec.NASConnectionID // as libnas.Encode: 1 (3GPP access) unless set
	if bearer == 0 {
		bearer = bearer3GPP
	}
	out, err := nassec.Cipher(u.sec.CipheringAlgorithm, u.sec.KnasEnc, uint32(u.sec.UplinkCount), bearer, uint8(libnas.Uplink), *c)
	if err != nil {
		return err
	}
	*c = out
	return nil
}
