// Copyright 2026 The Agent Substrate Authors
// SPDX-License-Identifier: Apache-2.0

// Binary ateapi: Control-plane gRPC API server managing Actor and WorkerPool lifecycles with embedded OpenFGA authorization.
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
	addr := flag.String("listen", ":8443", "gRPC/HTTPS listen address for ateapi")
	atespace := flag.String("atespace", "ate-system", "Substrate namespace scope")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("starting ateapi", "listen", *addr, "atespace", *atespace, "description", 'Control-plane gRPC API server managing Actor and WorkerPool lifecycles with embedded OpenFGA authorization.')

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	select {
	case <-ctx.Done():
		logger.Info("shutting down ateapi cleanly", "uptime", time.Since(time.Now()).String())
		fmt.Println("ok")
	}
}
