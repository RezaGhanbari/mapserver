FROM golang:1.10 AS builder
ENV	API_TOKEN=WP5D&s3ftd^NU3TG@JH2n?!!@!MLmquD5t?V7vCPdANyY4Vrq5F \
        MAP_NAME=GOOGLE \
        REDIS_HOST=redis \
        REDIS_PORT=6379 \
        PORT=3001 \
        LIMIT=100 \
        DB=http://172.16.17.66:9200

ARG https_proxy=http://proxy.carpino.info:33080
ARG http_proxy=http://proxy.carpino.info:33080

RUN go get github.com/cnjack/throttle
RUN go get github.com/gin-gonic/gin
RUN go get github.com/gin-contrib/cache
RUN go get github.com/satori/go.uuid
RUN go get github.com/gomodule/redigo/redis
RUN go get github.com/olivere/elastic
RUN go get github.com/JamesMilnerUK/pip-go
RUN go get github.com/corpix/uarand
RUN go get github.com/kelseyhightower/envconfig

COPY . /go/src/geo/
WORKDIR /go/src/geo/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .


FROM phusion/baseimage
COPY --from=builder /go/src/geo/main /app/
WORKDIR /app
CMD ["./main"]

