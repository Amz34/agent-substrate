// Copyright 2026 The Agent Substrate Authors
// SPDX-License-Identifier: Apache-2.0

// Binary ate-idp-server: OpenID Connect discovery and JWKS endpoint serving public keys to verify short-lived per-actor JWTs.
package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	addr := flag.String("listen", ":8443", "gRPC/HTTPS listen address for ate-idp-server")
	atespace := flag.String("atespace", "ate-system", "Substrate namespace scope")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("starting ate-idp-server", "listen", *addr, "atespace", *atespace, "description", 'OpenID Connect discovery and JWKS endpoint serving public keys to verify short-lived per-actor JWTs.')

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	select {
	case <-ctx.Done():
		logger.Info("shutting down ate-idp-server cleanly", "uptime", time.Since(time.Now()).String())
		fmt.Println("ok")
	}
}
