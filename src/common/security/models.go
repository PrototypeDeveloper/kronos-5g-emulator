// Structs of package security.

package security

import "crypto/cipher"

// Milenage computes the authentication functions f1, f1*, f2, f3, f4, f5 and
// f5* for one subscriber key (TS 35.206 §4.1).
type Milenage struct {
	block cipher.Block
	opc   [16]byte
}
