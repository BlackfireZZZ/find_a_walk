FROM golang:1.22 as build
LABEL authors="Blackfire"

WORKDIR /go/src/app
COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app backend/cmd/main.go
EXPOSE 8010

FROM gcr.io/distroless/static-debian12
COPY --from=build /go/bin/app /
CMD ["/app"]
