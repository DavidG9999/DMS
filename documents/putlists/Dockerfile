FROM golang:alpine  AS builder

COPY . /github.com/DavidG9999/DMS/documents/putlists

WORKDIR /github.com/DavidG9999/DMS/documents/putlists

RUN go mod download
RUN go build -o ./bin/putlists cmd/main.go

FROM alpine

WORKDIR /root/

COPY --from=builder /github.com/DavidG9999/DMS/documents/putlists/bin/putlists .
COPY --from=builder /github.com/DavidG9999/DMS/documents/putlists/configs configs/
COPY --from=builder /github.com/DavidG9999/DMS/documents/putlists/.env .

EXPOSE 4042

CMD [ "./putlists" ]