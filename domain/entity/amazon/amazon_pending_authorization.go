package amazon

import (
	"time"

	"github.com/gocql/gocql"
)

type AmazonPendingAuthorization struct {
	EntityID        gocql.UUID `json:"entity_id"`
	OrderID         gocql.UUID `json:"order_id"`
	PaymentID       gocql.UUID `json:"payment_id"`
	AuthorizationID string     `json:"authorization_id"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	Processed       int        `json:"processed"`
	Capture         int        `json:"capture"`
	CaptureID       string     `json:"capture_id"`
}
