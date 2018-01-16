FROM golang:1.9.2-alpine3.7
LABEL AUTHOR="Valentin Stoican"
LABEL EMAIL="vali@seqvence.com"

RUN mkdir /app
ADD . /app

WORKDIR /app

RUN apk add --update --no-cache git

# Download dependencies
RUN go get -d ./...
# Compile
RUN go build -o amqp-go-mongodb .
# Run
CMD [ "/app/amqp-go-mongodb" ]