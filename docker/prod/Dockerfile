ARG GO_VERSION=1.23rc1
ARG ALPINE_VERSION=3.20

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as go-builder

WORKDIR /go/src/kotatuneko-backend

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o api cmd/main.go


FROM alpine:${ALPINE_VERSION}

WORKDIR /usr/src/kotatuneko-backend

COPY --from=go-builder /go/src/kotatuneko-backend/api api
COPY --from=go-builder /go/src/kotatuneko-backend/api .env
RUN chmod +x "/usr/src/kotatuneko-backend/api"

ENTRYPOINT ["/usr/src/kotatuneko-backend/api"]