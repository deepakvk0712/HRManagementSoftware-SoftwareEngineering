FROM golang:1.17 AS builder

WORKDIR /src
COPY . .
RUN go mod download
#RUN apt update && apt add ca-certificates && rm -rf/var/cache/apt/*
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true
RUN CGO_ENABLED=1 GOOS=linux go build -o main -a -ldflags '-linkmode external -extldflags "-static"' .

WORKDIR /dist

RUN cp /src/main .
RUN cp /src/Database/TestHR .
RUN cp /src/secrets.env .

FROM scratch
COPY --from=builder /dist /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080

ENTRYPOINT ["/main"]