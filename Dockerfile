FROM golang:1.19.6-alpine3.17 AS build

WORKDIR /go/src/account_balance
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/account_balance cmd/account_balance.go

FROM scratch

ENV GOPROXY=https://proxy.golang.org
ENV GIN_MODE=release

COPY --from=build /go/bin/account_balance /app

ENTRYPOINT ["/app"]
