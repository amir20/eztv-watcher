.PHONY: git, build, release

clean: 
	rm config_generated.go

build:
	go build

bump:
	$(eval VERSION=$(shell git describe --tags --abbrev=0 | awk -F. '{$$NF+=1; OFS="."; print $0}'))
	@git tag -a $(VERSION) -m "Bumping to $(VERSION)"

release:
	@goreleaser --rm-dist