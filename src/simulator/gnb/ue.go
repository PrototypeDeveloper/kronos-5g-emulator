package gnb

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"

	libnas "libnas"
	"libnas/nasMessage"
	"libnas/nasType"
	nassec "libnas/security"
	libngap "libngap"
	"libngap/ngapMessage"
	"libngap/ngapType"

	"kronos/common/security"
	"kronos/models"
	"kronos/scenario"
	"kronos/simulator/role"
)

var _ role.UE = (*UE)(nil)

// RANID implements role.UE.
func (u *UE) RANID() (int64, bool) { return u.ranID, u.hasRAN }

// AMFID implements role.UE.
func (u *UE) AMFID() (int64, bool) { return u.amfID, u.hasAMF }

// ReleaseNGAP implements role.UE: the UE goes to CM-IDLE. NAS state and the
// PDU sessions stay; their N3 tunnels are set up again on reactivation.
func (u *UE) ReleaseNGAP() {
	u.hasRAN, u.hasAMF = false, false
	u.pendingSetup, u.pendingRelease = nil, nil
}

// deregister: the UE leaves 5GMM-REGISTERED and releases its PDU sessions
// locally (TS 24.501 5.5.2.2 and 5.5.2.3). With in, the N3 tunnels of the
// sessions are reported down and the release is noted; on a send (switch
// off), the engine sees the sessions gone after the message.
func (u *UE) deregister(in *models.Inbound) {
	u.registered = false
	if len(u.sessions) == 0 {
		return
	}
	ids := make([]int, 0, len(u.sessions))
	for id, s := range u.sessions {
		ids = append(ids, int(id))
		if in != nil && s.dlTEID != 0 {
			in.TunnelsDown = append(in.TunnelsDown, s.dlTEID)
		}
	}
	sort.Ints(ids)
	u.sessions = nil
	u.pendingSetup, u.pendingRelease = nil, nil
	if in != nil {
		in.Notes = append(in.Notes, fmt.Sprintf("PDU session(s) %s released locally (deregistered)", strings.Trim(fmt.Sprint(ids), "[]")))
	}
}

// Reset implements role.UE.
func (u *UE) Reset() {
	*u = UE{r: u.r, index: u.index, imsi: u.imsi, msin: u.msin, ngKSI: noKey}
}

// present reports whether a field of a decoded NGAP message is present (an
// optional IE has a Has<Name> flag).
func present(v reflect.Value, name string) bool {
	if !v.FieldByName(name).IsValid() {
		return false
	}
	has := v.FieldByName("Has" + name)
	return !has.IsValid() || has.Bool()
}

// Build implements role.UE: decode a fresh copy of the
// template, fix the NGAP IDs and the location, rebuild and protect the NAS.
func (u *UE) Build(sig *scenario.Signal, off models.MapOff) (models.Built, error) {
	u.off = off
	if sig.Mode == scenario.ModeRaw {
		return models.Built{PDU: sig.Bytes}, nil
	}
	b, plain, err := u.build(sig)
	if err != nil {
		return models.Built{}, err
	}
	return models.Built{PDU: b, Plain: plain}, nil
}

