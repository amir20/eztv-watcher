.PHONY: git, build, release

config_generate.go:
	go generate

clean:
	rm config_generated.go

build: config_generate.go
	go build

bump:
	$(eval VERSION=$(shell git describe --tags --abbrev=0 | awk -F. '{$$NF+=1; OFS="."; print $0}'))
	@git tag -a $(VERSION) -m "Bumping to $(VERSION)"

release:
	@goreleaser --rm-dist