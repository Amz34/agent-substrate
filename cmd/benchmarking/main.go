// Copyright 2026 The Agent Substrate Authors
// SPDX-License-Identifier: Apache-2.0

// Binary benchmarking: Synthetic load generator and glutton resource consumer for testing 30x oversubscription density.
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
	addr := flag.String("listen", ":8443", "gRPC/HTTPS listen address for benchmarking")
	atespace := flag.String("atespace", "ate-system", "Substrate namespace scope")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("starting benchmarking", "listen", *addr, "atespace", *atespace, "description", 'Synthetic load generator and glutton resource consumer for testing 30x oversubscription density.')

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	select {
	case <-ctx.Done():
		logger.Info("shutting down benchmarking cleanly", "uptime", time.Since(time.Now()).String())
		fmt.Println("ok")
	}
}
