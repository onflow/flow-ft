.PHONY: test
test:
	$(MAKE) generate -C lib/go
	$(MAKE) test -C lib/go
	flow test --cover tests/*.cdc

.PHONY: ci
ci:
	$(MAKE) ci -C lib/go
	$(MAKE) ci -C lib/js/test
	flow test --cover tests/*.cdc
