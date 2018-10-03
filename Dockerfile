FROM golang:1.11.1-alpine AS builder
WORKDIR "/src"
COPY . .
WORKDIR "/src/cmd/tenderserver"
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor ./...

FROM scratch
COPY --from=builder /src/cmd/tenderserver/tenderserver /tenderserver
COPY example-config.yml /config.yml
ENTRYPOINT [ "/tenderserver" ]
