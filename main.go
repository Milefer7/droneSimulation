package main

import (
	"fmt"
	"sync"
	"time"
)

// 参数定义模块****************************************************************************************************************************
const (
	numReconDrones  = 3 // 侦察无人机数量
	numCombatDrones = 5 // 战斗无人机数量
)

// Intelligence 情报结构体
type Intelligence struct {
	data       string // 情报数据
	taskNumber int    // 任务编号
}

// ControlCenter 控制中心结构体
type ControlCenter struct {
	mu           sync.RWMutex  // 读写锁
	condition    *sync.Cond    // 条件变量，用于协调侦察和战斗无人机
	intelligence *Intelligence // 当前的情报
	taskCounter  int           // 任务计数器
}

//控制中心模块****************************************************************************************************************************

// NewControlCenter 创建新的控制中心实例
func NewControlCenter() *ControlCenter {
	cc := &ControlCenter{
		taskCounter: 1,
	}
	cc.condition = sync.NewCond(cc.mu.RLocker())
	return cc
}

//侦察无人机模块****************************************************************************************************************************

// 侦察无人机函数
func scoutDrone(id int, cc *ControlCenter, submitMu *sync.Mutex) {
	for {
		cc.SubmitIntelligence(id, submitMu)
	}
}

// SubmitIntelligence 侦察无人机提交情报
func (cc *ControlCenter) SubmitIntelligence(id int, submitMu *sync.Mutex) {
	submitMu.Lock()
	defer submitMu.Unlock()
	cc.mu.Lock()
	defer cc.mu.Unlock()
	intel := &Intelligence{data: "侦察到的情报", taskNumber: cc.taskCounter}
	cc.taskCounter++
	// 输出侦察无人机提交情报的信息
	fmt.Printf("\033[34m[侦察无人机%d号]: 提交在%v侦察到的情报 (任务编号: %d)\033[0m\n", id, time.Now().Format("15:04:05"), intel.taskNumber)
	cc.intelligence = intel
	cc.condition.Broadcast() // 通知所有战斗无人机有新的情报
	//time.Sleep(time.Duration(id*2+rand.Intn(3)) * time.Second)
	time.Sleep(5 * time.Second)
}

//战斗无人机模块****************************************************************************************************************************

// 战斗无人机函数
func combatDrone(id int, cc *ControlCenter) {
	for {
		// 输出战斗无人机请求任务的信息
		fmt.Printf("\033[32m[战斗无人机%d号]: 正在请求战斗任务\033[0m\n", id)
		intel := cc.GetIntelligence()
		if intel != nil {
			// 输出战斗无人机收到情报的信息
			fmt.Printf("\033[33m[战斗无人机%d号]: 收到在%v的情报 (任务编号: %d)\033[0m\n", id, time.Now().Format("15:04:05"), intel.taskNumber)
		}
		//time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		time.Sleep(5 * time.Second)
	}
}

// GetIntelligence 获取当前的情报
func (cc *ControlCenter) GetIntelligence() *Intelligence {
	cc.mu.RLock() // 读锁, 多个战斗机都可以读
	defer cc.mu.RUnlock()

	for cc.intelligence == nil {
		cc.condition.Wait() // 等待新的情报，用信号的方式
	}
	return cc.intelligence
}

// 主函数****************************************************************************************************************************
func main() {
	cc := NewControlCenter()
	submitMu := &sync.Mutex{}

	// 启动战斗无人机
	for i := 1; i <= numCombatDrones; i++ {
		go combatDrone(i, cc)
	}

	// 启动侦察无人机2
	for i := 1; i <= numReconDrones; i++ {
		go scoutDrone(i, cc, submitMu)
	}

	// 主程序运行30秒后结束
	time.Sleep(30 * time.Second)
}
