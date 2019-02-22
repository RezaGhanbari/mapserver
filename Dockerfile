FROM golang:latest as builder
ENV	 REDIS_HOST=redis \
        REDIS_PORT=6379
RUN mkdir /app
COPY . /go/src/mapserver/
WORKDIR /go/src/mapserver/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM alpine:latest
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true

COPY --from=builder /go/src/mapserver/main /app/
WORKDIR /app
CMD ["./main"]

