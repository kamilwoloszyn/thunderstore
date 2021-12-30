package amazon

import "github.com/gocql/gocql"

type AmazonCustomer struct {
	EntityID   gocql.UUID `json:"amazon_customer"`
	CustomerID gocql.UUID `json:"customer_id"`
	AmazonID   string     `json:"amazon_id"`
}
