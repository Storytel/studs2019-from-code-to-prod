FROM golang:1.11.1-alpine3.8 as builder
WORKDIR /Users/johan.lejdung/localdev/go/studs-2019
COPY ./ .
RUN apk add build-base
RUN GOOS=linux GO111MODULE=on GOARCH=amd64 go build -mod=vendor -ldflags="-w -s" -v
RUN cp studs-2019 /

FROM alpine:3.8
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /studs-2019 /
CMD ["/studs-2019"]
