package amazon

import (
	"time"

	"github.com/gocql/gocql"
)

type AmazonPendingCapture struct {
	EntityID  gocql.UUID `json:"entity_id"`
	CaptureID string     `json:"capture_id"`
	CreatedAt *time.Time `json:"created_at"`
	OrderID   gocql.UUID `json:"order_id"`
	PaymentID gocql.UUID `json:"payment_id"`
}
