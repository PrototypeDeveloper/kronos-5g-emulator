// Structs of package state.

package state

import "kronos/common/diag"

// The configuration file and its sections.

// Config is config.yaml. Digit strings such as the IMSI, MNC
// and SD are strings: YAML would otherwise read `sd: 010203` as octal 4227.
type Config struct {
	Base   Base   `yaml:"base"`
	UE     UE     `yaml:"ue"`
	GNodeB GNodeB `yaml:"gnodeb"`

	Path string `yaml:"-"` // the file the config was read from
}

// Base is the base: section: the role, the snap and debug switches and the input files.
type Base struct {
	Role     string `yaml:"role"`     // gnb (default) | amf (reversed mode, not implemented yet)
	Snap     bool   `yaml:"snap"`     // snap() runs only when true
	Debug    bool   `yaml:"debug"`    // debug log level
	Scenario string `yaml:"scenario"` // relative to the config file; default scenario.txt
	Signal   string `yaml:"signal"`   // relative to the config file; default signal.txt
}

// The ue: section.

// UE is the ue: section: the emulated subscribers, their credentials and PDU sessions.
type UE struct {
	IMSI         string     `yaml:"imsi"`         // IMSI of the first subscriber
	Subscription int        `yaml:"subscription"` // number of subscribers, IMSIs counted up from IMSI
	IMEISV       string     `yaml:"imeisv"`
	SUCI         SUCI       `yaml:"suci"`
	Credential   Credential `yaml:"credential"`
	NSSAI        []SNSSAI   `yaml:"nssai"`
	Session      []Session  `yaml:"session"`
	Integrity    *Integrity `yaml:"integrity"`
	Ciphering    *Ciphering `yaml:"ciphering"`
	// TUN gives each UE with a PDU session a TUN device on the host
	// (kronos-ue1 …) holding the UE address, so host programs (iperf3 -B)
	// can use the session with the kernel's TCP/IP. Needs CAP_NET_ADMIN.
	TUN bool `yaml:"tun"`

	// Deprecated: ue.type and ue.apn are read as session[0].
	Type string `yaml:"type"`
	APN  string `yaml:"apn"`
}

// SUCI selects how the SUPI is concealed in the SUCI.
type SUCI struct {
	ProtectionScheme int    `yaml:"protectionscheme"` // 0 null, 1 Profile A, 2 Profile B
	KeyID            int    `yaml:"keyid"`
	Key              string `yaml:"key"` // home network public key (Profile A/B)
	RoutingIndicator string `yaml:"routingindicator"`
}

// Credential is the subscriber key K, OP or OPc (per Type) and the AMF field for MILENAGE.
type Credential struct {
	Key  string `yaml:"key"`
	OP   string `yaml:"op"`
	Type string `yaml:"type"` // OP | OPC
	AMF  string `yaml:"amf"`
}

// SNSSAI is an S-NSSAI: the slice/service type and an optional 6-hex-digit SD.
type SNSSAI struct {
	SST int    `yaml:"sst"`
	SD  string `yaml:"sd"`
}

// Session is one PDU session of the UE.
type Session struct {
	ID     int     `yaml:"id"`
	Type   string  `yaml:"type"` // IPv4 | IPv6 | IPv4v6
	DNN    string  `yaml:"dnn"`
	SNSSAI *SNSSAI `yaml:"snssai"`
}

// Integrity selects the 5G integrity algorithms the UE announces as supported.
type Integrity struct {
	NIA0 bool `yaml:"nia0"`
	NIA1 bool `yaml:"nia1"`
	NIA2 bool `yaml:"nia2"`
	NIA3 bool `yaml:"nia3"`
}

// Ciphering selects the 5G ciphering algorithms the UE announces as supported.
type Ciphering struct {
	NEA0 bool `yaml:"nea0"`
	NEA1 bool `yaml:"nea1"`
	NEA2 bool `yaml:"nea2"`
	NEA3 bool `yaml:"nea3"`
}

// The gnodeb: section.

// GNodeB is the gnodeb: section: the gNB identity, cell and endpoints.
type GNodeB struct {
	Name           string   `yaml:"name"`           // RANNodeName; empty keeps the template value
	GNodeBID       string   `yaml:"gnodebid"`       // hex
	GNodeBIDLength int      `yaml:"gnodebidlength"` // 22..32 bits; default 32
	NCI            string   `yaml:"nci"`            // 36-bit NR Cell Identity, 9 hex digits; default gNB ID followed by zero bits
	NGAP           Link     `yaml:"ngap"`
	GTPU           Link     `yaml:"gtpu"`
	PLMN           PLMN     `yaml:"plmn"`
	SNSSAI         []SNSSAI `yaml:"snssai"`
	TAC            TAC      `yaml:"tac"`

	// HTTP is where pseudo-calls talk to each other (sync_send /
	// sync_receive: JSON over HTTP); none disables it.
	HTTP *Listen `yaml:"http"`
	// Web is where the web console listens; none disables it.
	Web *Listen `yaml:"web"`
}

// Link is the local (src) and peer (dst) endpoint of an interface.
type Link struct {
	Src *Endpoint `yaml:"src"`
	Dst *Endpoint `yaml:"dst"`
}

// Endpoint is an IP address and port.
type Endpoint struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

// Listen is a listening endpoint written as src: {address, port}.
type Listen struct {
	Src *Endpoint `yaml:"src"`
}

// PLMN is the MCC and MNC as digit strings.
type PLMN struct {
	MCC string `yaml:"mcc"`
	MNC string `yaml:"mnc"`
}

// TAC accepts decimal (leading zeros allowed) or 0x-prefixed hex. A plain
// int field would read `tac: 000010` as octal.
type TAC struct {
	Value uint32
	Set   bool
	err   string
}

// Key positions in the parsed file.

// locator maps key paths such as "ue.credential.amf" to their position.
type locator struct {
	file string
	at   map[string]diag.Pos
}
