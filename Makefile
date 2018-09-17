watcher/config_generated.go: conf/config.yml
	go generate

.PHONY: clean
clean:
	@rm -fv watcher/config_generated.go
	@rm -fv eztv-watcher
	@rm -fv db.bin

.PHONY: build
build: watcher/config_generated.go
	go build

.PHONY: bump
bump:
	$(eval VERSION=$(shell git describe --tags --abbrev=0 | awk -F. '{$$NF+=1; OFS="."; print $0}'))
	@git tag -a $(VERSION) -m "Bumping to $(VERSION)"

.PHONY: release
release:
	@goreleaser --rm-dist