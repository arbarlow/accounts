# vi: ft=make
.PHONY: proto test get ci docker

proto:
	protoc -I $$GOPATH/src/ -I . accounts.proto --lile-server_out=. --go_out=plugins=grpc:$$GOPATH/src

test:
	go test -v ./... -cover

get:
	go get -u -t ./...

ci: get test

docker:
	docker build . -t lileio/accounts:latest
