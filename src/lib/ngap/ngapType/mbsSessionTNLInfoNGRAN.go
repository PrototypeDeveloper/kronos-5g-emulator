// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionTNLInfoNGRAN represents ASN.1 MBS-SessionTNLInfoNGRAN.
type MBSSessionTNLInfoNGRANPresent uint32

const (
	MBSSessionTNLInfoNGRANPresentNothing MBSSessionTNLInfoNGRANPresent = iota
	MBSSessionTNLInfoNGRANPresentLocationindependent
	MBSSessionTNLInfoNGRANPresentLocationdependent
	MBSSessionTNLInfoNGRANPresentChoiceExtensions
)

type MBSSessionTNLInfoNGRAN struct {
	Present             MBSSessionTNLInfoNGRANPresent
	Locationindependent UPTransportLayerInformation
	Locationdependent   MBSSessionTNLInfoNGRANList
	ChoiceExtensions    MBSSessionTNLInfoNGRANExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionTNLInfoNGRAN) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-SessionTNLInfoNGRAN", v)
}

// DecodeMBSSessionTNLInfoNGRAN decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionTNLInfoNGRAN(data []byte) (MBSSessionTNLInfoNGRAN, int, error) {
	var v MBSSessionTNLInfoNGRAN
	err := ngapcodec.DecodeParameter("MBS-SessionTNLInfoNGRAN", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionTNLInfoNGRAN{}, 0, err
	}
	return v, len(data), err
}
