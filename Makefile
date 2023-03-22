dev:
	docker-compose up -d

dist: clean
	docker compose run front npm run build
	cp frontend/dist.go.tmpl frontend/dist/dist.go
	go build -tags=dist -o bin/docmonk.bin ./cmd/docmonk.go

build:
	go build -o bin/docmonk.bin ./cmd/docmonk.go

clean:
	rm -rf frontend/dist

.PHONY: dev build clean
