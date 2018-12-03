// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package locker

import (
	"sync"
	"time"

	rc "github.com/go-redis/redis"
)

type redis struct {
	prefix string
	mtx    sync.RWMutex
	rds    rc.Cmdable
}

// NewRedisLocker constructor
func NewRedisLocker(rds rc.Cmdable) Locker {
	return &redis{
		prefix: "locker:",
		rds:    rds,
	}
}

func (l *redis) Lock(key string, ttl time.Duration) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	return l.rds.SetNX(l.prefix+key, 1, ttl).Err()
}

var luaRelease = rc.NewScript(`if redis.call("get", KEYS[1]) == ARGV[1] then return redis.call("del", KEYS[1]) else return 0 end`)

func (l *redis) Unlock(key string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	_, err := luaRelease.Run(l.rds, []string{l.prefix + key}, 1).Result()

	return err
}

func (l *redis) IsLocked(key string) bool {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	ok, err := l.rds.Exists(l.prefix + key).Result()
	if err != nil {
		return false
	}

	return ok == 1
}
