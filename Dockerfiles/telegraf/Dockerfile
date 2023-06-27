FROM golang:1.20-alpine3.18 AS builder

ENV ROOT=/go/src/
WORKDIR ${ROOT}

COPY ./src/input-github-webhook ./input-github-webhook

RUN ls -la
RUN apk update && apk add git
RUN go build -o /go/dist/input-github-webhook ./input-github-webhook/cmd/main.go

FROM telegraf

WORKDIR /etc/telegraf

COPY --from=builder /go/dist/input-github-webhook .

RUN ls -la

#ENTRYPOINT ["/usr/local/etc/input-github-webhook"]