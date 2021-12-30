package authorization

import "github.com/gocql/gocql"

type AuthorizationRule struct {
	RuleID     gocql.UUID `json:"rule_id"`
	RoleID     gocql.UUID `json:"role_id"`
	ResourceID string     `json:"resource_id"`
	Privileges string     `json:"privileges"`
	Permission string     `json:"permission"`
}
