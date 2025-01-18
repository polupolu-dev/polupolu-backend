ARG GO_VERSION=1.23
ARG ALPINE_VERSION=3.21

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as go-builder

WORKDIR /go/src/polupolu-backend

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o cmd/main cmd/main.go


FROM alpine:${ALPINE_VERSION}

WORKDIR /
COPY --from=go-builder /go/src/polupolu-backend/cmd/main /api
COPY --from=go-builder /go/src/polupolu-backend/.env /.env
EXPOSE 8080
RUN chmod +x "/api"

ENTRYPOINT ["/api"]
