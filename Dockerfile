FROM golang:1.9 as builder
WORKDIR /go/src/github.com/smarthut/smarthut
COPY . .
RUN make vendor
RUN make build

FROM alpine:latest
COPY --from=builder /go/src/github.com/smarthut/smarthut/smarthut /
EXPOSE 8080
ENTRYPOINT ["/smarthut"]
