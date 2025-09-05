APP_NAME=webapp-bridge

build:
	GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME)-linux main.go
	GOOS=windows GOARCH=amd64 go build -o bin/$(APP_NAME)-windows.exe main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/$(APP_NAME)-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/$(APP_NAME)-darwin-arm64 main.go

clear:
	rm -rf bin