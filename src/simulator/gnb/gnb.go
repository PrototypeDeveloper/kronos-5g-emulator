// Package gnb is the gNB role: the RAN side of NGAP towards
// a real AMF and the per-subscriber NAS mapping. It implements role.Role and
// never imports the engine.
package gnb

import (
	"encoding/hex"
	"fmt"
	"net/netip"
	"strings"

	"libngap/ngapMessage"
	"libngap/ngapType"

	"kronos/common/plmn"
	"kronos/common/security"
	"kronos/models"
	"kronos/simulator/role"
	"kronos/state"
)

var _ role.Role = (*Role)(nil)

// New derives the values the gNB writes into NGAP from a validated config.
func New(cfg *state.Config) (*Role, error) {
	g := cfg.GNodeB
	p, ok := plmn.Encode(g.PLMN.MCC, g.PLMN.MNC)
	if !ok {
		return nil, fmt.Errorf("invalid gnodeb.plmn %s/%s", g.PLMN.MCC, g.PLMN.MNC)
	}
	id, err := g.GNBIDValue()
	if err != nil {
		return nil, err
	}
	t := g.TAC.Value
	r := &Role{name: g.Name, plmn: p, gnbID: BitString(uint64(id), g.GNodeBIDLength), tac: []byte{byte(t >> 16), byte(t >> 8), byte(t)}}
	for _, s := range g.SNSSAI {
		item, err := sliceItem(s)
		if err != nil {
			return nil, err
		}
		r.slices = append(r.slices, item)
	}
	if err := r.loadSubscriber(cfg); err != nil {
		return nil, err
	}
	return r, nil
}

// loadSubscriber derives what the NAS mapping needs from the ue section.
func (r *Role) loadSubscriber(cfg *state.Config) error {
	g, u := cfg.GNodeB, cfg.UE
	r.mcc, r.mnc = g.PLMN.MCC, g.PLMN.MNC
	r.snn = security.ServingNetworkName(r.mcc, r.mnc)
	nci, err := g.NCIValue()
	if err != nil {
		return err
	}
	r.nci = BitString(nci, nciBits)
	// Credentials: K and OPc (derived from OP when the type is OP).
	if r.k, err = hex.DecodeString(u.Credential.Key); err != nil {
		return fmt.Errorf("ue.credential.key: %v", err)
	}
	op, err := hex.DecodeString(u.Credential.OP)
	if err != nil {
		return fmt.Errorf("ue.credential.op: %v", err)
	}
	r.opc = op
	if strings.EqualFold(u.Credential.Type, "OP") {
		if r.opc, err = security.OPc(r.k, op); err != nil {
			return err
		}
	}
	r.ri = u.SUCI.RoutingIndicator
	if u.IMEISV != "" {
		if r.imeisv, err = security.IMEISV(u.IMEISV); err != nil {
			return err
		}
	}
	// UE security capability: one bit per algorithm, NEA0/NIA0 first.
	for i, on := range []bool{u.Ciphering.NEA0, u.Ciphering.NEA1, u.Ciphering.NEA2, u.Ciphering.NEA3} {
		if on {
			r.ea |= secCapAlg0Bit >> i
		}
	}
	for i, on := range []bool{u.Integrity.NIA0, u.Integrity.NIA1, u.Integrity.NIA2, u.Integrity.NIA3} {
		if on {
			r.ia |= secCapAlg0Bit >> i
		}
	}
	// Requested NSSAI: S-NSSAI IEs of length 1 (SST) or 4 (SST and SD).
	for _, s := range u.NSSAI {
		if s.SD == "" {
			r.nssai = append(r.nssai, 1, byte(s.SST))
			continue
		}
		sd, err := hex.DecodeString(s.SD)
		if err != nil || len(sd) != sdLen {
			return fmt.Errorf("invalid sd %q in ue.nssai", s.SD)
		}
		r.nssai = append(append(r.nssai, 4, byte(s.SST)), sd...)
	}
	if r.sessions, err = loadSessions(u); err != nil {
		return err
	}
	// N3: the gNB's GTP-U address and an optional fixed UPF.
	if s := g.GTPU.Src; s != nil {
		if r.n3, err = netip.ParseAddr(s.Address); err != nil {
			return fmt.Errorf("gnodeb.gtpu.src: %v", err)
		}
	}
	if d := g.GTPU.Dst; d != nil {
		a, err := netip.ParseAddr(d.Address)
		if err != nil {
			return fmt.Errorf("gnodeb.gtpu.dst: %v", err)
		}
		r.upfOverride = netip.AddrPortFrom(a, uint16(d.Port))
	}
	return nil
}

