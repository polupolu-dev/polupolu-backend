ARG GO_VERSION=1.23
ARG ALPINE_VERSION=3.21

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION}

WORKDIR /go/src/polupolu-backend

RUN apk update && apk add git
RUN go install github.com/air-verse/air@latest

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

CMD ["air", "-log.silent", "true", "-c", ".air.toml"]
