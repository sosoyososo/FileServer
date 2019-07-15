rm ./FileServer*
rm ./*.zip

env GOOS=linux GOARCH=amd64 go build -o FileServer-linux
env GOOS=windows GOARCH=amd64 go build -o FileServer-windows.exe
env GOOS=darwin GOARCH=amd64 go build -o FileServer-mac

zip ./linux.zip ./FileServer-linux ./upload.html
zip ./mac.zip ./FileServer-mac ./upload.html
zip ./windows.zip ./FileServer-windows.exe ./upload.html

rm ./FileServer*