// NewUE implements role.Role.
func (r *Role) NewUE(index int, imsi string) role.UE {
	return &UE{r: r, index: index, imsi: imsi, msin: imsi[len(r.mcc)+len(r.mnc):], ngKSI: noKey}
}

// Name implements role.Role.
func (r *Role) Name() string { return "gnb" }

// Listen implements role.Role: the gNB dials the AMF.
func (r *Role) Listen() bool { return false }

// PatchNode implements role.Role. NGSetupRequest gets GlobalRANNodeID,
// RANNodeName and SupportedTAList from config.
func (r *Role) PatchNode(msg any) error {
	if m, ok := msg.(*ngapMessage.NGSetupRequest); ok {
		r.patchNGSetupRequest(m)
	}
	return nil
}

// patchNGSetupRequest writes the node identity, name and TA of the config.
func (r *Role) patchNGSetupRequest(m *ngapMessage.NGSetupRequest) {
	id := &m.GlobalRANNodeID
	id.Present = ngapType.GlobalRANNodeIDPresentGlobalGNBID
	id.GlobalGNBID.PLMNIdentity = ngapType.PLMNIdentity{Value: clone(r.plmn)}
	id.GlobalGNBID.GNBID.Present = ngapType.GNBIDPresentGNBID
	id.GlobalGNBID.GNBID.GNBID = ngapType.BitString{Bytes: clone(r.gnbID.Bytes), BitLength: r.gnbID.BitLength}
	if r.name != "" {
		m.HasRANNodeName = true
		m.RANNodeName.Value = r.name
	}
	// config describes one TA: keep the template's first entries (and their
	// extensions) as the pattern and replace the values.
	var ta ngapType.SupportedTAItem
	if len(m.SupportedTAList.List) > 0 {
		ta = m.SupportedTAList.List[0]
	}
	ta.TAC = ngapType.TAC{Value: clone(r.tac)}
	var bp ngapType.BroadcastPLMNItem
	if len(ta.BroadcastPLMNList.List) > 0 {
		bp = ta.BroadcastPLMNList.List[0]
	}
	bp.PLMNIdentity = ngapType.PLMNIdentity{Value: clone(r.plmn)}
	bp.TAISliceSupportList = ngapType.SliceSupportList{List: append([]ngapType.SliceSupportItem(nil), r.slices...)}
	ta.BroadcastPLMNList.List = []ngapType.BroadcastPLMNItem{bp}
	m.SupportedTAList.List = []ngapType.SupportedTAItem{ta}
}

// NodeReceived implements role.Role: record what NG Setup returned.
func (r *Role) NodeReceived(msg any, st *models.NodeStatus) string {
	switch m := msg.(type) {
	case *ngapMessage.NGSetupResponse:
		st.Setup, st.Cause = models.SetupOK, ""
		st.PeerName = m.AMFName.Value
		st.PeerInfo = nil
		for _, g := range m.ServedGUAMIList.List {
			amfID := BitValue(g.GUAMI.AMFRegionID.Value)<<(amfSetIDBits+amfPointerBits) | BitValue(g.GUAMI.AMFSetID.Value)<<amfPointerBits | BitValue(g.GUAMI.AMFPointer.Value)
			st.PeerInfo = append(st.PeerInfo, fmt.Sprintf("GUAMI %s-%06x", plmn.String(g.GUAMI.PLMNIdentity.Value), amfID))
		}
		for _, p := range m.PLMNSupportList.List {
			var slices []string
			for _, s := range p.SliceSupportList.List {
				slices = append(slices, snssaiString(s.SNSSAI))
			}
			st.PeerInfo = append(st.PeerInfo, fmt.Sprintf("PLMN %s (slices %s)", plmn.String(p.PLMNIdentity.Value), strings.Join(slices, ", ")))
		}
		return fmt.Sprintf("NG Setup succeeded: AMF %s, %s", st.PeerName, strings.Join(st.PeerInfo, ", "))
	case *ngapMessage.NGSetupFailure:
		st.Setup = models.SetupFailed
		st.Cause = CauseString(m.Cause)
		if m.HasTimeToWait {
			st.Cause += ", TimeToWait " + timeToWait(m.TimeToWait.Value)
		}
		return "NG Setup failed: " + st.Cause
	}
	return ""
}