// build fixes a decoded copy of the template for this UE and encodes it; it
// also returns the plaintext of the NAS-PDU it protects.
func (u *UE) build(sig *scenario.Signal) ([]byte, [][]byte, error) {
	msg, err := libngap.Decode(sig.Bytes)
	if msg == nil {
		return nil, nil, err
	}
	v := reflect.ValueOf(msg).Elem()
	_, initial := msg.(*ngapMessage.InitialUEMessage)
	ids := u.off.On(models.MapNGAPID)
	_, pathSwitch := msg.(*ngapMessage.PathSwitchRequest)
	if initial || pathSwitch {
		// A new UE association: new RAN-UE-NGAP-ID, the AMF's ID is learned
		// again. With ngapid off the template's ID is sent, and answers come to it.
		u.ranID, u.hasRAN, u.hasAMF = u.r.nextRAN.Add(1), true, false
		if !ids {
			u.ranID = v.FieldByName("RANUENGAPID").Interface().(ngapType.RANUENGAPID).Value
		}
	}
	// The NGAP IDs and the location of this UE association.
	if present(v, "RANUENGAPID") && ids {
		if !u.hasRAN {
			return nil, nil, errors.New("no RAN-UE-NGAP-ID (send an InitialUEMessage first)")
		}
		v.FieldByName("RANUENGAPID").Addr().Interface().(*ngapType.RANUENGAPID).Value = u.ranID
	}
	if present(v, "AMFUENGAPID") && ids {
		if !u.hasAMF {
			return nil, nil, errors.New("no AMF-UE-NGAP-ID received yet")
		}
		v.FieldByName("AMFUENGAPID").Addr().Interface().(*ngapType.AMFUENGAPID).Value = u.amfID
	}
	if present(v, "UserLocationInformation") {
		u.r.patchULI(v.FieldByName("UserLocationInformation").Addr().Interface().(*ngapType.UserLocationInformation), u.off)
	}
	if m, ok := msg.(*ngapMessage.InitialUEMessage); ok && m.HasFiveGSTMSI && u.off.On(models.MapIdentity) {
		// A UE without a 5G-GUTI has no 5G-S-TMSI to give the gNB.
		m.HasFiveGSTMSI = u.patchFiveGSTMSI(&m.FiveGSTMSI)
	}
	// The NAS-PDU, rebuilt and protected (a NASNonDeliveryIndication returns
	// a downlink one as it came).
	var plain [][]byte
	_, nonDelivery := msg.(*ngapMessage.NASNonDeliveryIndication)
	if present(v, "NASPDU") && !nonDelivery {
		pdu := v.FieldByName("NASPDU").Addr().Interface().(*ngapType.NASPDU)
		var p []byte
		if pdu.Value, p, err = u.buildNAS(pdu.Value, sig.Mode, initial); err != nil {
			return nil, nil, err
		}
		plain = [][]byte{p} // the top-level NAS-PDU comes first in scenario.NASPDUs
	}
	// The PDU session lists of the responses.
	if err := u.patchSetupLists(msg); err != nil {
		return nil, nil, err
	}
	if err := u.patchReleaseLists(msg); err != nil {
		return nil, nil, err
	}
	if err := u.patchProcedures(msg); err != nil {
		return nil, nil, err
	}
	b, err := libngap.Encode(msg)
	return b, plain, err
}

// patchULI writes the NR cell and TA of config (the items whose mapping is on).
func (r *Role) patchULI(uli *ngapType.UserLocationInformation, off models.MapOff) {
	if off.On(models.MapPLMN) || off.On(models.MapNCI) || off.On(models.MapTAC) {
		uli.Present = ngapType.UserLocationInformationPresentUserLocationInformationNR
	}
	nr := &uli.UserLocationInformationNR
	if off.On(models.MapPLMN) {
		nr.NRCGI.PLMNIdentity.Value = clone(r.plmn)
		nr.TAI.PLMNIdentity.Value = clone(r.plmn)
	}
	if off.On(models.MapNCI) {
		nr.NRCGI.NRCellIdentity.Value = ngapType.BitString{Bytes: clone(r.nci.Bytes), BitLength: r.nci.BitLength}
	}
	if off.On(models.MapTAC) {
		nr.TAI.TAC.Value = clone(r.tac)
	}
	if nr.HasTimeStamp && off.On(models.MapTimeStamp) {
		// The first 32 bits of an NTP timestamp: seconds since 1900 (TS 38.413 §9.3.1.75).
		nr.TimeStamp.Value = binary.BigEndian.AppendUint32(nil, uint32(time.Now().Unix()+ntpEpochOffset))
	}
}

// buildNAS rebuilds the NAS-PDU of a template and protects it. It also
// returns the plaintext of a protected NAS-PDU (nil when sent in clear).
func (u *UE) buildNAS(b []byte, mode scenario.Mode, initial bool) ([]byte, []byte, error) {
	plain := b
	if len(b) > protectedHeaderLen && b[0] == models.EPD5GMM && b[1]&shtMask != 0 {
		plain = b[protectedHeaderLen:] // a template captured with NEA0: the body is readable
	}
	m, err := libnas.Decode(plain, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot decode the NAS of the template: %v", err)
	}
	if mode == scenario.ModeAuto {
		if err := u.patchNAS(m, false); err != nil {
			return nil, nil, err
		}
	}
	hdr, err := u.protection(m, initial)
	if err != nil {
		return nil, nil, err
	}
	if hdr == libnas.SecurityHeaderIntegrityProtected && u.off.On(models.MapContainer) {
		if err := u.cipherContainer(m); err != nil {
			return nil, nil, err
		}
	}
	if hdr == libnas.SecurityHeaderPlain {
		out, err := libnas.Encode(m, nil)
		return out, nil, err
	}
	var seen []byte
	out, err := libnas.Encode(m, &libnas.EncodeContext{Security: u.sec, SecurityHeaderType: hdr, Direction: libnas.Uplink,
		Observe: func(_, p []byte) { seen = clone(p) }})
	if err != nil {
		return nil, nil, err
	}
	if t, ok := m.(*nasMessage.DeregistrationRequestUE); ok && t.DeregistrationTypeAndNgKSI&switchOffBit != 0 {
		u.deregister(nil) // switch off: the network answers with the release only
	}
	return out, seen, nil
}

