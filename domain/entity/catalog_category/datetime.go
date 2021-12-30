package catalogcategory

import (
	"time"

	"github.com/gocql/gocql"
)

type CatalogCategoryDateTime struct {
	ValueID     gocql.UUID `json:"value_id"`
	AttributeID gocql.UUID `json:"attribute_id"`
	StoreID     gocql.UUID `json:"store_id"`
	EntityID    gocql.UUID `json:"entity_id"`
	Value       *time.Time `json:"value"`
}
