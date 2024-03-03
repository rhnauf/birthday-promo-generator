# Birthday Promo Generator

for disclaimer, this is just a PoC so it is not suitable for production purposes. Couple of improvements that we can do is:
* Instead of using CRON job, we can implement a queue system using 3rd party application/service like [RabbitMQ](https://rabbitmq-website.pages.dev/) or [Temporal Workflow](https://temporal.io/), so everytime *user* register or an entry being inserted to the *users* table we can schedule a **delayed** queue to that 3rd party service then it will get sent/processed on the specific time or in this case on user birthday. So why we need to do that? for small application scale it is fine using CRON job, but when we want to do horizontal scaling on our server/service then this leads to a problem where the CRON job being invoked n number of times, forcing us to separate this service from our main server(business logic, api, etc) and tightly coupled with vertical scaling only.
* When we fetch *users* using ```GetUsersBirthdayToday``` function we store it to a ```[]User``` object saving it to the RAM, this leads to a problem where the bigger the rows goes (ex: 1.000.000 rows being fetched) then it also will consume more RAM and burden the CPU. So to handle this we could iterate ```rows.Scan()``` function and process the *user* concurrently over to a channel implementing [worker pool](https://gobyexample.com/worker-pools) like [this](https://stackoverflow.com/a/33662810).
* Integrate proper email/whatsapp messaging, for whatsapp we can use [Qontak](https://qontak.com/) and [SendGrid](https://sendgrid.com/en-us) for email. I suppose it is outside of the test scope, because the setup is pretty complicated(creating an account, creating message template, etc) if we want to do proper messaging.
---
run docker for postgres instance

```docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:15-alpine```

run migrations, using your own preferred tools (example using goose)

```cd ./sql/schema && goose postgres postgres://root:root@localhost:5432/birthday-promo up```

create env files & modify to your own configurations

```cp .env.example .env```

to run

```make run``` or ```go run ./cmd/app/```

alternatively if you prefer to build binary and run

```make build-run``` or ```go build -C cmd/app -o ../../birthday-promo.exe && ./birthday-promo.exe```

---
<ins>Flowchart</ins>

![flowchart](./flowchart.png)

The flow of this application is pretty straightforward, it is just consist of a go service and a postgres database instance.
 We fetch *users* data from db and also generate the promo based on the *user* that whose birthday is that day.
In real world we might use 3rd party service to send the messaging through email/whatsapp, but for this context I will just mock it
using ```time.Sleep()``` function.

<ins>Table Design</ins>

![table design](./erd.png)
