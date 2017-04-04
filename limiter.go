package main

import (
	"time"
)

type Limiter struct {
	Ids      map[string]int64
	Cooldown int64
}

func NewLimiter(cooldown int) *Limiter {
	return &Limiter{
		Ids:      make(map[string]int64),
		Cooldown: int64(cooldown) * int64(time.Second),
	}
}

func (l *Limiter) Check(id string) bool {
	now := time.Now().UnixNano()
	updated, ok := l.Ids[id]
	return !ok || (now-updated) > l.Cooldown
}

func (l *Limiter) Add(id string) {
	l.Ids[id] = time.Now().UnixNano()
}
