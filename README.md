# ccm-conformance-tests
To have interface for ccm testing across cloud-providers
This is still WIP - ( Makefile need some changes , build can be done successfully , but need to delete build folder for new build , will update that)
Prerequiste - kind cluster or any other cluster which have nodes running - More details [here](https://kind.sigs.k8s.io/docs/user/quick-start/#installation) .

# After 'make build' succeeds , run below- 
./build/e2e.test --ginkgo.focus 'When a new node is present' --kubeconfig=<kubeconfig path>
Running Suite: E2E Suite - ~go/src/github.com/ccm-conformance-tests
================================================================================ # 
Random Seed: 1747385408

Will run 1 of 1 specs
•

Ran 1 of 1 Specs in 0.023 seconds
SUCCESS! -- 1 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS' 
