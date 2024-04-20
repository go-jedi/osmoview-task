FROM golang:1.22.2-alpine AS builder

WORKDIR /github.com/go-jedi/osmoview-task/app/
COPY . /github.com/go-jedi/osmoview-task/app/

RUN go mod download
RUN go build -o .bin/task cmd/task/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/go-jedi/osmoview-task/app/.bin/task .
COPY .env /root/

CMD ["./task"]