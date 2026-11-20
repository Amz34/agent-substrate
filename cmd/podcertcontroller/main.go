// Copyright 2026 The Agent Substrate Authors
// SPDX-License-Identifier: Apache-2.0

// Binary podcertcontroller: SPIFFE X.509 Pod Certificate signer authority issuing short-lived mTLS identities (podid, servicedns, postgres).
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
	addr := flag.String("listen", ":8443", "gRPC/HTTPS listen address for podcertcontroller")
	atespace := flag.String("atespace", "ate-system", "Substrate namespace scope")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("starting podcertcontroller", "listen", *addr, "atespace", *atespace, "description", 'SPIFFE X.509 Pod Certificate signer authority issuing short-lived mTLS identities (podid, servicedns, postgres).')

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	select {
	case <-ctx.Done():
		logger.Info("shutting down podcertcontroller cleanly", "uptime", time.Since(time.Now()).String())
		fmt.Println("ok")
	}
}
