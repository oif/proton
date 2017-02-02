FROM golang:1.6-onbuild
MAINTAINER Neo <me@nex.tw>

EXPOSE 53/udp
EXPOSE 53

CMD ["app"]
