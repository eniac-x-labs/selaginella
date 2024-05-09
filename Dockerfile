FROM golang:1.21.1-alpine3.18 as builder

RUN apk add --no-cache make ca-certificates gcc musl-dev linux-headers git jq bash

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

WORKDIR /app

RUN go mod download

# build selaginella with the shared go.mod & go.sum files
COPY . /app/selaginella

WORKDIR /app/selaginella

RUN make build

FROM alpine:3.18

COPY --from=builder /app/selaginella/selaginella /usr/local/bin
COPY --from=builder /app/selaginella/selaginella.yaml /app/selaginella/selaginella.yaml
COPY --from=builder /app/selaginella/migrations /app/selaginella/migrations

WORKDIR /app
ENV SELAGINELLA_MIGRATIONS_DIR="/app/selaginella/migrations"
ENV SELAGINELLA_CONFIG="/app/selaginella/selaginella.yaml"

CMD ["selaginella", "grpc", "--config", "/app/selaginella/selaginella.yaml"]
