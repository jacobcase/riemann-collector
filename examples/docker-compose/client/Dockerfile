FROM ubuntu:14.04

COPY riemann-collector /usr/local/bin/

RUN mkdir /usr/local/etc/riemann-collector

COPY ./conf.yaml ./conf.d ./scripts /usr/local/etc/riemann-collector/

