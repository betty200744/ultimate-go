FROM golang:1.13 as mygo

WORKDIR /Users/betty/project/whale_go/src/gobyexample/docker
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app main.go

CMD "ENVIRON=localhost go test -json ./..."
