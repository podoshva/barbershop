package domain

import (
	"time"
)

type Order struct {
	ID            int64
	ProfileID     int64
	BranchID      int64
	Date          time.Time
	CustomerPhone string
	Description   string
	Status        string
	CreatedAt     time.Time
}
