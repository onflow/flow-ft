.PHONY: test
test:
	go test ./...

.PHONY: ci
ci: test
	$(MAKE) -C contracts ci
