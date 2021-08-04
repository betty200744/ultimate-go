FROM golang:1.16-alpine
WORKDIR /app
COPY ./ .
RUN go get
RUN go build -o /docker-go
EXPOSE 8080
CMD ["docker-go"]