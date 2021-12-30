package cataloginventory

import "github.com/gocql/gocql"

type CatalogInventoryStock struct {
	StockID   gocql.UUID `json:"stock_id"`
	WebsiteID gocql.UUID `json:"website_id"`
	StockName string     `json:"stock_name"`
}
