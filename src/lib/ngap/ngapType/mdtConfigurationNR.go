// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MDTConfigurationNR represents ASN.1 MDT-Configuration-NR.
type MDTConfigurationNR struct {
	MdtActivation                 MDTActivation
	AreaScopeOfMDT                AreaScopeOfMDTNR
	MDTModeNr                     MDTModeNr
	HasSignallingBasedMDTPLMNList bool
	SignallingBasedMDTPLMNList    MDTPLMNList
	HasIEExtensions               bool
	IEExtensions                  []MDTConfigurationNRExtIEsEntry
	UnknownExtensions             []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MDTConfigurationNR) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MDT-Configuration-NR", v)
}

// DecodeMDTConfigurationNR decodes a standalone APER value and reports the consumed octets.
func DecodeMDTConfigurationNR(data []byte) (MDTConfigurationNR, int, error) {
	var v MDTConfigurationNR
	err := ngapcodec.DecodeParameter("MDT-Configuration-NR", data, &v)
	if ngapcodec.IsFatal(err) {
		return MDTConfigurationNR{}, 0, err
	}
	return v, len(data), err
}
