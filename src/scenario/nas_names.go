package scenario

import (
	"kronos/models"
	"sort"
	"strings"
)

// nasTable lists the known 5GS NAS messages.
var nasTable = []models.NASInfo{
	// 5GMM
	{Name: "RegistrationRequest", EPD: models.EPD5GMM, Type: 0x41, Supported: true},
	{Name: "RegistrationAccept", EPD: models.EPD5GMM, Type: 0x42, Supported: true},
	{Name: "RegistrationComplete", EPD: models.EPD5GMM, Type: 0x43, Supported: true},
	{Name: "RegistrationReject", EPD: models.EPD5GMM, Type: 0x44, Supported: true},
	{Name: "DeregistrationRequestUE", EPD: models.EPD5GMM, Type: 0x45, Supported: true}, // UE originating
	{Name: "DeregistrationAcceptUE", EPD: models.EPD5GMM, Type: 0x46, Supported: true},
	{Name: "DeregistrationRequestNW", EPD: models.EPD5GMM, Type: 0x47, Supported: true}, // UE terminated (network initiated)
	{Name: "DeregistrationAcceptNW", EPD: models.EPD5GMM, Type: 0x48, Supported: true},
	{Name: "ServiceRequest", EPD: models.EPD5GMM, Type: 0x4c, Supported: true},
	{Name: "ServiceReject", EPD: models.EPD5GMM, Type: 0x4d, Supported: true},
	{Name: "ServiceAccept", EPD: models.EPD5GMM, Type: 0x4e, Supported: true},
	{Name: "ControlPlaneServiceRequest", EPD: models.EPD5GMM, Type: 0x4f, Supported: true},
	{Name: "NetworkSliceSpecificAuthenticationCommand", EPD: models.EPD5GMM, Type: 0x50, Supported: true},
	{Name: "NetworkSliceSpecificAuthenticationComplete", EPD: models.EPD5GMM, Type: 0x51, Supported: true},
	{Name: "NetworkSliceSpecificAuthenticationResult", EPD: models.EPD5GMM, Type: 0x52, Supported: true},
	{Name: "ConfigurationUpdateCommand", EPD: models.EPD5GMM, Type: 0x54, Supported: true},
	{Name: "ConfigurationUpdateComplete", EPD: models.EPD5GMM, Type: 0x55, Supported: true},
	{Name: "AuthenticationRequest", EPD: models.EPD5GMM, Type: 0x56, Supported: true},
	{Name: "AuthenticationResponse", EPD: models.EPD5GMM, Type: 0x57, Supported: true},
	{Name: "AuthenticationReject", EPD: models.EPD5GMM, Type: 0x58, Supported: true},
	{Name: "AuthenticationFailure", EPD: models.EPD5GMM, Type: 0x59, Supported: true},
	{Name: "AuthenticationResult", EPD: models.EPD5GMM, Type: 0x5a, Supported: true},
	{Name: "IdentityRequest", EPD: models.EPD5GMM, Type: 0x5b, Supported: true},
	{Name: "IdentityResponse", EPD: models.EPD5GMM, Type: 0x5c, Supported: true},
	{Name: "SecurityModeCommand", EPD: models.EPD5GMM, Type: 0x5d, Supported: true},
	{Name: "SecurityModeComplete", EPD: models.EPD5GMM, Type: 0x5e, Supported: true},
	{Name: "SecurityModeReject", EPD: models.EPD5GMM, Type: 0x5f, Supported: true},
	{Name: "Status5GMM", EPD: models.EPD5GMM, Type: 0x64, Supported: true},
	{Name: "Notification", EPD: models.EPD5GMM, Type: 0x65, Supported: true},
	{Name: "NotificationResponse", EPD: models.EPD5GMM, Type: 0x66, Supported: true},
	{Name: "ULNASTransport", EPD: models.EPD5GMM, Type: 0x67, Supported: true},
	{Name: "DLNASTransport", EPD: models.EPD5GMM, Type: 0x68, Supported: true},
	// 5GSM
	{Name: "PDUSessionEstablishmentRequest", EPD: models.EPD5GSM, Type: 0xc1, Supported: true},
	{Name: "PDUSessionEstablishmentAccept", EPD: models.EPD5GSM, Type: 0xc2, Supported: true},
	{Name: "PDUSessionEstablishmentReject", EPD: models.EPD5GSM, Type: 0xc3, Supported: true},
	{Name: "PDUSessionAuthenticationCommand", EPD: models.EPD5GSM, Type: 0xc5, Supported: true},
	{Name: "PDUSessionAuthenticationComplete", EPD: models.EPD5GSM, Type: 0xc6, Supported: true},
	{Name: "PDUSessionAuthenticationResult", EPD: models.EPD5GSM, Type: 0xc7, Supported: true},
	{Name: "PDUSessionModificationRequest", EPD: models.EPD5GSM, Type: 0xc9, Supported: true},
	{Name: "PDUSessionModificationReject", EPD: models.EPD5GSM, Type: 0xca, Supported: true},
	{Name: "PDUSessionModificationCommand", EPD: models.EPD5GSM, Type: 0xcb, Supported: true},
	{Name: "PDUSessionModificationComplete", EPD: models.EPD5GSM, Type: 0xcc, Supported: true},
	{Name: "PDUSessionModificationCommandReject", EPD: models.EPD5GSM, Type: 0xcd, Supported: true},
	{Name: "PDUSessionReleaseRequest", EPD: models.EPD5GSM, Type: 0xd1, Supported: true},
	{Name: "PDUSessionReleaseReject", EPD: models.EPD5GSM, Type: 0xd2, Supported: true},
	{Name: "PDUSessionReleaseCommand", EPD: models.EPD5GSM, Type: 0xd3, Supported: true},
	{Name: "PDUSessionReleaseComplete", EPD: models.EPD5GSM, Type: 0xd4, Supported: true},
	{Name: "Status5GSM", EPD: models.EPD5GSM, Type: 0xd6, Supported: true},
}

