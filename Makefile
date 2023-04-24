APP_NAME=thoth-notes
APP_VERSION=v0.9.0


.PHONEY: run dev prod clean
build: $(shell find frontend -type f) $(shell find backend -type f)
	@CGO_ENABLED=1 wails build --debug

run: build
	@./build/bin/${APP_NAME}.exe

dev:
	@CGO_ENABLED=1 wails dev

dev.quick:
	@CGO_ENABLED=1 wails dev -skipbindings -nosyncgomod

test:
	@go test ./...

scan:
	@golangci-lint run
	@gosec ./...

prod:
	@CGO_ENABLED=1 wails build -o ${APP_NAME}-${APP_VERSION}.exe -upx -race -nsis

clean:
	rm -f ./build/bin/${APP_NAME}-${APP_VERSION}*
