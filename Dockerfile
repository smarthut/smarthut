FROM scratch
LABEL maintainer="devs@smarthut.cc"
COPY . /
ADD zoneinfo.tar.gz /
EXPOSE 80
ENTRYPOINT ["/smarthut"]
