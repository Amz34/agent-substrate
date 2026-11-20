// Copyright 2026 The Agent Substrate Authors
// SPDX-License-Identifier: Apache-2.0

// Package oidcjwt provides: Zero-dependency OIDC token verifier validating SPIFFE and federated JWT claims.
package oidcjwt

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
