FROM golang:1.22.3-alpine AS build_base


WORKDIR /tmp/bot

COPY . .

RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go mod download
#RUN go build -o ./out/bot main.go
RUN go build -gcflags all='-N -l' -o ./out/bot main.go

# Start fresh from a smaller image
FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=build_base /go/bin/dlv /opt/dlv
COPY --from=build_base /tmp/bot/out/bot /app/bot

# This container exposes port 8080 to the outside world
EXPOSE 8081
EXPOSE 3344

# Run the binary program produced by `go install`
#CMD ["/app/bot"]
CMD ["/opt/dlv", "--headless", "--listen", ":3344", "--continue", "--accept-multiclient", "exec", "./app/bot"]