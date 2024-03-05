##@ Development

.PHONY: update-swagger
update-swagger: swagger
	@echo "Updating swagger"
	$(SWAGGER) generate spec --scan-models -o docs/spec.json
	sed -i 's/"uint64"/"int64"/g' docs/spec.json

##@ Tools

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
SWAGGER ?= $(LOCALBIN)/swagger

## Tool Versions
SWAGGER_VERSION ?= v0.30.5

.PHONY: swagger
swagger: $(SWAGGER)
$(SWAGGER): | $(LOCALBIN)
	@echo "Installing swagger"
	test -s $(SWAGGER) || GOBIN=$(LOCALBIN) go install github.com/go-swagger/go-swagger/cmd/swagger@$(SWAGGER_VERSION)
