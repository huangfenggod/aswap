package main

import (
	"pugg/config"
	"pugg/cron"
	"pugg/log"
	"pugg/router"
	"pugg/service"
	"pugg/sql"
)

func main()  {
	config.Config()
	log.InitLog()
	sql.InitDatabase()
	service.InitDial()
	cron.InitCron()
	router.InitGin()
	router.ApiGroup()
	router.GIN.Run(":8081")


}

