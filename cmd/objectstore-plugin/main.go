// Copyright 2026 The Agent Substrate Authors
// SPDX-License-Identifier: Apache-2.0

// Binary objectstore-plugin: Pluggable snapshot persistence driver for GCS, S3, and NVMe node-local cache tiers.
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
	addr := flag.String("listen", ":8443", "gRPC/HTTPS listen address for objectstore-plugin")
	atespace := flag.String("atespace", "ate-system", "Substrate namespace scope")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("starting objectstore-plugin", "listen", *addr, "atespace", *atespace, "description", 'Pluggable snapshot persistence driver for GCS, S3, and NVMe node-local cache tiers.')

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	select {
	case <-ctx.Done():
		logger.Info("shutting down objectstore-plugin cleanly", "uptime", time.Since(time.Now()).String())
		fmt.Println("ok")
	}
}
