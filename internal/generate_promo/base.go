package generate_promo

import "database/sql"

func Start(db *sql.DB, cronExpr string) {
	repository := newRepository(db)
	service := newService(repository)
	handler := newCronJob(service, cronExpr)

	handler.RunCRON()
}
