FROM golang:1.22.3-alpine AS build_base

WORKDIR /tmp/bot

COPY . .

RUN go mod download

EXPOSE 8080

#RUN CGO_ENABLED=0 go test -v

RUN go build -o ./out/bot main.go

# Start fresh from a smaller image
FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=build_base /tmp/bot/out/bot /app/bot

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/app/bot"]