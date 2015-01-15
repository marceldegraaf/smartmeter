.PHONY: build

default: build

build:
	go build -o release/smartmeter main.go

build-pi:
	GOOS=linux GOARCH=arm GOARM=6 go build -o release/smartmeter-pi main.go

release: build-pi
	scp ./release/smartmeter-pi pi@pi:~/smartmeter
