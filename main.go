package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numReconDrones  = 3 // 侦察无人机数量
	numCombatDrones = 5 // 战斗无人机数量
)

type Intelligence struct {
	data       string
	taskNumber int
}

type ControlCenter struct {
	mu           sync.RWMutex
	condition    *sync.Cond
	intelligence *Intelligence
	taskCounter  int
}

func NewControlCenter() *ControlCenter {
	cc := &ControlCenter{
		taskCounter: 1,
	}
	cc.condition = sync.NewCond(cc.mu.RLocker())
	return cc
}

func (cc *ControlCenter) GetIntelligence() *Intelligence {
	cc.mu.RLock()
	defer cc.mu.RUnlock()

	for cc.intelligence == nil {
		cc.condition.Wait() // 等待新的情报
	}

	return cc.intelligence
}

func (cc *ControlCenter) SubmitIntelligence(intel *Intelligence, id int) {
	cc.mu.Lock()
	defer cc.mu.Unlock()

	fmt.Printf("\033[34m[侦察无人机%d号]: 提交在%v侦察到的情报 (任务编号: %d)\033[0m\n", id, time.Now().Format("15:04:05"), intel.taskNumber)
	cc.intelligence = intel
	cc.taskCounter++
	cc.condition.Broadcast() // 通知所有战斗无人机有新的情报
	time.Sleep(time.Duration(id*2+rand.Intn(10)) * time.Second)
}

func scoutDrone(id int, cc *ControlCenter) {
	for {
		intel := &Intelligence{data: "侦察到的情报", taskNumber: cc.taskCounter}
		cc.SubmitIntelligence(intel, id)
	}
}

func combatDrone(id int, cc *ControlCenter) {
	for {
		fmt.Printf("\033[32m[战斗无人机%d号]: 正在请求战斗任务\033[0m\n", id)
		intel := cc.GetIntelligence()
		if intel != nil {
			fmt.Printf("\033[33m[战斗无人机%d号]: 收到在%v的情报 (任务编号: %d)\033[0m\n", id, time.Now().Format("15:04:05"), intel.taskNumber)
		}
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	cc := NewControlCenter()

	for i := 1; i <= numReconDrones; i++ {
		go scoutDrone(i, cc)
	}

	for i := 1; i <= numCombatDrones; i++ {
		go combatDrone(i, cc)
	}

	time.Sleep(30 * time.Second)
}
