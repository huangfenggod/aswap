package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"pugg/config"
	"time"
)

var DB *gorm.DB
func InitDatabase()  {
	sqlSource := config.Cfg.Database.DBUserName + ":" +config.Cfg.Database.DBPassword + "@tcp(" + config.Cfg.Database.DBHost + ":" + config.Cfg.Database.DBPort + ")/"+config.Cfg.Database.DBSchema+"?"+config.Cfg.Database.DBArgs
	db, err := gorm.Open("mysql", sqlSource)
	if err !=nil {
		log.Fatal(err)
	}
	db.DB().SetMaxIdleConns(100)
	DB =db
}



type AAUser struct {
	Address string `db:"address"`
	FAddress string `db:"f_address"`
	Code string `db:"code"`
	Amount float32 `db:"amount"`
	Nft int `db:"nft"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}


type AATransaction struct {
	Address string `db:"address"`
	Amount float32 `db:"amount"`
	Transaction string `db:"transaction"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}


