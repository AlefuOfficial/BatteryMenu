package main

import (
	"fmt"
	"time"

	"github.com/caseymrm/menuet"
	"github.com/distatus/battery"
)

func BatteryBar() {
	for {
		batteries, err := battery.GetAll()
		if err != nil {
			fmt.Println("Could not get battery info!", err)
			return
		}
		for _, battery := range batteries {
			fmt.Println(battery.Full)
			menuet.App().SetMenuState(&menuet.MenuState{
				Title: fmt.Sprint(battery.Full),
			})
		}
		time.Sleep(5 * time.Second)
	}
}

func main() {
	go BatteryBar()
	menuet.App().RunApplication()
}
