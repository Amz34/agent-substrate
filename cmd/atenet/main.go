// Copyright 2026 The Agent Substrate Authors
// SPDX-License-Identifier: Apache-2.0

// Binary atenet: Agent-aware networking controller providing Envoy ingress routing (atenet-router) and egress credential injection (atenet-egress).
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
	addr := flag.String("listen", ":8443", "gRPC/HTTPS listen address for atenet")
	atespace := flag.String("atespace", "ate-system", "Substrate namespace scope")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("starting atenet", "listen", *addr, "atespace", *atespace, "description", 'Agent-aware networking controller providing Envoy ingress routing (atenet-router) and egress credential injection (atenet-egress).')

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	select {
	case <-ctx.Done():
		logger.Info("shutting down atenet cleanly", "uptime", time.Since(time.Now()).String())
		fmt.Println("ok")
	}
}
