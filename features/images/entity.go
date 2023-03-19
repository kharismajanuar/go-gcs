package images

import "time"

type Core struct {
	ID        uint
	UserID    uint
	Url       string `validate:"required,max=50"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
