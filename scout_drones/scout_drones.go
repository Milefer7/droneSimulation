package scout_drones

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Milefer7/droneSimulation/control_center"
)

// ScoutDrone 侦察无人机
type ScoutDrone struct {
	id int
	cc *control_center.ControlCenter
}

// NewScoutDrone 创建侦察无人机实例
func NewScoutDrone(id int, cc *control_center.ControlCenter) *ScoutDrone {
	return &ScoutDrone{
		id: id,
		cc: cc,
	}
}

// Run 启动侦察无人机
func (sd *ScoutDrone) Run() {
	for {
		// 模拟获取情报数据
		time.Sleep(time.Duration(rand.Intn(30)) * time.Second)
		t := time.Now().Format("15:04:05")
		data := fmt.Sprintf("在%v获取到情报", t)
		fmt.Printf("\033[33m[侦察无人机%d号]: 提交在%v侦察到的情报\033[0m\n", sd.id, t) // 使用黄色显示侦察无人机的信息
		sd.cc.SubmitScoutData(data)
	}
}
