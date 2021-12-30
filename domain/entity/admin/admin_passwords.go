package admin

import (
	"time"

	"github.com/gocql/gocql"
)

type AdminPasswords struct {
	PasswordID   gocql.UUID `json:"password_id"`
	UserID       gocql.UUID `json:"user_id"`
	PasswordHash string     `json:"password_hash"`
	Expires      *time.Time `json:"expires"`
	LastUpdated  time.Time  `json:"last_updated"`
}
