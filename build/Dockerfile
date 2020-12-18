FROM golang:1.15-alpine as builder

RUN apk add --no-cache build-base

WORKDIR /src
COPY ./go.mod ./go.mod
RUN go mod download
ARG VERSION
RUN test -n "$VERSION"
COPY . .
RUN CGO_ENABLED=1 go build -ldflags "-X main.versionString=${VERSION}" -o serverhd .


FROM alpine

LABEL maintainer="Will Fantom <willfantom@gmail.com>"

WORKDIR /app
COPY --from=builder /src/serverhd ./serverhd

ENTRYPOINT ["./serverhd"]
