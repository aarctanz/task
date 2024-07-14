.DEFAULT_GOAL:=build

run:
	go run src/main.go

build: install
	go build -o task src/main.go
	sudo mv task /usr/bin


install:
	go mod download