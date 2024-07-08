ARG GO_VERSION=1.23rc1
ARG ALPINE_VERSION=3.20

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION}

WORKDIR /go/src/kotatuneko-backend

RUN apk update && apk add git
RUN go install github.com/air-verse/air@latest

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

CMD ["air", "-c", ".air.toml"]