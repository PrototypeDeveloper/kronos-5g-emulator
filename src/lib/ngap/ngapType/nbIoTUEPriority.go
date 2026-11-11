// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NBIoTUEPriority represents ASN.1 NB-IoT-UEPriority.
type NBIoTUEPriority struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:255"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NBIoTUEPriority) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NB-IoT-UEPriority", v)
}

// DecodeNBIoTUEPriority decodes a standalone APER value and reports the consumed octets.
func DecodeNBIoTUEPriority(data []byte) (NBIoTUEPriority, int, error) {
	var v NBIoTUEPriority
	err := ngapcodec.DecodeParameter("NB-IoT-UEPriority", data, &v)
	if ngapcodec.IsFatal(err) {
		return NBIoTUEPriority{}, 0, err
	}
	return v, len(data), err
}
