package catalogcategory

import "github.com/gocql/gocql"

type CatalogCategortyProduct struct {
	EntityID   gocql.UUID `json:"entity_id"`
	CategoryID gocql.UUID `json:"category_id"`
	ProductID  gocql.UUID `json:"product_id"`
	Position   int        `json:"position"`
}
