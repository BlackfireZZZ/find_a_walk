FROM golang:1.22 as build
LABEL authors="Blackfire"

WORKDIR /go/src/app
COPY backend/go.mod go.mod
COPY backend/go.sum go.sum

RUN go mod download
COPY backend/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app cmd/project/main.go
EXPOSE 8010

FROM gcr.io/distroless/static-debian12
COPY --from=build /go/bin/app /
COPY backend/.env /
COPY backend/migrations /migrations
CMD ["/app"]