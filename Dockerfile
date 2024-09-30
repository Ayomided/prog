ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

RUN apt-get update && apt-get install -y sqlite3 libsqlite3-dev
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app .

FROM debian:bookworm
COPY --from=builder templates/ /usr/local/bin/
COPY --from=builder posts/ /usr/local/bin/
COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]
