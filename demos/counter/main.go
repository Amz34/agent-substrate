// Copyright 2026 The Agent Substrate Authors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var counter int64

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		val := atomic.AddInt64(&counter, 1)
		fmt.Fprintf(w, "Actor state preserved across suspend/resume cycles. Count: %d
", val)
	})
	_ = http.ListenAndServe(":8080", nil)
}
