FROM golang:1.16-alpine3.13 as builder

# prepare dependencies
RUN apk add git openssh-client
RUN apk add --update alpine-sdk

# prepage go mod permissions

COPY . /src
WORKDIR /src

RUN go build -o /backend

FROM alpine:latest
COPY --from=builder /backend /backend

CMD ["/backend"]