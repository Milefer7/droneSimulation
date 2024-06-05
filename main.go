package main

import (
	"github.com/Milefer7/droneSimulation/combat_drones"
	"github.com/Milefer7/droneSimulation/control_center"
	"github.com/Milefer7/droneSimulation/scout_drones"
	"time"
)

func main() {
	// 创建控制中心实例
	cc := control_center.NewControlCenter()

	// 启动控制中心
	go cc.Run()

	// 创建侦察无人机
	numScoutDrones := 2
	for i := 0; i < numScoutDrones; i++ {
		sd := scout_drones.NewScoutDrone(i, cc)
		go sd.Run()
	}

	// 创建战斗无人机
	numCombatDrones := 5
	for i := 0; i < numCombatDrones; i++ {
		cd := combat_drones.NewCombatDrone(i, cc)
		go cd.Run()
	}

	// 让程序运行一段时间以便观察输出
	time.Sleep(90 * time.Second)
}
