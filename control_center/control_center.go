package control_center

import (
	"container/list"
	"fmt"
	"sync"
)

// ControlCenter 控制中心
type ControlCenter struct {
	CombatDrone
	ScoutDrone
}

// 创建侦察机结构体
type ScoutDrone struct {
	scoutData   chan string // 侦察数据
	taskQueue   *list.List  // 任务队列
	taskCounter int         // 任务编号计数器
	mutex       sync.Mutex
	cond        *sync.Cond
}

// 创建战斗机结构体
type CombatDrone struct {
	combatRequests chan chan string // 战斗请求
	combatNum      int
	sendMsgNum     int
}

// NewControlCenter 创建控制中心实例
func NewControlCenter(numScoutDrones int) *ControlCenter {
	scoutDrone := ScoutDrone{
		scoutData:   make(chan string), // 设置一个侦察数据提交通道。一次只可以一架侦察机往控制中心提交数据，所以使用了无缓冲通道
		taskQueue:   list.New(),        // 初始化任务编号队列
		taskCounter: 0,
	}
	scoutDrone.cond = sync.NewCond(&scoutDrone.mutex)

	combatDrone := CombatDrone{
		combatRequests: make(chan chan string, numScoutDrones), // 设置一个战斗请求通道。多架战斗机可以同时请求数据，所以使用了有缓冲通道
		combatNum:      numScoutDrones,
		sendMsgNum:     0,
	}
	return &ControlCenter{
		CombatDrone: combatDrone,
		ScoutDrone:  scoutDrone,
	}
}

// Pop 从队列中获取并移除第一个元素
func (cc *ControlCenter) Pop() {
	cc.mutex.Lock()
	defer cc.mutex.Unlock()
	if cc.taskQueue.Len() > 0 {
		front := cc.taskQueue.Front()
		cc.taskQueue.Remove(front)
	}
}

// Run 启动控制中心
func (cc *ControlCenter) Run() {
	// 处理侦察数据
	go func() {
		for {
			select {
			case scoutData := <-cc.scoutData:
				// 处理侦察数据（生成任务编号，并将任务编号加入队列）
				cc.mutex.Lock()
				cc.taskCounter++
				taskID := cc.taskCounter
				cc.taskQueue.PushBack(taskID) // 将任务编号添加到队列
				fmt.Printf("\033[34m[控制中心]: %s (任务编号: %d)\033[0m\n", scoutData, taskID)
				cc.cond.Signal() // 通知有新任务到来
				cc.mutex.Unlock()
			}
		}
	}()

	// 处理战斗请求
	go func() {
		for {
			cc.mutex.Lock()
			for cc.taskQueue.Len() == 0 {
				cc.cond.Wait() // 等待有新任务
			}
			//cc.mutex.Lock()
			taskID := cc.taskQueue.Front().Value.(int)
			cc.mutex.Unlock()

			select {
			case request := <-cc.combatRequests:
				request <- fmt.Sprintf("任务编号: %d", taskID)
				cc.mutex.Lock()
				cc.sendMsgNum++
				if cc.sendMsgNum == cc.combatNum {
					cc.Pop()
					cc.sendMsgNum = 0
				}
				cc.mutex.Unlock()
			}
		}
	}()
}

// SubmitScoutData 提交侦察数据
func (cc *ControlCenter) SubmitScoutData(data string) {
	cc.scoutData <- data
}

// RequestCombatData 请求战斗数据
func (cc *ControlCenter) RequestCombatData() chan string {
	response := make(chan string)
	cc.combatRequests <- response // 向控制中心的战斗请求通道提交一个通道，用于接收数据
	return response
}
