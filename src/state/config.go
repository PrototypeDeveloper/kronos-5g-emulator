// Package state reads and checks the configuration (config.yaml).
package state

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/netip"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"

	"kronos/common/diag"
	"kronos/models"
)

// String returns "sst" or "sst/sd".
func (s SNSSAI) String() string {
	if s.SD == "" {
		return strconv.Itoa(s.SST)
	}
	return fmt.Sprintf("%d/%s", s.SST, strings.ToLower(s.SD))
}

// UnmarshalYAML parses the TAC; an invalid value is kept in err for validate to report.
func (t *TAC) UnmarshalYAML(n *yaml.Node) error {
	s, base := n.Value, 10
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		s, base = s[2:], 16
	}
	v, err := strconv.ParseUint(s, base, 32)
	if err != nil || v > maxTAC {
		t.err = fmt.Sprintf("gnodeb.tac %q must be an integer from 0 to 16777215 (3 octets); write hex with 0x", n.Value)
		return nil
	}
	t.Value, t.Set = uint32(v), true
	return nil
}

// ScenarioPath and SignalPath return the input files, resolved against the
// directory of config.yaml.
func (c *Config) ScenarioPath() string { return c.Base.Scenario }
func (c *Config) SignalPath() string   { return c.Base.Signal }

// SubscriberIMSI returns the IMSI of the n-th subscriber (0-based): the MSIN
// of ue.imsi plus n, keeping the MCC/MNC prefix.
func (c *Config) SubscriberIMSI(n int) (string, error) {
	imsi := c.UE.IMSI
	plen := mccDigits + len(c.GNodeB.PLMN.MNC)
	if plen != mccDigits+mncMinDigits && plen != mccDigits+mncMaxDigits {
		plen = mccDigits + mncMinDigits
	}
	if len(imsi) <= plen || !isDigits(imsi, 1, models.IMSIDigits) {
		return "", fmt.Errorf("invalid IMSI %q", imsi)
	}
	msin := imsi[plen:]
	v, err := strconv.ParseUint(msin, 10, 64)
	if err != nil {
		return "", err
	}
	next := fmt.Sprintf("%0*d", len(msin), v+uint64(n))
	if len(next) > len(msin) {
		return "", fmt.Errorf("the MSIN does not fit in %d digits (%s%s)", len(msin), imsi[:plen], next)
	}
	return imsi[:plen] + next, nil
}

// SubscriberIndex returns the 0-based index of imsi among the subscribers.
func (c *Config) SubscriberIndex(imsi string) (int, bool) {
	first, err := c.SubscriberIMSI(0)
	if err != nil || len(imsi) != len(first) || !isDigits(imsi, 1, models.IMSIDigits) {
		return 0, false
	}
	a, _ := strconv.ParseUint(first, 10, 64)
	b, _ := strconv.ParseUint(imsi, 10, 64)
	if b < a || b-a >= uint64(c.UE.Subscription) {
		return 0, false
	}
	return int(b - a), true
}

// GNBIDValue parses gnodeb.gnodebid.
func (g GNodeB) GNBIDValue() (uint32, error) {
	if !isHex(g.GNodeBID, 1, maxGNBIDHexDigits) {
		return 0, fmt.Errorf("gnodeb.gnodebid must be 1 to 8 hex digits (%q)", g.GNodeBID)
	}
	v, _ := strconv.ParseUint(g.GNodeBID, 16, 32)
	if g.GNodeBIDLength < maxGNBIDLength && v >= 1<<g.GNodeBIDLength {
		return 0, fmt.Errorf("gnodeb.gnodebid %s does not fit in %d bits", g.GNodeBID, g.GNodeBIDLength)
	}
	return uint32(v), nil
}

// NCIValue returns the 36-bit NR Cell Identity: gnodeb.nci, or by default the
// gNB ID in the upper bits and zeros below.
func (g GNodeB) NCIValue() (uint64, error) {
	if g.NCI == "" {
		id, err := g.GNBIDValue()
		if err != nil {
			return 0, err
		}
		return uint64(id) << (nciBits - g.GNodeBIDLength), nil
	}
	if !isHex(g.NCI, nciHexDigits, nciHexDigits) {
		return 0, fmt.Errorf("gnodeb.nci must be 9 hex digits, 36 bits (%q)", g.NCI)
	}
	v, _ := strconv.ParseUint(g.NCI, 16, 64)
	return v, nil
}

