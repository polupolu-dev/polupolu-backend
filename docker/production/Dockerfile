ARG GO_VERSION=1.23
ARG ALPINE_VERSION=3.21

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as go-builder

WORKDIR /go/src/polupolu-backend

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o api cmd/main.go


FROM alpine:${ALPINE_VERSION}

WORKDIR /usr/src/polupolu-backend

COPY --from=go-builder /go/src/polupolu-backend/api/v1 api
COPY --from=go-builder /go/src/polupolu-backend/api/v1 .env
RUN chmod +x "/usr/src/polupolu-backend/api/v1"

ENTRYPOINT ["/usr/src/polupolu-backend/api/v1"]
