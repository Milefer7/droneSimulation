package models

import (
	"sync"
)

type ControlCenter struct {
	mu   sync.Mutex
	cond *sync.Cond
}

func NewControlCenter() *ControlCenter {
	cc := &ControlCenter{}
	cc.cond = sync.NewCond(&cc.mu)
	return cc
}

func (cc *ControlCenter) ReceiveIntel(id int) {
	cc.mu.Lock()
	defer cc.mu.Unlock()
	cc.cond.L.Lock()
	defer cc.cond.L.Unlock()
	cc.cond.Broadcast()
}

func (cc *ControlCenter) GetIntel(id int) {
	cc.mu.Lock()
	defer cc.mu.Unlock()
	cc.cond.L.Lock()
	defer cc.cond.L.Unlock()
	cc.cond.Wait()
}
