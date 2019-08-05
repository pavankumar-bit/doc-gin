FROM golang:1.12.0-alpine3.9
RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*  
RUN go get -u github.com/gin-gonic/gin

RUN mkdir -p /app

ADD . /app
WORKDIR /app
RUN go build -o ./app main.go

EXPOSE 8088

ENTRYPOINT ["./app"]