// Load reads config.yaml, applies defaults and the deprecated keys, and
// validates it. The config is nil only when the file cannot be
// read, is empty, or cannot be parsed as YAML at all.
func Load(path string) (*Config, diag.List) {
	var ds diag.List
	b, err := os.ReadFile(path)
	if err != nil {
		ds.Errorf(diag.Pos{File: path}, "cannot read the config file: %v", err)
		return nil, ds
	}
	var root yaml.Node
	if err := yaml.Unmarshal(b, &root); err != nil {
		ds = append(ds, yamlDiag(path, err.Error()))
		return nil, ds
	}
	if root.Kind == 0 {
		ds.Errorf(diag.Pos{File: path}, "the config file is empty")
		return nil, ds
	}
	cfg := &Config{Path: path}
	dec := yaml.NewDecoder(bytes.NewReader(b))
	dec.KnownFields(true)
	if err := dec.Decode(cfg); err != nil && !errors.Is(err, io.EOF) {
		var te *yaml.TypeError
		if !errors.As(err, &te) {
			ds = append(ds, yamlDiag(path, err.Error()))
			return nil, ds
		}
		for _, m := range te.Errors {
			ds = append(ds, yamlDiag(path, m))
		}
	}
	// Lines that YAML could not decode hold zero values; checking them again
	// would only repeat the error, so later diagnostics on those lines are dropped.
	broken := map[int]bool{}
	for _, d := range ds {
		broken[d.Pos.Line] = true
	}
	var more diag.List
	loc := newLocator(path, &root)
	cfg.normalize(&more, loc)
	cfg.validate(&more, loc)
	quotingHint(&more, path, &root)
	for _, d := range more {
		if !broken[d.Pos.Line] || d.Sev != diag.Error {
			ds = append(ds, d)
		}
	}
	return cfg, ds
}

// normalize resolves the input paths, fills in defaults and converts the deprecated keys.
func (c *Config) normalize(ds *diag.List, loc *locator) {
	// Input files are relative to the directory of config.yaml.
	dir := filepath.Dir(c.Path)
	resolve := func(p, def string) string {
		if p == "" {
			p = def
		}
		if filepath.IsAbs(p) {
			return p
		}
		return filepath.Join(dir, p)
	}
	c.Base.Scenario = resolve(c.Base.Scenario, "scenario.txt")
	c.Base.Signal = resolve(c.Base.Signal, "signal.txt")
	// Defaults for keys that are not set.
	if c.Base.Role == "" {
		c.Base.Role = "gnb"
	}
	if !loc.has("ue.subscription") {
		c.UE.Subscription = defaultSubscription
	}
	// The deprecated ue.type / ue.apn become session 1 unless ue.session is set.
	if c.UE.Type != "" || c.UE.APN != "" {
		where := loc.pos("ue.type")
		if !loc.has("ue.type") {
			where = loc.pos("ue.apn")
		}
		if len(c.UE.Session) > 0 {
			ds.Warnf(where, "ue.session is set, so the old ue.type / ue.apn are not used")
		} else {
			s := Session{ID: deprecatedSessionID, Type: c.UE.Type, DNN: c.UE.APN}
			if s.Type == "" {
				s.Type = "IPv4"
			}
			if len(c.GNodeB.SNSSAI) > 0 {
				sn := c.GNodeB.SNSSAI[0]
				s.SNSSAI = &sn
			}
			c.UE.Session = []Session{s}
			ds.Infof(where, "ue.type / ue.apn are the old form: read as ue.session (id 1, %s, dnn %s)", s.Type, s.DNN)
		}
	}
	// By default the UE supports NIA0-2 and NEA0-2.
	if c.UE.Integrity == nil {
		c.UE.Integrity = &Integrity{NIA0: true, NIA1: true, NIA2: true}
	}
	if c.UE.Ciphering == nil {
		c.UE.Ciphering = &Ciphering{NEA0: true, NEA1: true, NEA2: true}
	}
	if c.GNodeB.GNodeBIDLength == 0 {
		c.GNodeB.GNodeBIDLength = defaultGNBIDLength
	}
}

