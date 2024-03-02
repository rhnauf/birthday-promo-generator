package main

import (
	"github.com/joho/godotenv"

	"github.com/rhnauf/birthday-promo-generator/external/db"
	promo "github.com/rhnauf/birthday-promo-generator/internal/generate_promo"
)

const CronExpr = "*/1 * * * *"

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// connect to db
	dbConn, err := db.NewDatabase()
	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	// start cron job
	promo.Start(dbConn, CronExpr)

	// bloc
	done := make(chan bool)
	<-done
}
