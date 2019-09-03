package uniqid

import (
	"context"
	"time"
)

const (
	buffer = 10
)

// Manager unique id manager
type Manager struct {
	idChan chan uint64
	b      uint64
}

// NewManager create a unique id manager
func NewManager(ctx context.Context) *Manager {
	m := &Manager{
		idChan: make(chan uint64, buffer),
		b:      uint64(time.Now().Unix()),
	}
	go func() {
	out:
		for {
			m.b++
			select {
			case <-ctx.Done():
				return
			case m.idChan <- pseudoEncrypt(m.b):
				continue out
			}
		}
	}()
	return m
}

// NewID generate a new unique id
func (m *Manager) NewID() uint64 {
	return <-m.idChan
}
