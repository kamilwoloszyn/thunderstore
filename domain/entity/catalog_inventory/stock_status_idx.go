package cataloginventory

import "github.com/gocql/gocql"

type CatalogInventoryStockStatusIdx struct {
	ProductID   gocql.UUID `json:"product_id"`
	WebsiteID   gocql.UUID `json:"website_id"`
	StockID     gocql.UUID `json:"stock_id"`
	Qty         float64    `json:"qty"`
	StockStatus int        `json:"stock_status"`
}