// protection picks the security header type.
func (u *UE) protection(m any, initial bool) (libnas.SecurityHeaderType, error) {
	switch m.(type) {
	case *nasMessage.SecurityModeComplete:
		if u.sec == nil {
			return 0, errors.New("no Security Mode Command received, so SecurityModeComplete cannot be protected")
		}
		return libnas.SecurityHeaderIntegrityProtectedAndCipheredWithNewContext, nil
	case *nasMessage.SecurityModeReject:
		// The UE does not take the new context into use: the reject goes in
		// clear and the context before the command comes back.
		u.sec = u.prevSec
		return libnas.SecurityHeaderPlain, nil
	}
	switch {
	case u.sec == nil:
		return libnas.SecurityHeaderPlain, nil
	case initial:
		// An initial NAS message under a security context is integrity
		// protected only; the container inside it is ciphered (TS 24.501 §4.4.6).
		return libnas.SecurityHeaderIntegrityProtected, nil
	default:
		return libnas.SecurityHeaderIntegrityProtectedAndCiphered, nil
	}
}

// patchNAS writes the dynamic NAS fields. inContainer is set
// for the RegistrationRequest inside a SecurityModeComplete.
func (u *UE) patchNAS(m any, inContainer bool) error {
	switch t := m.(type) {
	case *nasMessage.RegistrationRequest:
		u.patchRegistrationRequest(t, inContainer)
		if !inContainer && u.off.On(models.MapContainer) {
			return u.patchRegistrationContainer(t)
		}
	case *nasMessage.ServiceRequest:
		return u.patchServiceRequest(t)
	case *nasMessage.DeregistrationRequestUE:
		return u.patchDeregistration(t)
	case *nasMessage.IdentityResponse:
		if u.off.On(models.MapIdentity) {
			return u.patchIdentityResponse(t)
		}
	case *nasMessage.AuthenticationFailure:
		if u.off.On(models.MapRES) {
			return u.patchAuthenticationFailure(t)
		}
	case *nasMessage.AuthenticationResponse:
		if !u.off.On(models.MapRES) {
			return nil
		}
		if u.resStar == nil {
			return errors.New("no RES* (no AuthenticationRequest received, or the authentication failed)")
		}
		t.HasAuthenticationResponseParameter = true
		t.AuthenticationResponseParameter.Value = clone(u.resStar)
	case *nasMessage.ULNASTransport:
		return u.patchULNASTransport(t)
	case *nasMessage.SecurityModeComplete:
		if u.off.On(models.MapIMEISV) {
			had := t.HasIMEISV
			t.HasIMEISV = false
			if u.imeisvRequested {
				if u.r.imeisv != nil {
					t.HasIMEISV, t.IMEISV.Value = true, clone(u.r.imeisv)
				} else {
					t.HasIMEISV = had // keep the template's IMEISV
				}
			}
		}
		if t.HasNASMessageContainer && u.off.On(models.MapContainer) {
			inner, err := libnas.Decode(t.NASMessageContainer.Value, nil)
			if err != nil {
				return fmt.Errorf("cannot decode the NAS message container: %v", err)
			}
			if err := u.patchNAS(inner, true); err != nil {
				return err
			}
			if t.NASMessageContainer.Value, err = libnas.Encode(inner, nil); err != nil {
				return err
			}
		}
	}
	return nil
}

