VERSION = 0.0.1

updatev:
		git tag v${VERSION} && git push origin v${VERSION}


fix:
	go mod tidy

updatemod:
	go get -u ./...


test:
	go test -v ./...



example:
	go run example/main.go

