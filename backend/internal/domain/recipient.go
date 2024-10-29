package domain

import "time"

type Recipient struct {
	Email          string
	Subscribed     bool
	UnsubscribedAt time.Time
}
