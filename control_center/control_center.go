package control_center

import (
	"container/list"
	"fmt"
	"sync"
)

// ControlCenter 控制中心
type ControlCenter struct {
	combatRequests chan chan string // 战斗请求
	scoutQueue
}

// 创建侦察数据队列结构体
type scoutQueue struct {
	scoutData   chan string // 侦察数据
	taskQueue   *list.List  // 任务队列
	taskCounter int         // 任务编号计数器
	mutex       sync.Mutex
}

// NewControlCenter 创建控制中心实例
func NewControlCenter() *ControlCenter {
	return &ControlCenter{
		combatRequests: make(chan chan string), // 设置一个战斗请求通道。多架战斗机可以同时请求数据，所以使用了有缓冲通道
		scoutQueue: scoutQueue{
			scoutData:   make(chan string), // 设置一个侦察数据提交通道。一次只可以一架侦察机往控制中心提交数据，所以使用了无缓冲通道
			taskQueue:   list.New(),        // 初始化任务编号队列
			taskCounter: 0,
		},
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
	for {
		select {
		case scoutData := <-cc.scoutData:
			// 处理侦察数据（生成任务编号，并将任务编号加入队列）
			cc.mutex.Lock()
			cc.taskCounter++
			taskID := cc.taskCounter
			cc.taskQueue.PushBack(taskID) // 将任务编号添加到队列
			fmt.Printf("\033[34m[控制中心]: %s (任务编号: %d)\033[0m\n", scoutData, taskID)
			cc.mutex.Unlock()

			// 向所有等待的战斗无人机发送情报
			for {
				select {
				case request := <-cc.combatRequests: // 从战斗请求通道中获取一个通道
					request <- fmt.Sprintf("（任务编号: %d）", cc.taskQueue.Front().Value.(int))
				default:
					cc.Pop()
					goto next
				}
			}
		}
	next:
	}
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
