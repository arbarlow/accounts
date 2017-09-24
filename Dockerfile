FROM golang:1.9-alpine AS build-env
WORKDIR /go/src/github.com/lileio/accounts
COPY . /go/src/github.com/lileio/accounts
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build -o build/accounts ./accounts
RUN go build -o build/accounts_cli ./accounts_cli


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /go/src/github.com/lileio/accounts/build/accounts /bin/accounts
COPY --from=build-env /go/src/github.com/lileio/accounts/build/accounts_cli /bin/accounts_cli
CMD ["accounts", "up"]
