// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionTNLInfoNGRANItem represents ASN.1 MBS-SessionTNLInfoNGRANItem.
type MBSSessionTNLInfoNGRANItem struct {
	MBSAreaSessionID                  MBSAreaSessionID
	HasSharedNGUUnicastTNLInformation bool
	SharedNGUUnicastTNLInformation    UPTransportLayerInformation
	HasIEExtensions                   bool
	IEExtensions                      []MBSSessionTNLInfoNGRANItemExtIEsEntry
	UnknownExtensions                 []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionTNLInfoNGRANItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-SessionTNLInfoNGRANItem", v)
}

// DecodeMBSSessionTNLInfoNGRANItem decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionTNLInfoNGRANItem(data []byte) (MBSSessionTNLInfoNGRANItem, int, error) {
	var v MBSSessionTNLInfoNGRANItem
	err := ngapcodec.DecodeParameter("MBS-SessionTNLInfoNGRANItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionTNLInfoNGRANItem{}, 0, err
	}
	return v, len(data), err
}
