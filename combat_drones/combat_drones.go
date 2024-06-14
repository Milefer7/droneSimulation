package combat_drones

import (
	"fmt"
	"time"

	"github.com/Milefer7/droneSimulation/control_center"
)

// CombatDrone 战斗无人机
type CombatDrone struct {
	id int
	cc *control_center.ControlCenter
}

// NewCombatDrone 创建战斗无人机实例
func NewCombatDrone(id int, cc *control_center.ControlCenter) *CombatDrone {
	return &CombatDrone{
		id: id,
		cc: cc,
	}
}

// Run 启动战斗无人机
func (cd *CombatDrone) Run() {
	for {
		// 请求情报数据
		fmt.Printf("\033[32m[战斗无人机%d号]: 正在请求战斗任务\033[0m\n", cd.id) // 使用绿色显示战斗无人机的信息
		data := <-cd.cc.RequestCombatData()
		fmt.Printf("\033[32m[战斗无人机%d号]: 已接收到战斗任务： 根据%v的安排开展秘密打击任务\033[0m\n", cd.id, data) // 使用绿色显示战斗无人机的信息

		// 模拟战斗任务
		time.Sleep(5 * time.Second)
		fmt.Printf("\033[32m[战斗无人机%d号]: 完成战斗任务：%s\033[0m\n", cd.id, data)
	}
}
