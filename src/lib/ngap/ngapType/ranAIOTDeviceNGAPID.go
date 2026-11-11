// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RANAIOTDeviceNGAPID represents ASN.1 RAN-AIOT-Device-NGAP-ID.
type RANAIOTDeviceNGAPID struct {
	Value int64 `aper:"valueLB:0,valueUB:4294967295"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANAIOTDeviceNGAPID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RAN-AIOT-Device-NGAP-ID", v)
}

// DecodeRANAIOTDeviceNGAPID decodes a standalone APER value and reports the consumed octets.
func DecodeRANAIOTDeviceNGAPID(data []byte) (RANAIOTDeviceNGAPID, int, error) {
	var v RANAIOTDeviceNGAPID
	err := ngapcodec.DecodeParameter("RAN-AIOT-Device-NGAP-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANAIOTDeviceNGAPID{}, 0, err
	}
	return v, len(data), err
}
