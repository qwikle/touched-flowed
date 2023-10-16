run:
	clear && go run .

test:
	go test -v ./tests/...

reset-db:
	sqitch revert -y ; sqitch deploy

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/windows/$(APP_NAME).exe .

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/$(APP_NAME) .

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/$(APP_NAME) .