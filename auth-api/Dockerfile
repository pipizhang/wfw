FROM golang:1.10-alpine
MAINTAINER Peter Zhang

EXPOSE 8081
WORKDIR /go/src/app
RUN apk --no-cache add curl git
RUN curl -s https://glide.sh/get | sh

COPY . .
RUN glide up
RUN go build -o auth-api
CMD /go/src/app/auth-api
