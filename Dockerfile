FROM golang:1.9 as builder
WORKDIR /go/src/github.com/smarthut/smarthut
COPY . .
RUN make vendor ; make build

FROM alpine:3.6 as alpine
RUN apk --no-cache add tzdata zip ca-certificates
WORKDIR /usr/share/zoneinfo
RUN zip -r -0 /zoneinfo.zip .

FROM scratch
COPY --from=builder /go/src/github.com/smarthut/smarthut/smarthut /smarthut
ENV ZONEINFO /zoneinfo.zip
COPY --from=alpine /zoneinfo.zip /
EXPOSE 8080
VOLUME ["/data"]
ENTRYPOINT ["/smarthut"]
