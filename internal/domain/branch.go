package domain

import "time"

type Branch struct {
	ID        int64
	Name      string
	CreatedAt time.Time
}
