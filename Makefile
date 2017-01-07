all:  mac linux

mac:
	@mkdir -p bin
	@GOOS=darwin GOARCH=amd64 go build -o ./bin/eth0.me.darwin.amd64 *.go

linux:
	@mkdir -p bin
	@GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o ./bin/eth0.me.linux.amd64 *.go

clean:
	@rm -vf ./bin/eth0.me.darwin.amd64 || true
	@rm -vf ./bin/eth0.me.linux.amd64  || true

.PHONY: all
