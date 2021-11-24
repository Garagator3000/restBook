FROM golang:1.17.3 AS builder
LABEL maintainer="al.suhodubenko@yandex.ru"

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o restBook ./cmd/main.go


FROM alpine:latest

COPY --from=builder /app/configs ./configs
COPY --from=builder /app/restBook .
COPY --from=builder /app/wait-for-postgres.sh .
COPY --from=builder /app/db ./db

RUN apk update \
 && apk add --no-cache postgresql-client \
 && apk add --no-cache bash

RUN chmod +x wait-for-postgres.sh

EXPOSE 8000
CMD ["./restBook"]
