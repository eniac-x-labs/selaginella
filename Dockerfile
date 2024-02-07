FROM --platform=$BUILDPLATFORM golang:1.21.1-alpine3.18 as builder

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
#ENV SELAGINELLA_CONFIG="./selaginella.yaml"
#ENV SELAGINELLA_MIGRATIONS_DIR="./migrations"
ENV SELAGINELLA_ENABLE_HSM=false
ENV SELAGINELLA_HSM_API_NAME=""
ENV SELAGINELLA_HSM_ADDRESS=""
ENV SELAGINELLA_HSM_CREDEN=""
ENV SELAGINELLA_PRIVATE_KEY="0f0b59bddf091da85ebee2d547b8e8c2d2a92fa23982bc54fe13d6e439b5f4e8"
ENV SELAGINELLA_L1_CHAIN_ID=11155111
ENV SELAGINELLA_L1_TRANSFER_MULTIPLE=3
ENV SELAGINELLA_L2_TRANSFER_MULTIPLE=3

CMD ["selaginella", "grpc", "--config", "/app/selaginella/selaginella.yaml"]
