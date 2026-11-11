// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRPagingTimeWindowForRRCINACTIVE represents ASN.1 NR-Paging-Time-Window-for-RRC-INACTIVE.
const (
	NRPagingTimeWindowForRRCINACTIVEPresentS1  Enumerated = 0
	NRPagingTimeWindowForRRCINACTIVEPresentS2  Enumerated = 1
	NRPagingTimeWindowForRRCINACTIVEPresentS3  Enumerated = 2
	NRPagingTimeWindowForRRCINACTIVEPresentS4  Enumerated = 3
	NRPagingTimeWindowForRRCINACTIVEPresentS5  Enumerated = 4
	NRPagingTimeWindowForRRCINACTIVEPresentS6  Enumerated = 5
	NRPagingTimeWindowForRRCINACTIVEPresentS7  Enumerated = 6
	NRPagingTimeWindowForRRCINACTIVEPresentS8  Enumerated = 7
	NRPagingTimeWindowForRRCINACTIVEPresentS9  Enumerated = 8
	NRPagingTimeWindowForRRCINACTIVEPresentS10 Enumerated = 9
	NRPagingTimeWindowForRRCINACTIVEPresentS11 Enumerated = 10
	NRPagingTimeWindowForRRCINACTIVEPresentS12 Enumerated = 11
	NRPagingTimeWindowForRRCINACTIVEPresentS13 Enumerated = 12
	NRPagingTimeWindowForRRCINACTIVEPresentS14 Enumerated = 13
	NRPagingTimeWindowForRRCINACTIVEPresentS15 Enumerated = 14
	NRPagingTimeWindowForRRCINACTIVEPresentS16 Enumerated = 15
	NRPagingTimeWindowForRRCINACTIVEPresentS17 Enumerated = 16
	NRPagingTimeWindowForRRCINACTIVEPresentS18 Enumerated = 17
	NRPagingTimeWindowForRRCINACTIVEPresentS19 Enumerated = 18
	NRPagingTimeWindowForRRCINACTIVEPresentS20 Enumerated = 19
	NRPagingTimeWindowForRRCINACTIVEPresentS21 Enumerated = 20
	NRPagingTimeWindowForRRCINACTIVEPresentS22 Enumerated = 21
	NRPagingTimeWindowForRRCINACTIVEPresentS23 Enumerated = 22
	NRPagingTimeWindowForRRCINACTIVEPresentS24 Enumerated = 23
	NRPagingTimeWindowForRRCINACTIVEPresentS25 Enumerated = 24
	NRPagingTimeWindowForRRCINACTIVEPresentS26 Enumerated = 25
	NRPagingTimeWindowForRRCINACTIVEPresentS27 Enumerated = 26
	NRPagingTimeWindowForRRCINACTIVEPresentS28 Enumerated = 27
	NRPagingTimeWindowForRRCINACTIVEPresentS29 Enumerated = 28
	NRPagingTimeWindowForRRCINACTIVEPresentS30 Enumerated = 29
	NRPagingTimeWindowForRRCINACTIVEPresentS31 Enumerated = 30
	NRPagingTimeWindowForRRCINACTIVEPresentS32 Enumerated = 31
)

type NRPagingTimeWindowForRRCINACTIVE struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:31"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRPagingTimeWindowForRRCINACTIVE) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NR-Paging-Time-Window-for-RRC-INACTIVE", v)
}

// DecodeNRPagingTimeWindowForRRCINACTIVE decodes a standalone APER value and reports the consumed octets.
func DecodeNRPagingTimeWindowForRRCINACTIVE(data []byte) (NRPagingTimeWindowForRRCINACTIVE, int, error) {
	var v NRPagingTimeWindowForRRCINACTIVE
	err := ngapcodec.DecodeParameter("NR-Paging-Time-Window-for-RRC-INACTIVE", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRPagingTimeWindowForRRCINACTIVE{}, 0, err
	}
	return v, len(data), err
}
