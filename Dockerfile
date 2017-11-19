FROM golang:1.9 as builder
WORKDIR /go/src/github.com/smarthut/smarthut
COPY . .
RUN make vendor
RUN make build

FROM alpine:3.6
RUN apk add -U tzdata
COPY --from=builder /go/src/github.com/smarthut/smarthut/smarthut /
EXPOSE 8080
VOLUME ["/data"]
ENTRYPOINT ["/smarthut"]
