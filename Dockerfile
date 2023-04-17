FROM golang:1.19.6-alpine3.17 AS build

WORKDIR /go/src/account_balance
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/account_balance cmd/account_balance.go

FROM alpine:3.17

ENV GOPROXY=https://proxy.golang.org
ENV GIN_MODE=release

RUN mkdir -p /csv_files/pending && \
    mkdir /csv_files/processed && \
    mkdir /csv_files/error

COPY --from=build /go/bin/account_balance /app
COPY --from=build /go/src/account_balance/resources* /resources
COPY --from=build /go/src/account_balance/resources/csv/example.csv /csv_files/pending/example.csv

ENTRYPOINT ["/app"]
