// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EUTRAPagingEDRXCycle represents ASN.1 EUTRA-Paging-eDRX-Cycle.
const (
	EUTRAPagingEDRXCyclePresentHfhalf Enumerated = 0
	EUTRAPagingEDRXCyclePresentHf1    Enumerated = 1
	EUTRAPagingEDRXCyclePresentHf2    Enumerated = 2
	EUTRAPagingEDRXCyclePresentHf4    Enumerated = 3
	EUTRAPagingEDRXCyclePresentHf6    Enumerated = 4
	EUTRAPagingEDRXCyclePresentHf8    Enumerated = 5
	EUTRAPagingEDRXCyclePresentHf10   Enumerated = 6
	EUTRAPagingEDRXCyclePresentHf12   Enumerated = 7
	EUTRAPagingEDRXCyclePresentHf14   Enumerated = 8
	EUTRAPagingEDRXCyclePresentHf16   Enumerated = 9
	EUTRAPagingEDRXCyclePresentHf32   Enumerated = 10
	EUTRAPagingEDRXCyclePresentHf64   Enumerated = 11
	EUTRAPagingEDRXCyclePresentHf128  Enumerated = 12
	EUTRAPagingEDRXCyclePresentHf256  Enumerated = 13
)

type EUTRAPagingEDRXCycle struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:13"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRAPagingEDRXCycle) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRA-Paging-eDRX-Cycle", v)
}

// DecodeEUTRAPagingEDRXCycle decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRAPagingEDRXCycle(data []byte) (EUTRAPagingEDRXCycle, int, error) {
	var v EUTRAPagingEDRXCycle
	err := ngapcodec.DecodeParameter("EUTRA-Paging-eDRX-Cycle", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRAPagingEDRXCycle{}, 0, err
	}
	return v, len(data), err
}
