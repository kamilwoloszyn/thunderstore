package cataloginventory

import "github.com/gocql/gocql"

type CatalogInventoryStockStatus struct {
	ProductID   gocql.UUID `json:"product_id"`
	WebsiteID   gocql.UUID `json:"website_id"`
	StockID     gocql.UUID `json:"stock_id"`
	Qty         int        `json:"qty"`
	StockStatus int        `json:"stock_status"`
}
