FROM ubuntu:18.04 as build
RUN add-apt-repository ppa:longsleep/golang-backports; \
    apt-get update; \
    apt-get install golang-goapt update; \
    apt install libsnmp-dev
COPY . /snmp_exporter
RUN cd /snmp_exporter; \
    make

FROM        quay.io/prometheus/busybox:latest
LABEL       maintainer="The Prometheus Authors <prometheus-developers@googlegroups.com>"

COPY --from=gatherer /snmp_exporter/snmp_exporter  /bin/snmp_exporter
COPY --from=gatherer /snmp_exporter/snmp.yml       /etc/snmp_exporter/snmp.yml

EXPOSE      9116
ENTRYPOINT  [ "/bin/snmp_exporter" ]
CMD         [ "--config.file=/etc/snmp_exporter/snmp.yml" ]
