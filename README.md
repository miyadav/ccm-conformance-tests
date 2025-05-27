# ccm-conformance-tests

`ccm-conformance-tests` is a Kubernetes end-to-end (E2E) conformance testing suite focused on validating Cloud Controller Manager (CCM) implementations. It provides a framework for writing, organizing, and running tests using [Ginkgo](https://onsi.github.io/ginkgo/) and [kind](https://kind.sigs.k8s.io/) clusters.

## Features

* Automated setup and teardown of test namespaces
* Kind-based Kubernetes cluster creation using config files
* Ginkgo-based E2E test runner
* GitHub Actions workflow support for CI/CD integration

## Prerequisites

* Go (>= 1.20)
* Docker
* [kind](https://kind.sigs.k8s.io/) installed and accessible in your path
* kubectl
* Ginkgo CLI: `go install github.com/onsi/ginkgo/v2/ginkgo@latest`

## Project Structure

```
ccm-conformance-tests/
├── cmd/ccm-test/          # Entry point for test binary
├── examples/kind/         # Sample kind cluster configuration
├── tests/e2e/             # Ginkgo-based test suites
│   ├── e2e_suite_test.go  # Ginkgo setup
│   └── framework/         # Framework for client and namespace setup
├── .github/workflows/     # CI configuration
├── go.mod / go.sum        # Go module definitions
├── Makefile               # Build and test shortcuts
```

## Setup

1. Clone the repository:

```bash
git clone https://github.com/miyadav/ccm-conformance-tests.git
cd ccm-conformance-tests
```

2. Create a kind cluster:

```bash
kind create cluster \
  --config=examples/kind/kind-config.yaml \
  --name test-cluster \
  --kubeconfig=$HOME/.kube/config
```

3. Run the tests:

```bash
make test-e2e
```

4. Tear down the kind cluster:

```bash
kind delete cluster --name test-cluster
```

## Make Targets

* `make test-e2e` - Run the E2E test suite using Ginkgo

## CI Integration

The `.github/workflows/ci.yaml` workflow handles:

* Go setup
* Kind cluster provisioning
* Running Ginkgo E2E tests
* Cleaning up the cluster post-test

## Contributing

Pull requests are welcome! Feel free to fork and submit changes to expand coverage or improve the framework.

---
