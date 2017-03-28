FROM golang:1.8
MAINTAINER Neo <me@nex.tw>

EXPOSE 53/udp
EXPOSE 53

COPY . /go/src/proton
WORKDIR /go/src/proton

RUN godep restore
RUN go build
RUN ls -al

CMD ["./proton"]
