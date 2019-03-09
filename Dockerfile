FROM golang:1.12.0-alpine3.9
WORKDIR /go/src/github.com/kofoworola/definethephrase

#ENV
ENV HANDLE="@__define__"

#commands and entry point
RUN go build
ENTRYPOINT ./definethephrase