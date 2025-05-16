GINKGO_FLAGS ?= -v --junit-report=report.xml
KUBECONFIG ?= $(HOME)/.kube/config

build:
	go mod tidy
	cd tests/e2e && go test -c -o ../../build/e2e.test

test-e2e: build
	./build/e2e.test $(GINKGO_FLAGS) --kubeconfig=$(KUBECONFIG)

lint:
	go fmt ./...
	go vet ./...

