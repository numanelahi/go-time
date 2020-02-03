build:
	go generate
	GOOS=linux CGO_ENABLED=0 go build

install:
	go generate
	GOOS=linux CGO_ENABLED=0 go install