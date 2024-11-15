FROM golang:1.23.3-alpine3.20
ARG UID
ARG GID

RUN apk update && apk upgrade && apk add --no-cache make

RUN mkdir /workspace
RUN chown $UID:$GID /workspace

RUN mkdir /.cache
RUN chown $UID:$GID /.cache
USER $UID:$GID

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.0

WORKDIR /workspace
