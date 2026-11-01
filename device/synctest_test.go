/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2025 WireGuard LLC. All Rights Reserved.
 */

package device

import (
	"runtime"
	"testing"
	"testing/synctest"
	"time"
)

// TestRemovedPeerQueuesOutsideSynctestBubble checks that removing a peer
// drains its queues explicitly and clears their finalizers. Otherwise, once
// the peer is collected, the GC finalizer receives from a channel created
// inside the synctest bubble, which is a fatal error outside the bubble.
func TestRemovedPeerQueuesOutsideSynctestBubble(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		pair := genTestPair(t, false)

		// Remove the peer on the first device explicitly, and let Close
		// remove the peer on the second device through RemoveAllPeers.
		for k := range pair[0].dev.peers.keyMap {
			pair[0].dev.RemovePeer(k)
		}
		for i := range pair {
			pair[i].dev.Close()
		}
	})

	for range 5 {
		runtime.GC()
		time.Sleep(10 * time.Millisecond)
	}
}
