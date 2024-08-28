FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o crm.shopdev.com ./cmd/server/main.go

FROM scratch

COPY ./config /config

COPY --from=builder /build/crm.shopdev.com /

ENTRYPOINT [ "/crm.shopdev.com", "/config/config.yaml"]

