package token

import (
	"time"

	"github.com/o1egl/paseto"
)

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*paseto.JSONToken, error)
}