// patchRegistrationRequest writes the ngKSI, the identity, the security
// capability, the NSSAI and the PDU session status IEs.
func (u *UE) patchRegistrationRequest(t *nasMessage.RegistrationRequest, inContainer bool) {
	if !inContainer && u.off.On(models.MapNgKSI) { // the container repeats the initial message as it was sent
		ksi := uint8(noKey)
		if u.sec != nil {
			ksi = u.ngKSI & ksiMask
		}
		t.RegistrationTypeAndNgKSI = ksi<<ngKSIShift | t.RegistrationTypeAndNgKSI&lowHalfOctet
	}
	// The identity type of the template decides SUCI or 5G-GUTI.
	if len(t.MobileIdentity) > 0 && u.off.On(models.MapIdentity) {
		switch t.MobileIdentity[0] & identityTypeMask {
		case nasType.MobileIdentitySUCI:
			if suci, err := security.NullSUCI(u.r.mcc, u.r.mnc, u.r.ri, u.msin); err == nil {
				t.MobileIdentity = suci
			}
		case nasType.MobileIdentity5GGUTI:
			if u.guti != nil {
				t.MobileIdentity = clone(u.guti)
			}
		}
	}
	if t.HasUESecurityCapability && len(t.UESecurityCapability.Value) >= secCapLen && u.off.On(models.MapSecCap) {
		v := clone(t.UESecurityCapability.Value)
		v[0], v[1] = u.r.ea, u.r.ia
		t.UESecurityCapability.Value = v
	}
	if t.HasRequestedNSSAI && u.r.nssai != nil && u.off.On(models.MapSNSSAI) {
		t.RequestedNSSAI.Value = clone(u.r.nssai)
	}
	if !u.off.On(models.MapStatus) {
		return
	}
	status := u.sessionStatus()
	if t.HasPDUSessionStatus {
		t.PDUSessionStatus.Value = clone(status)
	}
	if t.HasUplinkDataStatus {
		t.UplinkDataStatus.Value = clone(status)
	}
	if t.HasAllowedPDUSessionStatus {
		t.AllowedPDUSessionStatus.Value = clone(status)
	}
}

// Receive implements role.UE.
func (u *UE) Receive(msg any) models.Inbound {
	var in models.Inbound
	if _, _, amf, ok := scenario.UEIDs(msg); ok {
		u.amfID, u.hasAMF = amf, true
	}
	if _, ok := msg.(*ngapMessage.UEContextReleaseCommand); ok {
		in.Release = true
	}
	u.receiveNGAP(msg, &in)
	u.receiveProcedures(msg, &in)
	for _, b := range scenario.NASPDUs(msg) {
		u.receiveNAS(b, &in)
	}
	u.tunnelsUp(&in)
	return in
}

// receiveNAS checks, names and decodes one downlink NAS-PDU and updates the
// UE from it.
func (u *UE) receiveNAS(b []byte, in *models.Inbound) {
	slot := len(in.Plain) // one entry per NAS-PDU, in order
	in.Plain = append(in.Plain, nil)
	fail := func(format string, args ...any) { in.Errors = append(in.Errors, fmt.Sprintf(format, args...)) }
	if len(b) < 3 {
		fail("NAS too short (%d octets)", len(b))
		return
	}
	// Remove the security header, checking the MAC and deciphering.
	plain, hdr := b, uint8(0)
	if b[0] == models.EPD5GMM && b[1]&shtMask != 0 {
		hdr = b[1] & shtMask
		if len(b) < 10 {
			fail("protected NAS too short (%d octets)", len(b))
			return
		}
		if hdr == shtIntegrityNewContext {
			plain = b[protectedHeaderLen:] // Security Mode Command: verified with the new keys below
		} else {
			if u.sec == nil {
				fail("protected NAS (type %d) received without a NAS security context", hdr)
				return
			}
			var err error
			if plain, err = openDL(u.sec, b); err != nil {
				fail("%v", err)
				return
			}
		}
		in.Plain[slot] = clone(plain)
	}
	// Name the message, then decode it and act on it.
	info, ok := nasInfo(plain)
	if !ok {
		fail("unknown NAS message type (%x)", plain[:min(len(plain), 4)])
		return
	}
	names := []string{info.Name}
	if info.Name == "DeregistrationRequestNW" {
		u.deregister(in) // network initiated de-registration
	}
	if !info.Supported {
		in.NAS = append(in.NAS, names...)
		in.Notes = append(in.Notes, info.Name+" is recognized by name only (libnas cannot decode its body)")
		return
	}
	m, err := libnas.Decode(plain, nil)
	if err != nil {
		in.NAS = append(in.NAS, names...)
		fail("cannot decode %s: %v", info.Name, err)
		return
	}
	switch t := m.(type) {
	case *nasMessage.AuthenticationRequest:
		u.authenticate(t, in)
	case *nasMessage.SecurityModeCommand:
		if hdr != shtIntegrityNewContext {
			fail("the security header type of Security Mode Command is %d (3 expected)", hdr)
			break
		}
		u.securityMode(t, b, in)
	case *nasMessage.RegistrationAccept:
		u.registered = true
		if t.HasGUTI5GS {
			u.guti = clone(t.GUTI5GS.Value)
			in.PagingKey = u.pagingKey()
			in.Notes = append(in.Notes, fmt.Sprintf("Registration Accept: 5G-GUTI %x", u.guti))
		} else {
			in.Notes = append(in.Notes, "Registration Accept")
		}
	case *nasMessage.ServiceAccept:
		in.Notes = append(in.Notes, "Service Accept (connected again)")
	case *nasMessage.DeregistrationAcceptUE:
		in.Notes = append(in.Notes, "Deregistration Accept")
		u.deregister(in)
	case *nasMessage.RegistrationReject:
		in.Notes = append(in.Notes, "registration rejected (RegistrationReject)")
		u.deregister(in)
	case *nasMessage.AuthenticationReject:
		// The UE deletes its 5G-GUTI, ngKSI and security context (TS 24.501 5.4.1.3.5).
		u.resStar, u.sec, u.guti, u.ngKSI = nil, nil, nil, noKey
		in.Notes = append(in.Notes, "authentication rejected (AuthenticationReject): the 5G-GUTI and the NAS security context are deleted")
		u.deregister(in)
	case *nasMessage.IdentityRequest:
		u.identityType = t.IdentityType & identityTypeMask
		in.Notes = append(in.Notes, fmt.Sprintf("Identity Request: identity type %d", u.identityType))
	case *nasMessage.DLNASTransport:
		if c := t.PayloadContainer; t.PayloadContainerType == payloadN1SMInfo && len(c) >= 4 {
			if inner, ok := scenario.NASByType(c[0], c[3]); ok {
				names = []string{inner.Name, info.Name} // the 5GSM name matches first
				u.receive5GSM(c, in)
			}
		}
	}
	in.NAS = append(in.NAS, names...)
}

