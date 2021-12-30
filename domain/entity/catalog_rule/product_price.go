package catalogrule

import (
	"time"

	"github.com/gocql/gocql"
)

type CatalogRuleProductPrice struct {
	RuleProductPriceID gocql.UUID `json:"rule_product_price_id "`
	RuleDate           *time.Time `json:"rule_date"`
	CustomerGroupID    gocql.UUID `json:"customer_group_id"`
	ProductID          gocql.UUID `json:"product_id"`
	RulePrice          float64    `json:"rule_price"`
	WebsiteID          gocql.UUID `json:"website_id"`
	LatestStartDate    *time.Time `json:"latest_start_date"`
	EarliestEndDate    *time.Time `json:"earliest_end_date"`
}
