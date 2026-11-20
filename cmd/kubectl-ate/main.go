// Copyright 2026 The Agent Substrate Authors
// SPDX-License-Identifier: Apache-2.0

// Binary kubectl-ate: Kubernetes CLI plugin for creating, inspecting, proxying, and teleporting Substrate actors.
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
	addr := flag.String("listen", ":8443", "gRPC/HTTPS listen address for kubectl-ate")
	atespace := flag.String("atespace", "ate-system", "Substrate namespace scope")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("starting kubectl-ate", "listen", *addr, "atespace", *atespace, "description", 'Kubernetes CLI plugin for creating, inspecting, proxying, and teleporting Substrate actors.')

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	select {
	case <-ctx.Done():
		logger.Info("shutting down kubectl-ate cleanly", "uptime", time.Since(time.Now()).String())
		fmt.Println("ok")
	}
}
