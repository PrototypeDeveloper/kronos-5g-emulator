// Constants of package plmn.

package plmn

// The PLMN identity.
const (
	// mccDigits is the number of digits of an MCC.
	mccDigits = 3
	// mncMinDigits and mncMaxDigits bound the number of digits of an MNC.
	mncMinDigits = 2
	mncMaxDigits = 3
	// encodedLen is the length of an encoded PLMN identity in octets.
	encodedLen = 3
	// fillerNibble stands in for the third MNC digit of a 2-digit MNC.
	fillerNibble = 0xf
)
