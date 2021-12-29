package catalogrule

import (
	"time"

	"github.com/gocql/gocql"
)

type CatalogRule struct {
	RuleID               gocql.UUID `json:"rule_id"`
	Name                 string     `json:"name"`
	Description          string     `json:"description"`
	FromDate             *time.Time `json:"from_date"`
	ToDate               *time.Time `json:"to_date"`
	IsActive             int        `json:"is_active"`
	ConditionsSerialized string     `json:"conditions_serialized"`
	ActionsSerialized    string     `json:"actions_serialized"`
	StopRulesProcessing  int        `json:"stop_rules_processing"`
	SortOrder            int        `json:"sort_order"`
	SimpleAction         string     `json:"simple_action"`
	DiscountAmount       float64    `json:"discount_amount"`
}
