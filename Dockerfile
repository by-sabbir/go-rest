FROM golang:1.18 AS builder

RUN mkdir /app

ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

FROM alpine:latest AS prod
COPY --from=builder /app/app .

CMD [ "./app" ]