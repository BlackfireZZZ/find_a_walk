FROM golang:1.22 AS builder

WORKDIR /app

COPY migrations_module/go.mod migrations_module/go.sum ./

RUN go mod download

ADD migrations_module/migrations migartions/
COPY migrations_module/ .


RUN CGO_ENABLED=0 GOOS=linux go build -o /migrations_module

FROM gcr.io/distroless/static-debian12

COPY --from=builder /migrations_module /migrations_module

ADD migrations_module/migrations migrations/
COPY migrations_module/migrations migrations/


ENTRYPOINT ["/migrations_module"]