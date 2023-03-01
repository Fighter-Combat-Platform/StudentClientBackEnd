set GOARCH=amd64
set GOOS=linux
go build main.go 
pscp main root@10.119.9.110:/root/wzy