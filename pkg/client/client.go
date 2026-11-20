// Copyright 2026 The Agent Substrate Authors
// SPDX-License-Identifier: Apache-2.0

// Package client: Public Go gRPC and HTTP client for interacting with ate-api-server and atenet-router.
package client

type SandboxRuntime string

const (
	RuntimeGVisor  SandboxRuntime = "gvisor"
	RuntimeMicroVM SandboxRuntime = "microvm"
)
