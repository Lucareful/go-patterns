package main

import "fmt"

func main() {
	electricCooker := new(ElectricCooker)
	// 创建电饭煲指令触发器
	electricCookerInvoker := new(ElectricCookerInvoker)

	// 蒸饭
	steamRiceCommand := NewSteamRiceCommand(electricCooker)
	electricCookerInvoker.SetCookCommand(steamRiceCommand)
	fmt.Println(electricCookerInvoker.ExecuteCookCommand())

	// 煮粥
	cookCongeeCommand := NewCookCongeeCommand(electricCooker)
	electricCookerInvoker.SetCookCommand(cookCongeeCommand)
	fmt.Println(electricCookerInvoker.ExecuteCookCommand())

	// 停止
	shutdownCommand := NewShutdownCommand(electricCooker)
	electricCookerInvoker.SetCookCommand(shutdownCommand)
	fmt.Println(electricCookerInvoker.ExecuteCookCommand())
}
