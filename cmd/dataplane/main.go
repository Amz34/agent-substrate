// Copyright 2026 The Agent Substrate Authors
// SPDX-License-Identifier: Apache-2.0

// Binary dataplane: Low-latency atunnel packet forwarder bridging ingress router traffic to private /30 actor network namespaces.
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
	addr := flag.String("listen", ":8443", "gRPC/HTTPS listen address for dataplane")
	atespace := flag.String("atespace", "ate-system", "Substrate namespace scope")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("starting dataplane", "listen", *addr, "atespace", *atespace, "description", 'Low-latency atunnel packet forwarder bridging ingress router traffic to private /30 actor network namespaces.')

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	select {
	case <-ctx.Done():
		logger.Info("shutting down dataplane cleanly", "uptime", time.Since(time.Now()).String())
		fmt.Println("ok")
	}
}
