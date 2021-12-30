package catalogrule

import (
	"time"

	"github.com/gocql/gocql"
)

type CatalogRuleProduct struct {
	RuleProductID   gocql.UUID `json:"rule_product_id"`
	RuleID          gocql.UUID `json:"rule_id"`
	FromTime        *time.Time `json:"from_time"`
	ToTime          *time.Time `json:"to_time"`
	CustomerGroupID gocql.UUID `json:"customer_group_id"`
	ProductID       gocql.UUID `json:"product_id"`
	ActionOperator  string     `json:"action_operator"`
	ActionAmount    float64    `json:"action_amount"`
	ActionStop      int        `json:"action_stop"`
	SortOrder       int        `json:"sort_order"`
	WebsiteID       int        `json:"website_id"`
}
