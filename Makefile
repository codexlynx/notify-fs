build: dist/notify-fs_amd64

gofmt:
	@gofmt -s -w .

clean:
	@rm dist/*

dist/notify-fs_amd64:
	@DOCKER_BUILDKIT=1 docker build -f build/binary.dockerfile --target binary --output dist/ .
#
all: build
.PHONY: clean gofmt
