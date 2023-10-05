package token

import "time"

// the interface for managing tokens
type Maker interface {
	//creates a new token given a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	//checks if a given token is valid
	VerifyToken(token string) (*Payload, error)
}
