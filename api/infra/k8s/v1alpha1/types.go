// Copyright 2026 The Agent Substrate Authors
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

type WorkerPoolSpec struct {
	Replicas       int32  `json:"replicas"`
	SandboxRuntime string `json:"sandboxRuntime"` // "gvisor" or "microvm"
	WarmPoolBuffer int32  `json:"warmPoolBuffer"`
	Atespace       string `json:"atespace"`
}

type ActorTemplateSpec struct {
	Image          string   `json:"image"`
	WorkerPoolRef  string   `json:"workerPoolRef"`
	EgressPolicy   string   `json:"egressPolicy"`
	AllowedSecrets []string `json:"allowedSecrets,omitempty"`
}
