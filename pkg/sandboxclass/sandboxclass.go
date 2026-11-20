// Copyright 2026 The Agent Substrate Authors
// SPDX-License-Identifier: Apache-2.0

// Package sandboxclass: Runtime specification definitions for gVisor (runsc) and Cloud Hypervisor microVM classes.
package sandboxclass

type SandboxRuntime string

const (
	RuntimeGVisor  SandboxRuntime = "gvisor"
	RuntimeMicroVM SandboxRuntime = "microvm"
)
