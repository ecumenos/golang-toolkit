SHELL=/bin/sh

.PHONY: gomodcheck
gomodcheck:
	! grep "replace.*=>.*\.\.\\.*" go.mod

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

.PHONY: format
format:
	@GO111MODULE=on go fmt ./...

# Configure git hooks
.PHONY: hooks
hooks:
	git config core.hooksPath hooks

.PHONY: lint
lint:
	golangci-lint run -v --fix -c .golangci.yaml ./...

.PHONY: lint-ci
lint_ci: lint_deps
	golangci-lint run -v -E revive --print-issued-lines=false --timeout 2m

.PHONE: lint_deps
lint_deps:
	wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.30.0

GOTEST=GO111MODULE=on go test
GOTESTOPTS=-count=1 -timeout 5m
WITHCOVERAGE=-cover -coverprofile=/tmp/identification.cover -covermode count

.PHONY: build
build:
	GO111MODULE=on go build ./...

.PHONY: test_deps
test_deps:
	CGO_ENABLED=0 GO111MODULE=off go get -u -v github.com/jstemmer/go-junit-report

.PHONY: test_generate
test_generate:
	go generate ./test/mocks/interf/...

# test_compile just builds the tests without running them, useful for
# generating minimal output if one test fails to build, because it
# won't run all the other tests
.PHONY: test_compile
test_compile:
	${GOTEST} ./... ${GOTESTOPTS} -run xxxxx

.PHONY: test
test:
	${GOTEST} ./... ${GOTESTOPTS} ${WITHCOVERAGE}

test_integration:
	${GOTEST} ./... -tags=integration ${GOTESTOPTS} -run Integration

precommit: hooks vendor test_generate format lint build test test_integration gomodcheck
