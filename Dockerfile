ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

RUN apt-get update && apt-get install -y sqlite3 libsqlite3-dev
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN go build -v -o /run-app .

COPY data/blog.db /usr/local/bin/blog.db

FROM debian:bookworm

RUN mkdir -p /data
COPY --from=builder /run-app /usr/local/bin/
COPY --from=builder /usr/local/bin/blog.db /usr/local/bin/blog.db
COPY entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/entrypoint.sh
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["run-app"]
