package types

import "time"

type WorkersRecord struct {
	Id          int
	UserId      int
	UUID        string
	Fullname    string
	Active      bool
	Avg_ratings float32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
