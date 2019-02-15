GIT_COMMIT=$(shell git rev-list -1 HEAD)
DATE=$(shell date)

build:
	go build -o dbacAlphaPriceList \
		-ldflags "-X main.GitCommit=${GIT_COMMIT} -X 'main.BuildDate=${DATE}'"

