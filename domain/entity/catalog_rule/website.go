package catalogrule

import "github.com/gocql/gocql"

type CatalogRuleWebsite struct {
	RuleID    gocql.UUID `json:"rule_id"`
	WebsiteID gocql.UUID `json:"website_id"`
}
