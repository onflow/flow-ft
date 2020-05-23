.PHONY: test
test:
	go test ./...

# TODO: enable unit tests once flow-emulator is public
.PHONY: ci
ci:
	$(MAKE) -C contracts ci