// nasInfo looks up the type of a plain 5GMM or 5GSM message.
func nasInfo(plain []byte) (*models.NASInfo, bool) {
	switch {
	case len(plain) >= 3 && plain[0] == models.EPD5GMM && plain[1] == 0:
		return scenario.NASByType(models.EPD5GMM, plain[2])
	case len(plain) >= 4 && plain[0] == models.EPD5GSM:
		return scenario.NASByType(models.EPD5GSM, plain[3])
	}
	return nil, false
}

// authenticate checks AUTN and derives RES* and K_AMF (5G AKA, TS 33.501 §6.1.3.2).
func (u *UE) authenticate(t *nasMessage.AuthenticationRequest, in *models.Inbound) {
	u.resStar = nil
	if !t.HasAuthenticationParameterRAND || !t.HasAuthenticationParameterAUTN ||
		len(t.AuthenticationParameterRAND.Value) != randAUTNLen || len(t.AuthenticationParameterAUTN.Value) != randAUTNLen {
		in.Errors = append(in.Errors, "authentication: no RAND / AUTN (EAP-AKA' is not supported)")
		return
	}
	rand, autn := t.AuthenticationParameterRAND.Value, t.AuthenticationParameterAUTN.Value
	u.rand, u.sqn = clone(rand), nil
	m, err := security.NewMilenage(u.r.k, u.r.opc)
	if err != nil {
		in.Errors = append(in.Errors, "authentication: "+err.Error())
		return
	}
	res, ck, ik, ak := m.F2345(rand)
	sqnXorAK, amf, mac := autn[:6], autn[6:8], autn[8:]
	sqn := make([]byte, 6)
	for i := range sqn {
		sqn[i] = sqnXorAK[i] ^ ak[i]
	}
	u.sqn = sqn // SQN_MS of a synch failure (AUTS)
	if xmac, _ := m.F1(rand, sqn, amf); !bytes.Equal(xmac, mac) {
		in.Errors = append(in.Errors, "authentication: the MAC of AUTN does not match (K or OPc may differ from the subscriber data of the network)")
		return
	}
	if amf[0]&amfSeparationBit == 0 {
		in.Errors = append(in.Errors, "authentication: the AMF separation bit of AUTN is 0 (not a 5G authentication)")
		return
	}
	u.resStar = security.ResStar(ck, ik, u.r.snn, rand, res)
	kseaf := security.Kseaf(security.Kausf(ck, ik, u.r.snn, sqnXorAK), u.r.snn)
	u.kamf = security.Kamf(kseaf, u.imsi, t.ABBA)
	u.ngKSI = t.NgKSI
	in.Notes = append(in.Notes, fmt.Sprintf("authentication: AUTN verified, RES* computed (SQN %x, ngKSI %d)", sqn, t.NgKSI&ksiMask))
}

