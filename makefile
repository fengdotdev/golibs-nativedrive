VERSION = 0.0.1

updatev:
		git tag v${VERSION} && git push origin v${VERSION}



sand:
	go run -v ./cmd/sandbox/main.go