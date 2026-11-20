// Copyright 2026 The Agent Substrate Authors
// SPDX-License-Identifier: Apache-2.0

// Package objectstore: Content-addressable snapshot transfer client supporting delta memory and filesystem layers.
package objectstore

type SandboxRuntime string

const (
	RuntimeGVisor  SandboxRuntime = "gvisor"
	RuntimeMicroVM SandboxRuntime = "microvm"
)
