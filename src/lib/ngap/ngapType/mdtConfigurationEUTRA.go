// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MDTConfigurationEUTRA represents ASN.1 MDT-Configuration-EUTRA.
type MDTConfigurationEUTRA struct {
	MdtActivation                 MDTActivation
	AreaScopeOfMDT                AreaScopeOfMDTEUTRA
	MDTMode                       MDTModeEutra
	HasSignallingBasedMDTPLMNList bool
	SignallingBasedMDTPLMNList    MDTPLMNList
	HasIEExtensions               bool
	IEExtensions                  []MDTConfigurationEUTRAExtIEsEntry
	UnknownExtensions             []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MDTConfigurationEUTRA) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MDT-Configuration-EUTRA", v)
}

// DecodeMDTConfigurationEUTRA decodes a standalone APER value and reports the consumed octets.
func DecodeMDTConfigurationEUTRA(data []byte) (MDTConfigurationEUTRA, int, error) {
	var v MDTConfigurationEUTRA
	err := ngapcodec.DecodeParameter("MDT-Configuration-EUTRA", data, &v)
	if ngapcodec.IsFatal(err) {
		return MDTConfigurationEUTRA{}, 0, err
	}
	return v, len(data), err
}
