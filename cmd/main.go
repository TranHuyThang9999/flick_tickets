package main

import (
	"flag"
	"flick_tickets/common/log"
	"flick_tickets/common/utils"
	"flick_tickets/configs"
	"fmt"
)

func init() {
	log.LoadLogger() // Initialize the logger
	var pathConfig string
	flag.StringVar(&pathConfig, "configs", "./configs/configs.json", "path config")
	flag.Parse()
	configs.LoadConfig(pathConfig)
}
func main() {
	//thuy := "thuynguyen151387@gmail.com"
	//tranhuythang9999@gmail.com

	err := utils.GeneratesQrCodeAndSendQrWithEmail("thuynguyen151387@gmail.com", "highlight obnl egfwerv 666")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("generating highlight")
	// err := utils.SendEmail("thuynguyen151387@gmail.com", "http://localhost:1234/manager/shader/huythang/458478123.png")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}
