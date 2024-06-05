package control_center

import (
	"fmt"
	"sync"
)

// ControlCenter 控制中心
type ControlCenter struct {
	scoutData      chan string      // 侦察数据
	combatRequests chan chan string // 战斗请求
	mutex          sync.Mutex
}

// NewControlCenter 创建控制中心实例
func NewControlCenter() *ControlCenter {
	return &ControlCenter{
		scoutData:      make(chan string),      // 设置一个侦察数据提交通道。一次只可以一架侦察机往控制中心提交数据，所以使用了无缓冲通道
		combatRequests: make(chan chan string), // 设置一个战斗请求通道。多架战斗机可以同时请求数据，所以使用了有缓冲通道
	}
}

// Run 启动控制中心
func (cc *ControlCenter) Run() {
	for {
		select {
		case data := <-cc.scoutData: // 表示当有侦察数据提交时
			fmt.Printf("\033[34m[控制中心]: %s\033[0m\n", data) // 使用蓝色显示控制中心的信息
			cc.mutex.Lock()
			// 向所有等待的战斗无人机发送情报
			for {
				select {
				case request := <-cc.combatRequests:
					request <- data // 向取出的通道发送数据
				default:
					cc.mutex.Unlock()
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
