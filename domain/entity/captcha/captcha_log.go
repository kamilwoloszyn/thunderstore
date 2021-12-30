package captcha

import "time"

type CaptchaLog struct {
	Type      string     `json:"type"`
	Value     string     `json:"value"`
	Count     int        `json:"count"`
	UpdatedAt *time.Time `json:"updated_at"`
}
