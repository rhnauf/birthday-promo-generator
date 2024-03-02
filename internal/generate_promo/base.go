package generate_promo

import "database/sql"

func Start(db *sql.DB, cronExpr string) {
	s := newService()
	c := newCronJob(s, cronExpr)

	c.RunCRON()
}
