// Copyright 2026 The Agent Substrate Authors
// SPDX-License-Identifier: Apache-2.0

// Binary credential-provider: Out-of-band Kubernetes Secret credential injector for atenet-egress policies.
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
	addr := flag.String("listen", ":8443", "gRPC/HTTPS listen address for credential-provider")
	atespace := flag.String("atespace", "ate-system", "Substrate namespace scope")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("starting credential-provider", "listen", *addr, "atespace", *atespace, "description", 'Out-of-band Kubernetes Secret credential injector for atenet-egress policies.')

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	select {
	case <-ctx.Done():
		logger.Info("shutting down credential-provider cleanly", "uptime", time.Since(time.Now()).String())
		fmt.Println("ok")
	}
}
