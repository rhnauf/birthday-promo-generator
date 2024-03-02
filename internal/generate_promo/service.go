package generate_promo

import "fmt"

type service struct{}

func newService() service {
	return service{}
}

func (s service) GeneratePromo() {
	fmt.Println("trigger every minute")

	// get users

	// generate promo code

	// send notification
}
