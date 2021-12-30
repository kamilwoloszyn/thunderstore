package catalogrule

import "github.com/gocql/gocql"

type CatalogRuleGroupWebsite struct {
	RuleID          gocql.UUID `json:"rule_id"`
	CustomerGroupID gocql.UUID `json:"customer_group_id"`
	WebsiteID       gocql.UUID `json:"website_id"`
}
