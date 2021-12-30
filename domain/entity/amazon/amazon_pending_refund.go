package amazon

import (
	"time"

	"github.com/gocql/gocql"
)

type AmazonPendingRefund struct {
	EntityID  gocql.UUID `json:"entity_id"`
	RefundID  string     `json:"refund_id"`
	CreatedAt *time.Time `json:"created_at"`
	OrderID   gocql.UUID `json:"order_id"`
	PaymentID gocql.UUID `json:"payment_id"`
}
