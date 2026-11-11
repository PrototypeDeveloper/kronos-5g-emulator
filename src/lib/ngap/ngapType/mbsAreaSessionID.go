// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSAreaSessionID represents ASN.1 MBS-AreaSessionID.
type MBSAreaSessionID struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSAreaSessionID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-AreaSessionID", v)
}

// DecodeMBSAreaSessionID decodes a standalone APER value and reports the consumed octets.
func DecodeMBSAreaSessionID(data []byte) (MBSAreaSessionID, int, error) {
	var v MBSAreaSessionID
	err := ngapcodec.DecodeParameter("MBS-AreaSessionID", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSAreaSessionID{}, 0, err
	}
	return v, len(data), err
}
