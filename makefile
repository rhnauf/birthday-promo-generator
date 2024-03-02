run:
	go run ./cmd/app/

hot-reload:
	nodemon --exec go run ./cmd/app/. --ext go
