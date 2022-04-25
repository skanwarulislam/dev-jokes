default: install app/resources app/config app/local

install:
	go get ./...

build: install  **/*.go
	go build .

test: install
	go test ./...

app:
	mkdir -p app

app/config: app
	# Copy config files
	cp config.yml app/

app/resources: app
	 cp -r templates app/

app/server: app **/*.go
	# Creates a linux build for copying into the alpine container
    GOOS=linux GOARCH=amd64 go build -ldflags '-linkmode=external' -o ./app/server .

app/local: app **/*.go
	go build -ldflags '-linkmode=external' -o ./app/server .

deploy: docker-compose.yml ./app/server install app/resources app/config app/local
	docker-compose up --build

clean:
	rm -rf ./app ./main

.PHONY: install test app clean