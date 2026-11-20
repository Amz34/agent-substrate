// Copyright 2026 The Agent Substrate Authors
// SPDX-License-Identifier: Apache-2.0

package benchmarking

// AllocateActorFootprint simulates stateful agent memory and file-descriptor consumption prior to snapshot hibernation.
func AllocateActorFootprint(megabytes int) []byte {
	buf := make([]byte, megabytes*1024*1024)
	for i := 0; i < len(buf); i += 4096 {
		buf[i] = 1
	}
	return buf
}
