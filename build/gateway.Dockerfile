FROM golang:1.21 AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o gateway ./cmd/gateway

WORKDIR /dist
RUN cp /build/gateway .

FROM alpine

COPY --chown=0:0 --from=builder /dist /

USER 65534
CMD ["/gateway"]