test:
	go test -race ./... -coverprofile=profile.out -count=1 -covermode=atomic

testv:
	go test -race ./... -v -coverprofile=profile.out -count=1 -covermode=atomic

coverage:
	go tool cover -html=profile.out

compile:
	./compile.sh

validate:
	go build -v -o rei && ./rei validate

scaffold:
	go build -v -o rei && ./rei scaffold

build:
	go build -v -o rei && ./rei build