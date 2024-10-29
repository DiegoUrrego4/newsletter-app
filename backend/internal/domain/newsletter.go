package domain

import "time"

type Newsletter struct {
	ID          string
	Title       string
	Content     string
	Attachment  string
	CreatedAt   time.Time
	ScheduledAt time.Time
}
