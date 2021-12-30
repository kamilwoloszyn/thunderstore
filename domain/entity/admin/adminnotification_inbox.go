package admin

import (
	"time"

	"github.com/gocql/gocql"
)

type AdminNotificationInbox struct {
	NotificationID gocql.UUID `json:"notification_id"`
	Severity       int        `json:"severity"`
	DateAdded      *time.Time `json:"date_added"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	URL            string     `json:"url"`
	IsRead         int        `json:"is_read"`
	IsRemove       int        `json:"is_remove"`
}
