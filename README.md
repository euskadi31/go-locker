Go Locker ![Last release](https://img.shields.io/github/release/euskadi31/go-locker.svg)
=========

[![Go Report Card](https://goreportcard.com/badge/github.com/euskadi31/go-locker)](https://goreportcard.com/report/github.com/euskadi31/go-locker)

| Branch  | Status | Coverage |
|---------|--------|----------|
| master  | [![Build Status](https://img.shields.io/travis/euskadi31/go-locker/master.svg)](https://travis-ci.org/euskadi31/go-locker) | [![Coveralls](https://img.shields.io/coveralls/euskadi31/go-locker/master.svg)](https://coveralls.io/github/euskadi31/go-locker?branch=master) |
| develop | [![Build Status](https://img.shields.io/travis/euskadi31/go-locker/develop.svg)](https://travis-ci.org/euskadi31/go-locker) | [![Coveralls](https://img.shields.io/coveralls/euskadi31/go-locker/develop.svg)](https://coveralls.io/github/euskadi31/go-locker?branch=develop) |

Distributed Lock Manager

## Example

```go
import "github.com/euskadi31/go-locker"

rds := redis.New()

locker := server.NewRedisLocker(rds)

locker.Lock("foo", 1*time.Hour)
defer locker.Unlock("foo")

locker.IsLocked("foo") // true or false

```


## License

go-locker is licensed under [the MIT license](LICENSE.md).
