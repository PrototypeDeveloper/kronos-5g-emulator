// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MDTConfiguration represents ASN.1 MDT-Configuration.
type MDTConfiguration struct {
	HasMdtConfigNR    bool
	MdtConfigNR       MDTConfigurationNR
	HasMdtConfigEUTRA bool
	MdtConfigEUTRA    MDTConfigurationEUTRA
	HasIEExtensions   bool
	IEExtensions      []MDTConfigurationExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MDTConfiguration) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MDT-Configuration", v)
}

// DecodeMDTConfiguration decodes a standalone APER value and reports the consumed octets.
func DecodeMDTConfiguration(data []byte) (MDTConfiguration, int, error) {
	var v MDTConfiguration
	err := ngapcodec.DecodeParameter("MDT-Configuration", data, &v)
	if ngapcodec.IsFatal(err) {
		return MDTConfiguration{}, 0, err
	}
	return v, len(data), err
}