// securityMode derives the NAS keys for the selected algorithms and checks the
// Security Mode Command with them (TS 33.501 §6.7.2).
func (u *UE) securityMode(t *nasMessage.SecurityModeCommand, wire []byte, in *models.Inbound) {
	if u.kamf == nil {
		in.Errors = append(in.Errors, "Security Mode Command: no K_AMF (it came before the authentication)")
		return
	}
	ia, ca := t.SelectedNASAlgorithms&nasAlgMask, t.SelectedNASAlgorithms>>cipheringAlgShift&nasAlgMask
	if ia > nassec.NIA2 || ca > nassec.NEA2 {
		in.Errors = append(in.Errors, fmt.Sprintf("Security Mode Command: NIA%d / NEA%d is not implemented", ia, ca))
		return
	}
	enc, integ := security.NASKeys(u.kamf, ca, ia)
	sec := &libnas.SecurityContext{IntegrityAlgorithm: ia, CipheringAlgorithm: ca, KnasInt: integ, KnasEnc: enc}
	if _, err := openDL(sec, wire); err != nil {
		in.Errors = append(in.Errors, "Security Mode Command: "+err.Error()+" (the key derivation does not match the network)")
		return
	}
	u.prevSec, u.sec, u.ngKSI = u.sec, sec, t.NgKSI
	u.imeisvRequested = t.HasIMEISVRequest && t.IMEISVRequest.Value&1 == 1
	in.Notes = append(in.Notes, fmt.Sprintf("Security Mode: NAS protection starts with NIA%d / NEA%d", ia, ca))
	if r := t.ReplayedUESecurityCapabilities; len(r) < 2 || r[0] != u.r.ea || r[1] != u.r.ia {
		in.Errors = append(in.Errors, fmt.Sprintf("Security Mode Command: the replayed UE security capabilities %x differ from the ones sent (%02x%02x)", r, u.r.ea, u.r.ia))
	}
}

// openDL checks and deciphers a protected downlink 5GMM message, like
// libnas.Decode does, but also for messages libnas cannot decode, so their
// type can still be named. The downlink COUNT advances only on success.
func openDL(sec *libnas.SecurityContext, b []byte) ([]byte, error) {
	hdr := b[1] & shtMask
	next := uint32(sec.DownlinkCount)
	count := next&^sqnMask | uint32(b[6])
	if count < next {
		count += sqnRange
	}
	bearer := sec.NASConnectionID
	if bearer == 0 {
		bearer = bearer3GPP
	}
	mac, err := nassec.Integrity(sec.IntegrityAlgorithm, sec.KnasInt, count, bearer, uint8(libnas.Downlink), b[6:])
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(mac[:], b[2:6]) {
		return nil, fmt.Errorf("the NAS MAC does not match (downlink COUNT %d)", count)
	}
	plain := b[protectedHeaderLen:]
	if hdr == shtCiphered || hdr == shtCipheredNewContext {
		if plain, err = nassec.Cipher(sec.CipheringAlgorithm, sec.KnasEnc, count, bearer, uint8(libnas.Downlink), plain); err != nil {
			return nil, err
		}
	}
	sec.DownlinkCount = libnas.Count(count + 1)
	return plain, nil
}

// receive5GSM handles the 5GSM message that a DL NAS Transport carries.
func (u *UE) receive5GSM(b []byte, in *models.Inbound) {
	m, err := libnas.Decode(b, nil)
	if err != nil {
		return // the name was already given; unsupported 5GSM bodies are only named
	}
	switch t := m.(type) {
	case *nasMessage.PDUSessionEstablishmentAccept:
		u.acceptSession(t, in)
	case *nasMessage.PDUSessionEstablishmentReject:
		if s := u.sessions[t.PDUSessionID]; s != nil && !s.ueIP.IsValid() {
			delete(u.sessions, t.PDUSessionID) // it never came up
		}
		in.Notes = append(in.Notes, fmt.Sprintf("PDU session %d establishment rejected", t.PDUSessionID))
	case *nasMessage.PDUSessionReleaseCommand:
		u.session(t.PDUSessionID).relPTI = t.ProcedureTransactionID
		in.Notes = append(in.Notes, fmt.Sprintf("PDU session %d release commanded (cause %d)", t.PDUSessionID, t.Cause))
	}
}
