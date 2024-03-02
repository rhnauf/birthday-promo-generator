package main

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"

	"github.com/rhnauf/birthday-promo-generator/external/db"
	promo "github.com/rhnauf/birthday-promo-generator/internal/generate_promo"
)

// const CronExpr = "*/1 * * * *" // 1 minute for testing purposes
const CronExpr = "* * 1 * *" // daily

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// db connect
	dbConn, err := db.NewDatabase()
	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	// seed users table
	// err = seedUser(dbConn)
	// if err != nil {
	// 	panic(err)
	// }

	// start cron job
	promo.Start(dbConn, CronExpr)

	// blocking
	done := make(chan bool)
	<-done
}

func seedUser(db *sql.DB) error {
	qry := `
		INSERT INTO users (email, dob)
		VALUES
		('user_1@gmail.com', '2000-03-02'),
		('user_2@gmail.com', '2000-03-03'),
		('user_3@gmail.com', '2000-03-04'),
		('user_4@gmail.com', '2000-04-02'),
		('user_5@gmail.com', '1999-03-02'),
		('user_6@gmail.com', '1998-03-02');
	`

	_, err := db.Exec(qry)
	if err != nil {
		return err
	}

	log.Println("users table seeded successfully")

	return nil
}
