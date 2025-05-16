GINKGO_FLAGS ?= -v --junit-report=report.xml
KUBECONFIG ?= $(HOME)/.kube/config

test-e2e:
	cd tests/e2e && ginkgo $(GINKGO_FLAGS) -- --kubeconfig=$(KUBECONFIG)

lint:
	go fmt ./...
	go vet ./...


