# Use the official Golang image to build the migrations binary
FROM golang:1.22 AS builder

WORKDIR /app

COPY migrations/go.mod migrations/go.sum ./

RUN go mod download

COPY migrations/ .

RUN CGO_ENABLED=0 GOOS=linux go build -o /migrations

FROM gcr.io/distroless/static-debian12

COPY --from=builder /migrations /migrations

COPY migrations/sql_migrations /sql_migrations

ENTRYPOINT ["/migrations"]
