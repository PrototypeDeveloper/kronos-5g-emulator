package ngapcodec

import "libngap/aper"

// NewSchema returns an independent copy of the NGAP ASN.1 schema, constructed
// from Go declarations. No filesystem access or JSON parsing is performed.
func NewSchema() aper.Schema {
	return aper.Schema{Types: schemaTypes(), Names: schemaNames(), Sets: schemaSets()}
}
