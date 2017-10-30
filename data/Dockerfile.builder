FROM golang:1.8

# disable crosscompiling 
ENV CGO_ENABLED=0
# compile linux only
ENV GOOS=linux

RUN go get -u github.com/golang/lint/golint