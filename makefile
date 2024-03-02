run:
	go run ./cmd/app/

hot-reload:
	nodemon --exec go run ./cmd/app/. --ext go

up:
	 cd ./sql/schema && goose postgres postgres://root:root@localhost:5432/birthday-promo up

down:
	 cd ./sql/schema && goose postgres postgres://root:root@localhost:5432/birthday-promo down

build:
	go build -C cmd/app -o ../../birthday-promo.exe

build-run:
	go build -C cmd/app -o ../../birthday-promo.exe && ./birthday-promo.exe