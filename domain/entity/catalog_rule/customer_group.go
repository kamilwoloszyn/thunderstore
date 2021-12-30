package catalogrule

import "github.com/gocql/gocql"

type CatalogRuleCustomerGroup struct {
	RuleID          gocql.UUID `json:"rule_id"`
	CustomerGroupID gocql.UUID `json:"customer_group_id"`
}
