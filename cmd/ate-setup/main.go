// Copyright 2026 The Agent Substrate Authors
// SPDX-License-Identifier: Apache-2.0

// Binary ate-setup: Cluster installer and lifecycle manager for deploying ate-system and sample actor workloads.
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
	addr := flag.String("listen", ":8443", "gRPC/HTTPS listen address for ate-setup")
	atespace := flag.String("atespace", "ate-system", "Substrate namespace scope")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("starting ate-setup", "listen", *addr, "atespace", *atespace, "description", 'Cluster installer and lifecycle manager for deploying ate-system and sample actor workloads.')

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	select {
	case <-ctx.Done():
		logger.Info("shutting down ate-setup cleanly", "uptime", time.Since(time.Now()).String())
		fmt.Println("ok")
	}
}
