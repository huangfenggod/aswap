package cron

import (
	"github.com/robfig/cron/v3"
	"log"
	"pugg/service"
)
func InitCron() {

	crontab := cron.New(cron.WithSeconds())

	spec :="0 */1 * * * ?"
	//spec1 :="0 */2 * * * ?"

	//执行检查充值
	crontab.AddFunc(spec,checkBscRecharge)

	//crontab.AddFunc(spec1,checkEthRecharge)
	//crontab.AddFunc(spec2,checkTronRecharge)


	log.Println("cron task begin")
	crontab.Start()
}

func checkBscRecharge()  {
	service.CheckBscTransfer(1)
}
func checkEthRecharge()  {
	service.CheckBscTransfer(2)
}





