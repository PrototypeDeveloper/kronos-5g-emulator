// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionTNLInfo5GC represents ASN.1 MBS-SessionTNLInfo5GC.
type MBSSessionTNLInfo5GCPresent uint32

const (
	MBSSessionTNLInfo5GCPresentNothing MBSSessionTNLInfo5GCPresent = iota
	MBSSessionTNLInfo5GCPresentLocationindependent
	MBSSessionTNLInfo5GCPresentLocationdependent
	MBSSessionTNLInfo5GCPresentChoiceExtensions
)

type MBSSessionTNLInfo5GC struct {
	Present             MBSSessionTNLInfo5GCPresent
	Locationindependent SharedNGUMulticastTNLInformation
	Locationdependent   MBSSessionTNLInfo5GCList
	ChoiceExtensions    MBSSessionTNLInfo5GCExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionTNLInfo5GC) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-SessionTNLInfo5GC", v)
}

// DecodeMBSSessionTNLInfo5GC decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionTNLInfo5GC(data []byte) (MBSSessionTNLInfo5GC, int, error) {
	var v MBSSessionTNLInfo5GC
	err := ngapcodec.DecodeParameter("MBS-SessionTNLInfo5GC", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionTNLInfo5GC{}, 0, err
	}
	return v, len(data), err
}
