package generate_promo

import "time"

type User struct {
	Id        int64
	Email     string
	DoB       time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
