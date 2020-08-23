FROM golang:1.14 as build
WORKDIR /kodix
COPY ./ ./
RUN go build cmd/main.go
EXPOSE 8080
CMD ["/kodix/main"]