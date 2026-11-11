// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LinksToLog represents ASN.1 Links-to-log.
const (
	LinksToLogPresentUplink                Enumerated = 0
	LinksToLogPresentDownlink              Enumerated = 1
	LinksToLogPresentBothUplinkAndDownlink Enumerated = 2
)

type LinksToLog struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LinksToLog) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("Links-to-log", v) }

// DecodeLinksToLog decodes a standalone APER value and reports the consumed octets.
func DecodeLinksToLog(data []byte) (LinksToLog, int, error) {
	var v LinksToLog
	err := ngapcodec.DecodeParameter("Links-to-log", data, &v)
	if ngapcodec.IsFatal(err) {
		return LinksToLog{}, 0, err
	}
	return v, len(data), err
}
