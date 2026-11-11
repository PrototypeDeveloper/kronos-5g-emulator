// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NGRANCGI represents ASN.1 NGRAN-CGI.
type NGRANCGIPresent uint32

const (
	NGRANCGIPresentNothing NGRANCGIPresent = iota
	NGRANCGIPresentNRCGI
	NGRANCGIPresentEUTRACGI
	NGRANCGIPresentChoiceExtensions
)

type NGRANCGI struct {
	Present          NGRANCGIPresent
	NRCGI            NRCGI
	EUTRACGI         EUTRACGI
	ChoiceExtensions NGRANCGIExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGRANCGI) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("NGRAN-CGI", v) }

// DecodeNGRANCGI decodes a standalone APER value and reports the consumed octets.
func DecodeNGRANCGI(data []byte) (NGRANCGI, int, error) {
	var v NGRANCGI
	err := ngapcodec.DecodeParameter("NGRAN-CGI", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGRANCGI{}, 0, err
	}
	return v, len(data), err
}
