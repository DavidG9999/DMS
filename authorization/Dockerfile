FROM golang:alpine  AS builder

COPY . /github.com/DavidG9999/DMS/authorization

WORKDIR /github.com/DavidG9999/DMS/authorization

RUN go mod download
RUN go build -o ./bin/auth cmd/main.go

FROM alpine

WORKDIR /root/

COPY --from=builder /github.com/DavidG9999/DMS/authorization/bin/auth .
COPY --from=builder /github.com/DavidG9999/DMS/authorization/configs configs/
COPY --from=builder /github.com/DavidG9999/DMS/authorization/.env .
COPY --from=builder /github.com/DavidG9999/DMS/authorization/internal/logger/logfiles var/log


EXPOSE 4040

CMD [ "./auth" ]