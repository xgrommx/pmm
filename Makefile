default: help

help:                 ## Display this help message.
	@echo "Please use \`make <target>\` where <target> is one of:"
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | \
		awk -F ':.*?## ' 'NF==2 {printf "  %-26s%s\n", $$1, $$2}'

init:                 ## Install tools.
	rm -rf bin
	cd tools && go generate -x -tags=tools

	# Download third-party proto files
	$(eval GO_PROTO_VALIDATOR=$(shell go list -f '{{ .Version }}' -m github.com/mwitkow/go-proto-validators))
	curl --create-dirs -L https://raw.githubusercontent.com/mwitkow/go-proto-validators/$(GO_PROTO_VALIDATOR)/validator.proto -o ../third_party/github.com/mwitkow/go-proto-validators/validator.proto

release:             ## Build release versions of
	make -C agent release
	make -C admin release

gen: clean         ## Generate files.
#	# generated by descriptors target
	#./bin/prototool break check api -f api/api.descriptor
	bin/buf breaking --against descriptor.bin api

	bin/buf generate -v api

	for API in api/agentlocalpb api/serverpb api/inventorypb api/managementpb api/managementpb/dbaas api/managementpb/ia api/managementpb/backup api/managementpb/azure api/qanpb api/platformpb ; do \
		set -x ; \
		bin/swagger mixin $$API/json/header.json $$API/*.swagger.json --output=$$API/json/$$(basename $$API).json --keep-spec-order; \
		bin/swagger flatten --with-flatten=expand --with-flatten=remove-unused $$API/json/$$(basename $$API).json --output=$$API/json/$$(basename $$API).json ; \
		bin/swagger validate $$API/json/$$(basename $$API).json ; \
		bin/swagger generate client --with-flatten=expand --with-flatten=remove-unused --spec=$$API/json/$$(basename $$API).json --target=$$API/json \
			--additional-initialism=aws \
			--additional-initialism=db \
			--additional-initialism=ok \
			--additional-initialism=pmm \
			--additional-initialism=psmdb \
			--additional-initialism=pxc \
			--additional-initialism=pt \
			--additional-initialism=qan \
			--additional-initialism=rds \
			--additional-initialism=sql \
			--additional-initialism=ha ; \
	done

	# generate public API spec, omit agentlocalpb (always private),
	# and managementpb/dbaas, managementpb/ia, managementpb/backup , managementpb/azure and qanpb (not v1 yet)
	bin/swagger mixin --output=api/swagger/swagger.json \
		api/swagger/header.json \
		api/serverpb/json/serverpb.json \
		api/inventorypb/json/inventorypb.json \
		api/managementpb/json/managementpb.json
	bin/swagger validate api/swagger/swagger.json

	bin/swagger-order --output=api/swagger/swagger.json api/swagger/swagger.json

	# generate API spec with all PMM Server APIs (omit agentlocalpb)
	bin/swagger mixin --output=api/swagger/swagger-dev.json \
		api/swagger/header-dev.json \
		api/serverpb/json/serverpb.json \
		api/inventorypb/json/inventorypb.json \
		api/managementpb/json/managementpb.json \
		api/managementpb/dbaas/json/dbaas.json \
		api/managementpb/ia/json/ia.json \
		api/managementpb/backup/json/backup.json \
		api/managementpb/azure/json/azure.json \
		api/qanpb/json/qanpb.json \
		api/platformpb/json/platformpb.json
	bin/swagger validate api/swagger/swagger-dev.json

	bin/swagger-order --output=api/swagger/swagger-dev.json api/swagger/swagger-dev.json

	# generate API spec with only dev PMM Server APIs specifically for readme.io (omit agentlocalpb)
	bin/swagger mixin --output=api/swagger/swagger-dev-only.json \
		api/swagger/header-dev.json \
		api/managementpb/dbaas/json/dbaas.json \
		api/managementpb/ia/json/ia.json \
		api/managementpb/backup/json/backup.json \
		api/managementpb/azure/json/azure.json \
		api/qanpb/json/qanpb.json \
		api/platformpb/json/platformpb.json
	bin/swagger validate api/swagger/swagger-dev-only.json

	bin/swagger-order --output=api/swagger/swagger-dev-only.json api/swagger/swagger-dev-only.json

	make clean_swagger
	make -C agent gen
	make format
	make format ## TODO: One formatting run is not enough, figure out why.
	bin/go-sumtype ./...
	go install -v ./...

gen-alertmanager:     # Generate Alertmanager client.
	bin/swagger generate client --model-package=ammodels --client-package=amclient --spec=api/alertmanager/openapi.yaml --target=api/alertmanager

	make format
	go install -v ./api/alertmanager/...

clean_swagger:
	find api -name '*.swagger.json' -print -delete

clean: clean_swagger  ## Remove generated files.
	find api -name '*.pb.go' -print -delete
	find api -name '*.pb.gw.go' -print -delete

	for API in api/agentlocalpb api/serverpb api/inventorypb api/managementpb api/managementpb/dbaas api/managementpb/ia api/managementpb/backup api/qanpb api/platformpb ; do \
		rm -fr $$API/json/client $$API/json/models $$API/json/$$(basename $$API).json ; \
	done
	rm -f api/swagger/swagger.json api/swagger/swagger-dev.json api/swagger/swagger-dev-only.json

test:                 ## Run tests
	go test ./...

check:                          ## Run required checkers and linters.
	#go run .github/check-license.go ## TODO: This repo has multiple licenses, fix checker
	bin/golangci-lint run -c=.golangci-required.yml
	bin/go-consistent -pedantic ./...

check-all: check                ## Run golang ci linter to check new changes from main.
	bin/golangci-lint run -c=.golangci.yml --new-from-rev=main

FILES = $(shell find . -type f -name '*.go')

format:                         ## Format source code.
	bin/gofumpt -l -w $(FILES)
	bin/goimports -local github.com/percona/pmm-agent -l -w $(FILES)
	bin/gci write --Section Standard --Section Default --Section "Prefix(github.com/percona/pmm)" $(FILES)
	bin/goimports -local github.com/percona/pmm -l -w $(FILES) # Temporary fix, gci has bug with sorting black imports.

serve:                ## Serve API documentation with nginx.
	# http://127.0.0.1:8080/swagger-ui.html
	nginx -p . -c api/nginx/nginx.conf

descriptors:          ## Update API compatibility descriptors.
	#./prototool break descriptor-set . -o api/api.descriptor
	bin/buf build -o descriptor.bin --as-file-descriptor-set api

ci-reviewdog:         ## Runs reviewdog checks.
	# go run .github/check-license.go TODO: This repo has multiple licenses, fix checker
	bin/go-consistent -pedantic -exclude "tests" ./... | bin/reviewdog -f=go-consistent -name='Required go-consistent checks' -reporter=github-pr-review -fail-on-error
	bin/golangci-lint run -c=.golangci-required.yml --out-format=line-number | bin/reviewdog -f=golangci-lint -reporter=github-pr-review -fail-on-error
	bin/golangci-lint run -c=.golangci.yml --out-format=line-number | bin/reviewdog -f=golangci-lint -reporter=github-pr-review
