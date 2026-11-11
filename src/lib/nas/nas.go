// Package libnas encodes and decodes 5GS NAS messages with optional NAS security.
package libnas

import (
	"crypto/subtle"
	"errors"
	"fmt"
	"io"
	"libnas/nasType"
	"libnas/security"
)

var (
	ErrIntegrity      = errors.New("NAS MAC verification failed")
	ErrReplay         = errors.New("replayed NAS COUNT")
	ErrCountExhausted = errors.New("NAS COUNT exhausted; establish fresh keys")
)

// Encode selects protection using ctx.SecurityHeaderType. nil ctx means plain
// NAS. Protected encode consumes one COUNT on success. NAS procedure retries
// must encode again with the next COUNT (TS 24.501 4.4.3.1); lower-layer
// retransmissions do not call Encode. The caller serializes the context.
func Encode(value any, ctx *EncodeContext) ([]byte, error) {
	plain, err := encodePlain(value)
	if err != nil {
		return nil, err
	}
	if ctx == nil {
		return plain, nil
	}
	if ctx.Direction > Downlink {
		return nil, fmt.Errorf("invalid NAS direction %d", ctx.Direction)
	}
	sht := ctx.SecurityHeaderType
	if sht == SecurityHeaderPlain {
		if ctx.Observe != nil {
			ctx.Observe(plain, plain)
		}
		return plain, nil
	}
	if err = validateSecurityHeader(sht, plain); err != nil {
		return nil, err
	}
	count, bearer, err := securityInputs(ctx.Security, ctx.Direction)
	if err != nil {
		return nil, err
	}
	if *count > nasType.MaxCount {
		return nil, ErrCountExhausted
	}
	payload := plain
	if ciphered(sht) {
		payload, err = security.Cipher(ctx.Security.CipheringAlgorithm, ctx.Security.KnasEnc, uint32(*count), bearer, uint8(ctx.Direction), plain)
		if err != nil {
			return nil, err
		}
	}
	if len(payload) > nasType.MaxMessageSize-7 {
		return nil, fmt.Errorf("NAS message too large")
	}
	out := make([]byte, 7+len(payload))
	out[0] = nasType.EPD5GMM
	out[1] = byte(sht)
	out[6] = count.SequenceNumber()
	copy(out[7:], payload)
	mac, err := security.Integrity(ctx.Security.IntegrityAlgorithm, ctx.Security.KnasInt, uint32(*count), bearer, uint8(ctx.Direction), out[6:])
	if err != nil {
		return nil, err
	}
	copy(out[2:6], mac[:])
	*count++
	if ctx.Observe != nil {
		ctx.Observe(out, plain)
	}
	return out, nil
}

