package session

import (
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"sync"
	"sync/atomic"
	"time"
)

// NetworkStackLatencyHandler handles the NetworkStackLatency packet.
type NetworkStackLatencyHandler struct {
	mu        sync.Mutex
	forms     map[uint32]form.Form
	currentID atomic.Uint32
}

// Handle ...
func (h *NetworkStackLatencyHandler) Handle(p packet.Packet, s *Session, tx *world.Tx, c Controllable) error {
	pk := p.(*packet.NetworkStackLatency)
	s.latencyMu.Lock()
	defer s.latencyMu.Unlock()
	if s.lastTimestamp != nil && pk.Timestamp/1000000000 == s.lastTimestamp.Unix()/1000 {
		latency := time.Now().Sub(*s.lastTimestamp) / 4
		s.lastNetworkStackLatency.Store(&latency)
		s.lastTimestamp = nil
	}
	return nil
}
