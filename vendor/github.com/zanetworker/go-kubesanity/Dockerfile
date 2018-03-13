FROM golang:1.9 AS build-env
WORKDIR /go/src/github.com/zanetworker/go-kubesanity/
ADD . .
RUN make OS=linux dry


FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/zanetworker/go-kubesanity/kubesanity /app/
ENTRYPOINT [ "./kubesanity" ]

