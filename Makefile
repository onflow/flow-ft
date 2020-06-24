.PHONY: ci
ci:
	$(MAKE) -C test ci
	$(MAKE) -C templates ci
	$(MAKE) -C contracts ci
