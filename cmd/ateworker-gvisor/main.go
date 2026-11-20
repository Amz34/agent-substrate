// Copyright 2026 The Agent Substrate Authors
// SPDX-License-Identifier: Apache-2.0

// Binary ateworker-gvisor: Interior sandbox supervisor executing runsc checkpoint and restore commands inside gVisor worker pods.
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
	addr := flag.String("listen", ":8443", "gRPC/HTTPS listen address for ateworker-gvisor")
	atespace := flag.String("atespace", "ate-system", "Substrate namespace scope")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("starting ateworker-gvisor", "listen", *addr, "atespace", *atespace, "description", 'Interior sandbox supervisor executing runsc checkpoint and restore commands inside gVisor worker pods.')

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	select {
	case <-ctx.Done():
		logger.Info("shutting down ateworker-gvisor cleanly", "uptime", time.Since(time.Now()).String())
		fmt.Println("ok")
	}
}
