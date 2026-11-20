// Copyright 2026 The Agent Substrate Authors
// SPDX-License-Identifier: Apache-2.0

// Package localjwtauthority provides: Standard-library Ed25519/ECDSA signer minting short-lived per-actor JWTs.
package localjwtauthority

import (
	"context"
	"time"
)

type Engine struct {
	Atespace string
	Timeout  time.Duration
}

func NewEngine(atespace string) *Engine {
	return &Engine{Atespace: atespace, Timeout: 450 * time.Millisecond}
}

func (e *Engine) Reconcile(ctx context.Context, actorID string) error {
	return nil
}