// validate checks every value and reports problems at the position of their key.
func (c *Config) validate(ds *diag.List, loc *locator) {
	at := loc.pos
	// Role: only gnb is implemented.
	switch strings.ToLower(c.Base.Role) {
	case "gnb":
	case "amf":
		ds.Errorf(at("base.role"), "the AMF role (reversed mode) is not implemented")
	default:
		ds.Errorf(at("base.role"), "base.role must be gnb or amf (%q)", c.Base.Role)
	}

	// PLMN digits.
	plmn := c.GNodeB.PLMN
	mccOK, mncOK := isDigits(plmn.MCC, mccDigits, mccDigits), isDigits(plmn.MNC, mncMinDigits, mncMaxDigits)
	if !mccOK {
		ds.Errorf(at("gnodeb.plmn.mcc"), "gnodeb.plmn.mcc must be 3 digits (%q)", plmn.MCC)
	}
	if !mncOK {
		ds.Errorf(at("gnodeb.plmn.mnc"), "gnodeb.plmn.mnc must be 2 or 3 digits (%q)", plmn.MNC)
	}

	// Subscribers: the IMSI, the range of consecutive IMSIs and the IMEISV.
	ue := c.UE
	imsiOK := isDigits(ue.IMSI, models.IMSIDigits, models.IMSIDigits)
	if !imsiOK {
		ds.Errorf(at("ue.imsi"), "ue.imsi must be 15 digits (%q)", ue.IMSI)
	}
	if imsiOK && mccOK && mncOK && !strings.HasPrefix(ue.IMSI, plmn.MCC+plmn.MNC) {
		ds.Warnf(at("ue.imsi"), "IMSI %s does not start with gnodeb.plmn (%s%s): a mistake unless roaming", ue.IMSI, plmn.MCC, plmn.MNC)
	}
	if ue.Subscription < minSubscription {
		ds.Errorf(at("ue.subscription"), "ue.subscription must be 1 or more (%d)", ue.Subscription)
	} else if imsiOK && mncOK {
		if _, err := c.SubscriberIMSI(ue.Subscription - 1); err != nil {
			ds.Errorf(at("ue.subscription"), "%d consecutive IMSIs overflow the digits: %v", ue.Subscription, err)
		}
	}
	if ue.IMEISV != "" && !isDigits(ue.IMEISV, imeisvDigits, imeisvDigits) {
		ds.Errorf(at("ue.imeisv"), "ue.imeisv must be 16 digits (%q)", ue.IMEISV)
	}

	// SUCI: only the null scheme is implemented.
	switch ue.SUCI.ProtectionScheme {
	case suciNullScheme:
	case suciProfileA, suciProfileB:
		ds.Errorf(at("ue.suci.protectionscheme"), "SUCI Profile A/B (protectionscheme 1/2) is not implemented: use 0 (the null scheme)")
	default:
		ds.Errorf(at("ue.suci.protectionscheme"), "ue.suci.protectionscheme must be 0 to 2 (%d)", ue.SUCI.ProtectionScheme)
	}
	if ue.SUCI.KeyID < 0 || ue.SUCI.KeyID > maxSUCIKeyID {
		ds.Errorf(at("ue.suci.keyid"), "ue.suci.keyid must be 0 to 255 (%d)", ue.SUCI.KeyID)
	}
	if ri := ue.SUCI.RoutingIndicator; ri != "" && !isDigits(ri, 1, maxRoutingIndicatorDigits) {
		ds.Errorf(at("ue.suci.routingindicator"), "ue.suci.routingindicator must be 1 to 4 digits (%q)", ri)
	}

	// Authentication credentials.
	cred := ue.Credential
	if !isHex(cred.Key, keyHexDigits, keyHexDigits) {
		ds.Errorf(at("ue.credential.key"), "ue.credential.key must be 32 hex digits (128 bits)")
	}
	if !isHex(cred.OP, keyHexDigits, keyHexDigits) {
		ds.Errorf(at("ue.credential.op"), "ue.credential.op must be 32 hex digits (128 bits)")
	}
	if t := strings.ToUpper(cred.Type); t != "OP" && t != "OPC" {
		ds.Errorf(at("ue.credential.type"), "ue.credential.type must be OP or OPC (%q)", cred.Type)
	}
	if !isHex(cred.AMF, amfHexDigits, amfHexDigits) {
		ds.Errorf(at("ue.credential.amf"), "ue.credential.amf must be 4 hex digits (%q)", cred.AMF)
	} else if v, _ := strconv.ParseUint(cred.AMF, 16, 16); v&amfSeparationBit == 0 {
		ds.Warnf(at("ue.credential.amf"), "the separation bit (the top bit) of ue.credential.amf is 0: a 5G UE checks it, so it is normally 1, as in 8000")
	}

	// S-NSSAIs of the UE and the gNB, then the PDU sessions.
	for i, s := range ue.NSSAI {
		checkSNSSAI(ds, at(fmt.Sprintf("ue.nssai[%d]", i)), "ue.nssai", s)
	}
	if len(c.GNodeB.SNSSAI) == 0 {
		ds.Errorf(at("gnodeb"), "gnodeb.snssai is missing (the S-NSSAIs announced in NG Setup)")
	}
	for i, s := range c.GNodeB.SNSSAI {
		checkSNSSAI(ds, at(fmt.Sprintf("gnodeb.snssai[%d]", i)), "gnodeb.snssai", s)
	}
	ids := map[int]bool{}
	for i, s := range ue.Session {
		p := at(fmt.Sprintf("ue.session[%d]", i))
		if s.ID < minPDUSessionID || s.ID > maxPDUSessionID {
			ds.Errorf(p, "the PDU session ID must be 1 to 15 (%d)", s.ID)
		} else if ids[s.ID] {
			ds.Errorf(p, "PDU session ID %d is used twice", s.ID)
		}
		ids[s.ID] = true
		switch strings.ToLower(s.Type) {
		case "ipv4", "ipv6", "ipv4v6":
		default:
			ds.Errorf(p, "the PDU session type must be IPv4, IPv6 or IPv4v6 (%q)", s.Type)
		}
		if s.DNN == "" {
			ds.Errorf(p, "PDU session %d has no dnn", s.ID)
		}
		if s.SNSSAI != nil {
			checkSNSSAI(ds, p, "ue.session.snssai", *s.SNSSAI)
			if !containsSNSSAI(c.GNodeB.SNSSAI, *s.SNSSAI) {
				ds.Warnf(p, "PDU session %d: S-NSSAI %s is not in gnodeb.snssai", s.ID, s.SNSSAI)
			}
		}
	}

	// Security algorithms: at least one of each; ZUC is not implemented.
	in, ci := ue.Integrity, ue.Ciphering
	if !(in.NIA0 || in.NIA1 || in.NIA2 || in.NIA3) {
		ds.Errorf(at("ue.integrity"), "ue.integrity is all false (no integrity algorithm)")
	}
	if !(ci.NEA0 || ci.NEA1 || ci.NEA2 || ci.NEA3) {
		ds.Errorf(at("ue.ciphering"), "ue.ciphering is all false (no ciphering algorithm)")
	}
	if in.NIA3 {
		ds.Warnf(at("ue.integrity.nia3"), "nia3 is true, but ZUC (128-NIA3) is not implemented: if the AMF picks NIA3, NAS fails. Keep it false for now")
	}
	if ci.NEA3 {
		ds.Warnf(at("ue.ciphering.nea3"), "nea3 is true, but ZUC (128-NEA3) is not implemented: if the AMF picks NEA3, NAS fails. Keep it false for now")
	}

	// gNB ID, NCI, RAN node name and TAC.
	g := c.GNodeB
	if g.GNodeBIDLength < minGNBIDLength || g.GNodeBIDLength > maxGNBIDLength {
		ds.Errorf(at("gnodeb.gnodebidlength"), "gnodeb.gnodebidlength must be 22 to 32 (%d)", g.GNodeBIDLength)
	} else if id, err := g.GNBIDValue(); err != nil {
		ds.Errorf(at("gnodeb.gnodebid"), "%v", err)
	} else if nci, err := g.NCIValue(); err != nil {
		ds.Errorf(at("gnodeb.nci"), "%v", err)
	} else if top := nci >> (nciBits - g.GNodeBIDLength); top != uint64(id) {
		ds.Errorf(at("gnodeb.nci"), "the top %d bits of gnodeb.nci (%x) do not match gnodeb.gnodebid (%x)", g.GNodeBIDLength, top, id)
	}
	if len(g.Name) > maxRANNodeNameLength {
		ds.Errorf(at("gnodeb.name"), "gnodeb.name is at most 150 characters")
	}
	for _, r := range g.Name {
		if r < minPrintableASCII || r > maxPrintableASCII {
			ds.Errorf(at("gnodeb.name"), "gnodeb.name must be printable ASCII (RANNodeName is a PrintableString)")
			break
		}
	}
	switch {
	case g.TAC.err != "":
		ds.Errorf(at("gnodeb.tac"), "%s", g.TAC.err)
	case !g.TAC.Set:
		ds.Errorf(at("gnodeb"), "gnodeb.tac is missing")
	}

	// Endpoints: the AMF address (ngap.dst) and the local GTP-U address are required.
	checkEndpoint(ds, at("gnodeb.ngap.src"), "gnodeb.ngap.src", g.NGAP.Src, false, false)
	checkEndpoint(ds, at("gnodeb.ngap.dst"), "gnodeb.ngap.dst", g.NGAP.Dst, true, true)
	checkEndpoint(ds, at("gnodeb.gtpu.src"), "gnodeb.gtpu.src", g.GTPU.Src, true, true)
	checkEndpoint(ds, at("gnodeb.gtpu.dst"), "gnodeb.gtpu.dst", g.GTPU.Dst, false, true)
	if g.Web != nil {
		checkEndpoint(ds, at("gnodeb.web"), "gnodeb.web.src", g.Web.Src, true, true)
	}
	if g.HTTP != nil {
		checkEndpoint(ds, at("gnodeb.http"), "gnodeb.http.src", g.HTTP.Src, true, true)
	}
}

