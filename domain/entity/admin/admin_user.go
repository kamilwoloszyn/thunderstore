package admin

import (
	"time"

	"github.com/gocql/gocql"
)

type AdminUser struct {
	UserID           gocql.UUID `json:"user_id"`
	FirstName        string     `json:"firstname"`
	LastName         string     `json:"lastname"`
	Email            string     `json:"email"`
	Username         string     `json:"username"`
	Password         string     `json:"password"`
	Created          *time.Time `json:"created"`
	Modified         *time.Time `json:"modified"`
	Logdate          *time.Time `json:"logdate"`
	Lognum           int        `json:"lognum"`
	ReloadAclFlag    int        `json:"reload_acl_flag"`
	IsActive         int        `json:"is_active"`
	Extra            string     `json:"extra"`
	RpToken          string     `json:"rp_token"`
	RpTokenCreatedAt *time.Time `json:"rp_token_created_at"`
	InterfaceLocale  string     `json:"interface_locale"`
	FailuresNum      int        `json:"failures_num"`
	FirstFailure     *time.Time `json:"first_failure"`
	LockExpires      *time.Time `json:"lock_expires"`
	RefreshToken     string     `json:"refresh_token"`
}
