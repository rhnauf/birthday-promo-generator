package generate_promo

import (
	"fmt"
	"log"
	"time"

	"github.com/rhnauf/birthday-promo-generator/util"
)

type Repository interface {
	GetUsersBirthdayToday(currDate time.Time) ([]User, error)
	GeneratePromo(promo Promo) error
}

type service struct {
	repository Repository
}

func newService(repository Repository) service {
	return service{repository: repository}
}

const PromoCodeLen = 3

func (s service) GeneratePromo() {
	fmt.Println("CRON Job Starting...")

	currDate := time.Now()

	// get users whos birthday today
	users, err := s.repository.GetUsersBirthdayToday(currDate)
	if err != nil {
		return
	}

	// if no user found, then return early
	if len(users) == 0 {
		log.Println("no users found")
		return
	}

	// generate promo code for each birthday user
	for _, u := range users {
		promoCode := util.GeneratePromoCode(PromoCodeLen)
		promoAmount := util.GeneratePromoAmount()

		promo := Promo{
			UserId:    u.Id,
			PromoCode: promoCode,
			Amount:    float64(promoAmount),
			StartDate: currDate,
			EndDate:   currDate.AddDate(0, 1, 0), // expired after 1 month
		}

		err := s.repository.GeneratePromo(promo)
		if err != nil {
			continue
		}
		// log.Printf("%+v\n", u)
		go sendNotification(u.Email)
	}
}

/*
in real world scenario we would need to integrate 3rd party application/service
such as SendGrid to send email or Qontak to send whatsapp or maybe rabbitmq to do some preprocessing
and offload the server loads,
but that is outside of the scope for this test I suppose, so I will just mock the sending message process
*/
func sendNotification(target string) {
	time.Sleep(2 * time.Second)
	log.Printf("successfully send message to %s", target)
}
