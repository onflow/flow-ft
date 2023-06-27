.PHONY: test
test:
	flow test --cover tests/*.cdc

.PHONY: ci
ci:
	flow test --cover tests/*.cdc
