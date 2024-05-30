package main

import (
	"fmt"
	"github.com/Milefer7/droneSimulation/controllers"
	"github.com/Milefer7/droneSimulation/utils"
)

func main() {
	fmt.Println("无人机战斗小组仿真系统")
	for {
		choice := utils.DisplayMenu()
		switch choice {
		case 1:
			controllers.RunSimulation()
		case 2:
			fmt.Println("退出系统")
			return
		default:
			fmt.Println("无效选项，请重试")
		}
	}
}
