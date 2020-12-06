FROM golang:1.15

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o derp-gateway .

ENTRYPOINT ["/app/derp-gateway"]
