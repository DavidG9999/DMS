FROM golang:alpine  AS builder

COPY . /github.com/DavidG9999/DMS/users

WORKDIR /github.com/DavidG9999/DMS/users

RUN go mod download
RUN go build -o ./bin/users cmd/main.go

FROM alpine

WORKDIR /root/

COPY --from=builder /github.com/DavidG9999/DMS/users/bin/users .
COPY --from=builder /github.com/DavidG9999/DMS/users/configs configs/
COPY --from=builder /github.com/DavidG9999/DMS/users/.env .

EXPOSE 4041

CMD [ "./users" ]