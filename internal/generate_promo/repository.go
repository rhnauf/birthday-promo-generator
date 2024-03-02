package generate_promo

import (
	"database/sql"
	"log"
	"time"
)

type repository struct {
	db *sql.DB
}

func newRepository(db *sql.DB) repository {
	return repository{db: db}
}

func (r repository) GetUsersBirthdayToday() ([]User, error) {
	qry := `
		SELECT id, email, dob FROM users
		WHERE
			DATE_PART('DAY', dob) = $1
			AND DATE_PART('MONTH', dob) = $2
	`

	currDate := time.Now()
	day := currDate.Day()

	// get month in number, March => 3
	month := int(currDate.Month())

	rows, err := r.db.Query(qry, day, month)
	if err != nil {
		log.Println("error query =>", err)
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Id, &u.Email, &u.DoB); err != nil {
			log.Println("error scanning rows =>", err)
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (r repository) GeneratePromo(promo Promo) error {
	qry := `
		INSERT INTO promos (user_id, promo_code, amount, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(qry,
		promo.UserId,
		promo.PromoCode,
		promo.Amount,
		promo.StartDate,
		promo.EndDate,
	)
	if err != nil {
		log.Println("error query =>", err)
		return err
	}

	return nil
}
