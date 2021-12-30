package authorization

import "github.com/gocql/gocql"

type AuthorizationRole struct {
	RoleID    gocql.UUID `json:"role_id"`
	ParentId  gocql.UUID `json:"parent_id"`
	TreeLevel int        `json:"tree_level"`
	SortOrder int        `json:"sort_order"`
	RoleType  string     `json:"role_type"`
	UserID    gocql.UUID `json:"user_id"`
	UserType  string     `json:"user_type"`
	RoleName  string     `json:"role_name"`
}