// sliceItem converts a configured S-NSSAI to an NGAP slice support item.
func sliceItem(s state.SNSSAI) (ngapType.SliceSupportItem, error) {
	var it ngapType.SliceSupportItem
	it.SNSSAI.SST = ngapType.SST{Value: []byte{byte(s.SST)}}
	if s.SD != "" {
		sd, err := hex.DecodeString(s.SD)
		if err != nil || len(sd) != sdLen {
			return it, fmt.Errorf("invalid sd %q in S-NSSAI", s.SD)
		}
		it.SNSSAI.HasSD = true
		it.SNSSAI.SD = ngapType.SD{Value: sd}
	}
	return it, nil
}

// snssaiString renders an S-NSSAI as "SST" or "SST/SD".
func snssaiString(s ngapType.SNSSAI) string {
	out := "?"
	if len(s.SST.Value) == 1 {
		out = fmt.Sprint(s.SST.Value[0])
	}
	if s.HasSD {
		out += "/" + hex.EncodeToString(s.SD.Value)
	}
	return out
}

// BitString left-aligns the low n bits of v, as NGAP BIT STRINGs are encoded.
func BitString(v uint64, n int) ngapType.BitString {
	nb := (n + 7) / 8
	v <<= uint(nb*8 - n)
	b := make([]byte, nb)
	for i := nb - 1; i >= 0; i-- {
		b[i] = byte(v)
		v >>= 8
	}
	return ngapType.BitString{Bytes: b, BitLength: uint64(n)}
}

// BitValue is the inverse of BitString.
func BitValue(bs ngapType.BitString) uint64 {
	var v uint64
	for _, b := range bs.Bytes {
		v = v<<8 | uint64(b)
	}
	return v >> (uint64(len(bs.Bytes))*8 - bs.BitLength)
}

// clone returns a copy of b.
func clone(b []byte) []byte { return append([]byte(nil), b...) }

// Names of NGAP Cause values, indexed by the enumeration.
var (
	causeTransport = []string{"transport-resource-unavailable", "unspecified"}
	causeProtocol  = []string{"transfer-syntax-error", "abstract-syntax-error-reject", "abstract-syntax-error-ignore-and-notify",
		"message-not-compatible-with-receiver-state", "semantic-error", "abstract-syntax-error-falsely-constructed-message", "unspecified"}
	causeMisc = []string{"control-processing-overload", "not-enough-user-plane-processing-resources", "hardware-failure",
		"om-intervention", "unknown-PLMN-or-SNPN", "unspecified"}
)

// CauseString renders an NGAP Cause such as "misc(4) unknown-PLMN-or-SNPN".
func CauseString(c ngapType.Cause) string {
	named := func(group string, v ngapType.Enumerated, names []string) string {
		s := fmt.Sprintf("%s(%d)", group, v)
		if int(v) < len(names) {
			s += " " + names[v]
		}
		return s
	}
	switch c.Present {
	case ngapType.CausePresentRadioNetwork:
		return named("radioNetwork", c.RadioNetwork.Value, nil)
	case ngapType.CausePresentTransport:
		return named("transport", c.Transport.Value, causeTransport)
	case ngapType.CausePresentNas:
		return named("nas", c.Nas.Value, nil)
	case ngapType.CausePresentProtocol:
		return named("protocol", c.Protocol.Value, causeProtocol)
	case ngapType.CausePresentMisc:
		return named("misc", c.Misc.Value, causeMisc)
	}
	return "unknown cause"
}

// timeToWait renders an NGAP TimeToWait value.
func timeToWait(v ngapType.Enumerated) string {
	if names := []string{"1s", "2s", "5s", "10s", "20s", "60s"}; int(v) < len(names) {
		return names[v]
	}
	return fmt.Sprintf("(%d)", v)
}
