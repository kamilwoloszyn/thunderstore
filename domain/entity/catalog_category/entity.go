package catalogcategory

import (
	"time"

	"github.com/gocql/gocql"
)

type CatalogCategoryEntity struct {
	EntityID       gocql.UUID `json:"entity_id"`
	AttributeSetID gocql.UUID `json:"attribute_set_id"`
	ParentID       gocql.UUID `json:"parent_id"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	Path           string     `json:"path"`
	Position       int        `json:"position"`
	Level          gocql.UUID `json:"level"`
	ChildrenCount  int        `json:"children_count"`
}
