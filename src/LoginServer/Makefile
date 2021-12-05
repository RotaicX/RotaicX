GO = go

upgomod:
	$(GO) mod tidy
	$(GO) get -u

build:
	$(GO) build -o build/LoginServer LoginServer.go

configproxy:
	$(GO) env -w GOPROXY=https://goproxy.cn,direct

all: upgomod build