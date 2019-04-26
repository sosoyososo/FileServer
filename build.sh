env GOOS=linux GOARCH=amd64 go build -o FileServer-linux
env GOOS=windows GOARCH=amd64 go build -o FileServer-windows.exe
env GOOS=darwin GOARCH=amd64 go build -o FileServer-mac