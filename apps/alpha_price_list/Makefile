APP_NAME=dbacAlphaPriceList
GIT_COMMIT=$(shell git rev-list -1 HEAD)
DATE=$(shell date)

build: clean
	@echo "Building Mac and Windows executables..."
	GOOS=darwin GOARCH=amd64 go build \
		-o ${APP_NAME} \
		-ldflags "-X main.GitCommit=${GIT_COMMIT} -X 'main.BuildDate=${DATE}'"
	GOOS=windows GOARCH=amd64 go build \
		-o ${APP_NAME}.exe \
		-ldflags "-X main.GitCommit=${GIT_COMMIT} -X 'main.BuildDate=${DATE}'"

package: build
	@echo "Packaging executables into zip files..."
	zip -m ${APP_NAME}.Mac.zip ${APP_NAME}
	zip -m ${APP_NAME}.Windows.zip ${APP_NAME}.exe

clean:
	@if [ -f "${APP_NAME}" ]; then \
		echo "Removing existing executables..."; \
		rm -f ${APP_NAME}*; \
	fi
