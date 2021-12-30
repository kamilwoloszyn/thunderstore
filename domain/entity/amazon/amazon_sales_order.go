package amazon

import "github.com/gocql/gocql"

type AmazonSalesOrder struct {
	EntityID               gocql.UUID `json:"entity_id"`
	OrderID                gocql.UUID `json:"order_id"`
	AmazonOrderReferenceID string     `json:"amazon_order_reference_id"`
}
