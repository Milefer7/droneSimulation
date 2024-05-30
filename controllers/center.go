package controllers

import (
	"fmt"
	"github.com/Milefer7/droneSimulation/models"
	"github.com/Milefer7/droneSimulation/utils"
)

var controlCenter *models.ControlCenter

func init() {
	controlCenter = models.NewControlCenter()
	utils.Init()
}

func RunSimulation() {
	scouts := 5
	fighters := 5

	for i := 1; i <= scouts; i++ {
		go func(id int) {
			for {
				controlCenter.ReceiveIntel(id)
				fmt.Printf("侦察无人机 %d: 发送情报\n", id)
				utils.RandomDelay(1000, 2000)
			}
		}(i)
	}

	for i := 1; i <= fighters; i++ {
		go func(id int) {
			for {
				controlCenter.GetIntel(id)
				fmt.Printf("战斗无人机 %d: 接收情报并执行任务\n", id)
				utils.RandomDelay(500, 1500)
			}
		}(i)
	}

	// 运行一段时间后结束仿真
	utils.RandomDelay(10000, 20000)
	fmt.Println("仿真结束")
}
