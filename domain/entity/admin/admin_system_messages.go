package admin

import "time"

type AdminSystemMessages struct {
	Identity  string     `json:"identity"`
	Severity  int        `json:"severity"`
	CreatedAt *time.Time `json:"created_at"`
}
