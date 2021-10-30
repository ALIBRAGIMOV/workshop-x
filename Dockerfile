FROM golang:1.17-alpine

WORKDIR /Users/ali/project_x/workshop-x
ADD ./ /Users/ali/project_x/workshop-x

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN go mod download

RUN go build -o bin/workshop cmd/workshop/main.go
ENTRYPOINT ["bin/workshop"]