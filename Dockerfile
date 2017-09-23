FROM scratch
COPY . /
ADD zoneinfo.tar.gz /
EXPOSE 80
ENTRYPOINT ["/smarthut"]
