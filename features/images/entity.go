package images

import "time"

type Core struct {
	ID        uint
	UserID    uint
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
