package token

import (
	"time"

	"github.com/google/uuid"
)

// Maker is an interface  for managing tokens
type Maker interface {
	/// CreateToken creates a new token for a specifica username and duration
	CreateToken(userId uuid.UUID, duration time.Duration, access bool) (string, *Payload, error)

	/// VerifiyToken checks if the token is valid or not
	VerifiyToken(token string) (*Payload, error)
}
