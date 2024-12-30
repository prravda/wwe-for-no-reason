# build client.go
GOOS=linux GOARCH=amd64 go build -o client client.go

# build tcp_go/tcp_server.go
GOOS=linux GOARCH=amd64 go build -o ./tcp_go/tcp_server ./tcp_go/tcp_server.go