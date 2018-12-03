// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package locker

import (
	"time"
)

// Locker interface
type Locker interface {
	Lock(key string, ttl time.Duration) error
	Unlock(key string) error
	IsLocked(key string) bool
}
