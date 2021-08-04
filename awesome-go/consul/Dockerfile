FROM golang:1.16-alpine
WORKDIR /app
COPY ./ .
RUN go get
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o web-svr main.go
EXPOSE 30003
CMD ["./web-svr"]