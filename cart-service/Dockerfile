


FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY .. ./

RUN go build -o application


FROM alpine:3.17.2

COPY --from=builder /app/application /application


ENTRYPOINT ["./application"]