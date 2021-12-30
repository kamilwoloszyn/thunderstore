package admin

import "github.com/gocql/gocql"

type AdminAnalyticsUsageVersionLog struct {
	ID                  gocql.UUID `json:"id"`
	LastViewedInVersion string     `json:"last_viewed_in_version"`
}
