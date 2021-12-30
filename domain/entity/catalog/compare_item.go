package catalog

import "github.com/gocql/gocql"

type CompareItem struct {
	CatalogCompareItemID gocql.UUID `json:"catalog_compare_item_id"`
	VisitorID            gocql.UUID `json:"visitor_id"`
	CustomerID           gocql.UUID `json:"customer_id"`
	ProductID            gocql.UUID `json:"product_id"`
	StoreID              gocql.UUID `json:"store_id"`
}
