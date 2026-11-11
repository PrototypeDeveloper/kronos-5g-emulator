package gnb

import (
	"errors"

	"libnas/nasMessage"
	"libnas/nasType"

	"kronos/common/security"
)

// patchIdentityResponse gives the identity the IdentityRequest asked for
// (the template's type when none did): SUCI, 5G-GUTI, IMEI, 5G-S-TMSI or IMEISV.
func (u *UE) patchIdentityResponse(t *nasMessage.IdentityResponse) error {
	typ := u.identityType
	if typ == 0 && len(t.MobileIdentity) > 0 {
		typ = t.MobileIdentity[0] & identityTypeMask
	}
	switch typ {
	case nasType.MobileIdentitySUCI:
		suci, err := security.NullSUCI(u.r.mcc, u.r.mnc, u.r.ri, u.msin)
		if err != nil {
			return err
		}
		t.MobileIdentity = suci
	case nasType.MobileIdentity5GGUTI:
		if u.guti == nil {
			return errNoGUTI
		}
		t.MobileIdentity = clone(u.guti)
	case identity5GSTMSI:
		id, err := u.sTMSIIdentity()
		if err != nil {
			return err
		}
		t.MobileIdentity = id
	case identityIMEI:
		if u.r.imei != nil {
			t.MobileIdentity = clone(u.r.imei)
		}
	case identityIMEISV:
		if u.r.imeisv != nil {
			t.MobileIdentity = clone(u.r.imeisv)
		}
	}
	return nil
}

// patchAuthenticationFailure gives a synch failure the AUTS of the last
// AuthenticationRequest (TS 33.102 6.3.3); other causes carry no parameter.
func (u *UE) patchAuthenticationFailure(t *nasMessage.AuthenticationFailure) error {
	if t.Cause != causeSynchFailure {
		t.HasAuthenticationFailureParameter = false
		return nil
	}
	if u.rand == nil || u.sqn == nil {
		return errors.New("a synch failure needs the AUTS of an AuthenticationRequest, and none was received")
	}
	m, err := security.NewMilenage(u.r.k, u.r.opc)
	if err != nil {
		return err
	}
	// AUTS = SQN_MS xor AK*, then MAC-S; SQN_MS is the SQN of the request.
	akStar := m.F5Star(u.rand)
	_, macS := m.F1(u.rand, u.sqn, amfStar)
	auts := make([]byte, 0, len(u.sqn)+len(macS))
	for i := range u.sqn {
		auts = append(auts, u.sqn[i]^akStar[i])
	}
	t.HasAuthenticationFailureParameter = true
	t.AuthenticationFailureParameter.Value = append(auts, macS...)
	return nil
}
