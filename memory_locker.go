// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package locker

import (
	"sync"
	"time"
)

type memory struct {
	mtx   sync.RWMutex
	locks map[string]struct{}
}

// NewMemoryLocker constructor
func NewMemoryLocker() Locker {
	return &memory{
		locks: make(map[string]struct{}),
	}
}

func (l *memory) Lock(key string, ttl time.Duration) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	l.locks[key] = struct{}{}

	return nil
}

func (l *memory) Unlock(key string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	delete(l.locks, key)

	return nil
}

func (l *memory) IsLocked(key string) bool {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	_, ok := l.locks[key]

	return ok
}
