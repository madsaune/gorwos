BUILD_DATE=$(shell date)

default:
	echo "Copyright (c) 2022 - Mads Moi-Aune"

release-local:
	goreleaser release --snapshot --skip-publish --rm-dist

release:
	goreleaser release

build:
	goreleaser build --rm-dist --single-target

snapshot:
	goreleaser release --snapshot --skip-publish --rm-dist

lint:
	golangci-lint run
