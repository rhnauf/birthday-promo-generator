package generate_promo

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

type Service interface {
	GeneratePromo()
}

type cronJob struct {
	service  Service
	cronExpr string
}

func newCronJob(s Service, cronExpression string) cronJob {
	return cronJob{service: s, cronExpr: cronExpression}
}

func (c cronJob) RunCRON() func() context.Context {
	tz, err := time.LoadLocation(os.Getenv("TIMEZONE"))
	if err != nil {
		log.Fatal(err)
	}

	scheduler := cron.New(cron.WithLocation(tz))

	// [min (0-59)] [hour (0-23)] [day (1-31)] [month (1-12)] [weekday (0-7)]
	scheduler.AddFunc(c.cronExpr, c.service.GeneratePromo)

	go scheduler.Start()

	return scheduler.Stop
}
