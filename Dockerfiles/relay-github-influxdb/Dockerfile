FROM golang:1.20-alpine3.18

ENV ROOT=/go/src/relay-webhook-influxdb
WORKDIR ${ROOT}

#COPY ./src/input-github-webhook ./input-github-webhook

RUN apk update &&  apk add git bash
RUN go install github.com/cosmtrek/air@v1.29.0

CMD ["air", "-c", ".air.toml"]