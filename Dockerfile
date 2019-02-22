FROM golang:latest as builder
RUN mkdir /app
COPY . /go/src/mapserver/
WORKDIR /go/src/mapserver/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM alpine:latest
COPY --from=builder /go/src/mapserver/main /app/
WORKDIR /app
CMD ["./main"]

