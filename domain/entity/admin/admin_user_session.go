package admin

import (
	"time"

	"github.com/gocql/gocql"
)

type AdminUserSession struct {
	ID        gocql.UUID `json:"id"`
	SessionID string     `json:"session_id"`
	UserID    string     `json:"user_id"`
	Status    int        `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	IP        string     `json:"ip"`
}
