package generate_promo

import "time"

type Promo struct {
	Id        int64
	UserId    int64
	PromoCode string
	Amount    float64
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