// SyncListen returns gnodeb.http.src, where the pseudo-calls talk to each
// other (sync_send / sync_receive), or nil.
func (c *Config) SyncListen() *Endpoint {
	if c.GNodeB.HTTP == nil {
		return nil
	}
	return c.GNodeB.HTTP.Src
}

// WebListen returns where the web console listens, or nil when gnodeb.web is not set.
func (c *Config) WebListen() *Endpoint {
	if c.GNodeB.Web == nil {
		return nil
	}
	return c.GNodeB.Web.Src
}

// checkSNSSAI checks the sst range and the sd format of s.
func checkSNSSAI(ds *diag.List, p diag.Pos, what string, s SNSSAI) {
	if s.SST < 0 || s.SST > maxSST {
		ds.Errorf(p, "%s: sst must be 0 to 255 (%d)", what, s.SST)
	}
	if s.SD != "" && !isHex(s.SD, sdHexDigits, sdHexDigits) {
		ds.Errorf(p, "%s: sd must be 6 hex digits (%q)", what, s.SD)
	}
}

// containsSNSSAI reports whether list holds s, comparing sd case-insensitively.
func containsSNSSAI(list []SNSSAI, s SNSSAI) bool {
	for _, x := range list {
		if x.SST == s.SST && strings.EqualFold(x.SD, s.SD) {
			return true
		}
	}
	return false
}

