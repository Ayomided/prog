ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm AS builder

RUN apt-get update && apt-get install -y sqlite3 libsqlite3-dev
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
COPY posts/ /usr/src/app/posts/
COPY templates/ /usr/src/app/templates/
RUN go build -v -o /run-app .

FROM debian:bookworm
COPY --from=builder /run-app /usr/local/bin/
COPY --from=builder /usr/src/app/templates/ /usr/local/share/app/templates/
COPY --from=builder /usr/src/app/posts/ /usr/local/share/app/posts/
CMD ["run-app"]