// Decode checks EPD first: octet 2 in 5GSM is a session ID, not a security header.
// MAC is checked before decryption. On error both input and context are unchanged.
func Decode(data []byte, ctx *DecodeContext) (any, error) {
	if len(data) < 1 {
		return nil, io.ErrUnexpectedEOF
	}
	if len(data) > nasType.MaxMessageSize {
		return nil, fmt.Errorf("NAS message too large")
	}
	if ctx != nil && ctx.Direction > Downlink {
		return nil, fmt.Errorf("invalid NAS direction %d", ctx.Direction)
	}
	if data[0] == nasType.EPD5GSM {
		return decodeObserved(data, ctx)
	}
	if data[0] != nasType.EPD5GMM {
		return nil, fmt.Errorf("unsupported NAS EPD 0x%02x", data[0])
	}
	if len(data) < 2 {
		return nil, io.ErrUnexpectedEOF
	}
	if data[1]&0xf0 != 0 {
		return nil, fmt.Errorf("nonzero spare bits in NAS security header")
	}
	sht := SecurityHeaderType(data[1])
	if sht == SecurityHeaderPlain {
		return decodeObserved(data, ctx)
	}
	if sht > SecurityHeaderIntegrityProtectedAndCipheredWithNewContext {
		return nil, fmt.Errorf("unsupported NAS security header type %d", sht)
	}
	if len(data) < 10 {
		return nil, io.ErrUnexpectedEOF
	}
	if ctx == nil {
		return nil, fmt.Errorf("protected NAS requires a decode context")
	}
	next, bearer, err := securityInputs(ctx.Security, ctx.Direction)
	if err != nil {
		return nil, err
	}
	if *next > nasType.MaxCount {
		return nil, ErrCountExhausted
	}
	candidate := (*next &^ 0xff) | Count(data[6])
	verify := func(c Count) (bool, error) {
		mac, err := security.Integrity(ctx.Security.IntegrityAlgorithm, ctx.Security.KnasInt, uint32(c), bearer, uint8(ctx.Direction), data[6:])
		if err != nil {
			return false, err
		}
		return subtle.ConstantTimeCompare(mac[:], data[2:6]) == 1, nil
	}
	// Resolve an old SQN versus a forward rollover using the authenticated COUNT.
	if candidate < *next {
		if ctx.Security.IntegrityAlgorithm != security.NIA0 {
			valid, err := verify(candidate)
			if err != nil {
				return nil, err
			}
			if valid {
				return nil, ErrReplay
			}
		}
		candidate += 256
	}
	if candidate > nasType.MaxCount {
		return nil, ErrCountExhausted
	}
	valid, err := verify(candidate)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, ErrIntegrity
	}
	plain := append([]byte(nil), data[7:]...)
	if ciphered(sht) {
		plain, err = security.Cipher(ctx.Security.CipheringAlgorithm, ctx.Security.KnasEnc, uint32(candidate), bearer, uint8(ctx.Direction), plain)
		if err != nil {
			return nil, err
		}
	}
	if err = validateSecurityHeader(sht, plain); err != nil {
		return nil, err
	}
	msg, err := decodePlain(plain)
	if err != nil {
		return nil, err
	}
	*next = candidate + 1
	if ctx.Observe != nil {
		ctx.Observe(data, plain)
	}
	return msg, nil
}
func ciphered(sht SecurityHeaderType) bool {
	return sht == SecurityHeaderIntegrityProtectedAndCiphered || sht == SecurityHeaderIntegrityProtectedAndCipheredWithNewContext
}
func securityInputs(s *SecurityContext, d Direction) (*Count, uint8, error) {
	if s == nil {
		return nil, 0, fmt.Errorf("protected NAS requires SecurityContext")
	}
	if d > Downlink {
		return nil, 0, fmt.Errorf("invalid NAS direction %d", d)
	}
	bearer := s.NASConnectionID
	if bearer == 0 {
		bearer = 1
	}
	if bearer != 1 && bearer != 2 {
		return nil, 0, fmt.Errorf("NASConnectionID must be 1 or 2")
	}
	if d == Uplink {
		return &s.UplinkCount, bearer, nil
	}
	return &s.DownlinkCount, bearer, nil
}
func validateSecurityHeader(sht SecurityHeaderType, plain []byte) error {
	if sht < SecurityHeaderIntegrityProtected || sht > SecurityHeaderIntegrityProtectedAndCipheredWithNewContext {
		return fmt.Errorf("unsupported NAS security header type %d", sht)
	}
	h, err := nasType.DecodeHeader(plain)
	if err != nil {
		return err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM {
		return fmt.Errorf("wrap 5GSM inside UL/DL NAS Transport before applying NAS security")
	}
	if sht == SecurityHeaderIntegrityProtectedWithNewContext && h.MessageType != nasType.MsgSecurityModeCommand {
		return fmt.Errorf("new-context integrity header requires SecurityModeCommand")
	}
	if sht == SecurityHeaderIntegrityProtectedAndCipheredWithNewContext && h.MessageType != nasType.MsgSecurityModeComplete {
		return fmt.Errorf("new-context ciphering header requires SecurityModeComplete")
	}
	return nil
}

func decodeObserved(data []byte, ctx *DecodeContext) (any, error) {
	msg, err := decodePlain(data)
	if err == nil && ctx != nil && ctx.Observe != nil {
		ctx.Observe(data, data)
	}
	return msg, err
}
