#!/bin/bash
set -e

kind create cluster --config kind-config.yaml
kubectl cluster-info --context kind-kind