// checkEndpoint checks e's address and port; required and portRequired reject a missing endpoint or port 0.
func checkEndpoint(ds *diag.List, p diag.Pos, what string, e *Endpoint, required, portRequired bool) {
	if e == nil {
		if required {
			ds.Errorf(p, "%s (address and port) is missing", what)
		}
		return
	}
	if _, err := netip.ParseAddr(e.Address); err != nil {
		ds.Errorf(p, "%s.address %q is not an IP address", what, e.Address)
	}
	if e.Port < 0 || e.Port > maxPort || (portRequired && e.Port == 0) {
		ds.Errorf(p, "%s.port must be 1 to 65535 (%d)", what, e.Port)
	}
}

// isDigits reports whether s is min to max decimal digits.
func isDigits(s string, min, max int) bool {
	if len(s) < min || len(s) > max {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// isHex reports whether s is min to max hexadecimal digits.
func isHex(s string, min, max int) bool {
	if len(s) < min || len(s) > max {
		return false
	}
	for _, r := range s {
		if !(r >= '0' && r <= '9' || r >= 'a' && r <= 'f' || r >= 'A' && r <= 'F') {
			return false
		}
	}
	return true
}

// newLocator records the position of every key in the parsed YAML document.
func newLocator(file string, root *yaml.Node) *locator {
	l := &locator{file: file, at: map[string]diag.Pos{}}
	if root.Kind == yaml.DocumentNode && len(root.Content) > 0 {
		l.walk(root.Content[0], "")
	}
	return l
}

// walk records the positions of the keys and sequence items under n.
func (l *locator) walk(n *yaml.Node, path string) {
	switch n.Kind {
	case yaml.MappingNode:
		for i := 0; i+1 < len(n.Content); i += 2 {
			k := n.Content[i]
			p := k.Value
			if path != "" {
				p = path + "." + k.Value
			}
			l.at[p] = diag.Pos{File: l.file, Line: k.Line, Col: k.Column}
			l.walk(n.Content[i+1], p)
		}
	case yaml.SequenceNode:
		for i, c := range n.Content {
			p := fmt.Sprintf("%s[%d]", path, i)
			l.at[p] = diag.Pos{File: l.file, Line: c.Line, Col: c.Column}
			l.walk(c, p)
		}
	}
}

// has reports whether path appears in the file.
func (l *locator) has(path string) bool {
	_, ok := l.at[path]
	return ok
}

// pos returns the position of path, or of its nearest ancestor that exists.
func (l *locator) pos(path string) diag.Pos {
	for p := path; p != ""; {
		if pos, ok := l.at[p]; ok {
			return pos
		}
		i := strings.LastIndexAny(p, ".[")
		if i < 0 {
			break
		}
		p = p[:i]
	}
	return diag.Pos{File: l.file}
}

// stringKeys are the keys that hold digit strings.
var stringKeys = map[string]bool{
	"ue.imsi": true, "ue.imeisv": true, "ue.suci.key": true, "ue.suci.routingindicator": true,
	"ue.credential.key": true, "ue.credential.op": true, "ue.credential.amf": true,
	"ue.nssai[].sd": true, "ue.session[].snssai.sd": true, "gnodeb.gnodebid": true, "gnodeb.nci": true,
	"gnodeb.plmn.mcc": true, "gnodeb.plmn.mnc": true, "gnodeb.snssai[].sd": true,
}

var indexRE = regexp.MustCompile(`\[\d+\]`)

// quotingHint notes digit strings that YAML resolves as numbers. kronos reads
// them as strings, but other YAML tools turn `sd: 010203` into 4227.
func quotingHint(ds *diag.List, file string, root *yaml.Node) {
	var hits []string
	first := diag.Pos{}
	var walk func(n *yaml.Node, path string)
	walk = func(n *yaml.Node, path string) {
		switch n.Kind {
		case yaml.DocumentNode:
			for _, c := range n.Content {
				walk(c, path)
			}
		case yaml.MappingNode:
			for i := 0; i+1 < len(n.Content); i += 2 {
				k, v := n.Content[i], n.Content[i+1]
				p := k.Value
				if path != "" {
					p = path + "." + k.Value
				}
				if v.Kind == yaml.ScalarNode && v.Style == 0 && v.Tag != "!!str" && stringKeys[indexRE.ReplaceAllString(p, "[]")] {
					if len(hits) == 0 {
						first = diag.Pos{File: file, Line: v.Line, Col: v.Column}
					}
					hits = append(hits, fmt.Sprintf("%s (line %d)", k.Value, v.Line))
				}
				walk(v, p)
			}
		case yaml.SequenceNode:
			for i, c := range n.Content {
				walk(c, fmt.Sprintf("%s[%d]", path, i))
			}
		}
	}
	walk(root, "")
	if len(hits) > 0 {
		ds.Infof(first, "values that YAML could read as numbers are not quoted: %s. kronos reads them as strings, so nothing changes here, but other YAML tools change them (sd: 010203 becomes 4227, for example): quote them as \"…\"", strings.Join(hits, ", "))
	}
}

var (
	lineRE      = regexp.MustCompile(`^(?:yaml: )?line (\d+): (.*)$`)
	notFoundRE  = regexp.MustCompile(`^field (\S+) not found in type (\S+)$`)
	unmarshalRE = regexp.MustCompile("^cannot unmarshal !!\\w+ `(.*)` into (\\S+)$")
)

// yamlDiag turns a yaml.v3 error message into a positioned diagnostic.
func yamlDiag(file, msg string) diag.Diag {
	pos := diag.Pos{File: file}
	if m := lineRE.FindStringSubmatch(msg); m != nil {
		pos.Line, _ = strconv.Atoi(m[1])
		msg = m[2]
	}
	if m := notFoundRE.FindStringSubmatch(msg); m != nil {
		msg = fmt.Sprintf("unknown key %s", m[1])
	} else if m := unmarshalRE.FindStringSubmatch(msg); m != nil {
		if what, ok := map[string]string{"int": "an integer", "bool": "true or false", "string": "a string", "float64": "a number"}[m[2]]; ok {
			msg = fmt.Sprintf("cannot read %s as %s", m[1], what)
		} else {
			msg = fmt.Sprintf("cannot read %s as type %s", m[1], m[2])
		}
	} else {
		msg = "cannot read as YAML: " + strings.TrimPrefix(msg, "yaml: ")
	}
	return diag.Diag{Sev: diag.Error, Pos: pos, Msg: msg}
}
