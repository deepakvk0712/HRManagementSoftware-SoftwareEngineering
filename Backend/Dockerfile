FROM golang:1.17 AS builder

WORKDIR /src
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o main -a -ldflags '-linkmode external -extldflags "-static"' .

WORKDIR /dist

RUN cp /src/main .
RUN cp /src/Database/TestHR .
RUN cp /src/secrets.env .

FROM scratch
COPY --from=builder /dist /
EXPOSE 8080

ENTRYPOINT ["/main"]