// nasAliases maps alternative spellings (lower case) to the table names.
var nasAliases = map[string]string{
	"5gmmstatus":                         "Status5GMM",
	"5gsmstatus":                         "Status5GSM",
	"deregistrationrequestueoriginating": "DeregistrationRequestUE",
	"deregistrationacceptueoriginating":  "DeregistrationAcceptUE",
	"deregistrationrequestueterminated":  "DeregistrationRequestNW",
	"deregistrationacceptueterminated":   "DeregistrationAcceptNW",
	"uplinknastransport":                 "ULNASTransport",
	"downlinknastransport":               "DLNASTransport",
}

var (
	nasByLower = map[string]*models.NASInfo{}
	nasByType  = map[[2]uint8]*models.NASInfo{}
)

// init indexes nasTable by name, alias and (EPD, message type).
func init() {
	for i := range nasTable {
		n := &nasTable[i]
		nasByLower[strings.ToLower(n.Name)] = n
		nasByType[[2]uint8{n.EPD, n.Type}] = n
	}
	for alias, name := range nasAliases {
		nasByLower[alias] = nasByLower[strings.ToLower(name)]
	}
}

// LookupNAS finds a NAS message by name, ignoring case.
func LookupNAS(name string) (*models.NASInfo, bool) {
	n, ok := nasByLower[strings.ToLower(name)]
	return n, ok
}

// NASByType finds a NAS message by its discriminator and message type.
func NASByType(epd, typ uint8) (*models.NASInfo, bool) {
	n, ok := nasByType[[2]uint8{epd, typ}]
	return n, ok
}

// NASNames returns every NAS message name, sorted.
func NASNames() []string {
	out := make([]string, len(nasTable))
	for i, n := range nasTable {
		out[i] = n.Name
	}
	sort.Strings(out)
	return out
}
