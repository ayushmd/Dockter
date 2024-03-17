package internal

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

type Backend struct {
	URL            *url.URL
	InitialWeight  float64
	CurrentConnect int
	Alive          bool
	mux            sync.RWMutex
	ReverseProxy   *httputil.ReverseProxy
}

// SetAlive for this backend
func (b *Backend) SetAlive(alive bool) {
	b.mux.Lock()
	b.Alive = alive
	b.mux.Unlock()
}

// IsAlive returns true when backend is alive
func (b *Backend) IsAlive() (alive bool) {
	b.mux.RLock()
	alive = b.Alive
	b.mux.RUnlock()
	return
}

func (b *Backend) SetInitialWeight(score float64) {
	b.mux.Lock()
	b.InitialWeight = score
	b.mux.Unlock()
}

func (b *Backend) IncrConnection() {
	b.mux.Lock()
	b.CurrentConnect += 1
	b.mux.Unlock()
}

func (b *Backend) DecrConnection() {
	b.mux.Lock()
	b.CurrentConnect -= 1
	b.mux.Unlock()
}

type ServerPool struct {
	backends []*Backend
}

func (s *ServerPool) AddBackend(backend *Backend) {
	s.backends = append(s.backends, backend)
}
