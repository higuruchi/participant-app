FROM "golang:1.17-alpine3.14" AS builder

COPY ./go.mod /src/go.mod
COPY ./go.sum /src/go.sum

WORKDIR /src

RUN apk  add --update git
RUN go mod download

COPY ./ /src/
RUN go build -o ./cmd/main ./cmd/main.go

# ---------------------------------

FROM "alpine:3.14"

COPY --from=builder /src/cmd/main /app/main
WORKDIR /app
RUN apk add --update

CMD ["/app/main"]