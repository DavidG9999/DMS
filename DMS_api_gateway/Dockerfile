FROM golang:alpine  AS builder

COPY . /github.com/DavidG9999/DMS/DMS_api_gateway

WORKDIR /github.com/DavidG9999/DMS/DMS_api_gateway

RUN go mod download
RUN go build -o ./bin/dms_api_gateway cmd/main.go

FROM alpine

WORKDIR /root/

COPY --from=builder /github.com/DavidG9999/DMS/DMS_api_gateway/bin/dms_api_gateway .
COPY --from=builder /github.com/DavidG9999/DMS/DMS_api_gateway/configs configs/
COPY --from=builder /github.com/DavidG9999/DMS/DMS_api_gateway/internal/logger/logfiles var/log

EXPOSE 8080

CMD [ "./dms_api_gateway" ]