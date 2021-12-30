package catalogcategory

import "github.com/gocql/gocql"

type CatalogCategortyEntityText struct {
	ValueID     gocql.UUID `json:"value_id"`
	AttributeID gocql.UUID `json:"attribute_id"`
	StoreID     gocql.UUID `json:"store_id"`
	EntityID    gocql.UUID `json:"entity_id"`
	Value       string     `json:"value"`
}
