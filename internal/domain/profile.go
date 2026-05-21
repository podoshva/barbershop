package domain

import "time"

type Profile struct {
	ID        int64
	BranchID  int64
	FullName  string
	Login     string
	Password  string
	Role      string
	CreatedAt time.Time
}
