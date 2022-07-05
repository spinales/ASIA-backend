package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(tuition string, duration time.Duration) (string, error) {
	id := uuid.New().String()
	payload := paseto.JSONToken{
		Issuer:     tuition, // matricula
		Jti:        id,
		IssuedAt:   time.Now(),
		Expiration: time.Now().Add(duration),
		NotBefore:  time.Now(),
	}
	// footer := "some footer"

	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}

func (maker *PasetoMaker) VerifyToken(token string) (*paseto.JSONToken, error) {
	payload := &paseto.JSONToken{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, err
	}

	if time.Now().After(payload.Expiration) {
		return nil, ErrExpiredToken
	}

	return payload, nil
}
