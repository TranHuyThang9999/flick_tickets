package main

import (
	"flick_tickets/common/log"
	"flick_tickets/common/utils"
	"flick_tickets/configs"
)

func init() {
	configs.LoadConfig("./configs/configs.json")
	log.LoadLogger()
}
func main() {

	thuy := "thuynguyen151387@gmail.com"
	err := utils.SendOtpToEmail(thuy, "gui demo", utils.GenerateUniqueKey())
	if err != nil {
		log.Info(err.Error())
		return
	}
	log.Info("ok")
}
