package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type configuration struct {
	Port 		string `toml:"port"`
	Log			string `toml:"log_path"`
	ApiKey   	string `toml:"api_key"`
	Ethereum	Ethereum	`toml:"ethereum"`
	Database    Database `toml:"database"`
	RedisNetwork string `toml:"redis_network"`
	RedisPassword string `toml:"redis_password"`
	Arg Arg `toml:"arg"`
	LotteryRate int `toml:"lottery_rate"`
}
type Ethereum struct {
	Network string 	`toml:"network"`
	EthNetwork string `toml:"eth_network"`
	TronNetwork string `toml:"tron_network"`
	ChainId	  int	`toml:"chainId"`
	KeyStore 	string	`toml:"keyStore"`
	KeyPath string		`toml:"keyPath"`
	DofAddress	string	`toml:"pugg_address"`
	DofKey string `toml:"dof_key"`
	FuncAddress string `toml:"func_address"`
	HolderSlipKeyPath string `toml:"holder_slip_keyPath"`
	LpSlipKeyPath string `toml:"lp_slip_keyPath"`
	UsdtAddress string `toml:"usdt_address"`
	StaticAddress string `toml:"static_address"`
	TryAddress string `toml:"try_address"`
	ReceiveDbtcAddress string `toml:"receive_dbtc_address"`
}

type Database struct {
	DBHost       string  `toml:"db_host"`
	DBPort       string  `toml:"db_port"`
	DBSchema     string  `toml:"db_schema"`
	DBUserName 	 string  `toml:"db_username"`
	DBPassword   string  `toml:"db_password"`
	DBArgs       string  `toml:"db_args"`
}

type Arg struct {
	Morning string `toml:"morning"`
	Afternoon string `toml:"afternoon"`
	DofAmount int64 `toml:"dof_amount"`
}

var Cfg  *configuration


func Config() *configuration {
	if _, err := toml.DecodeFile("./config/config.toml", &Cfg); err != nil {
		log.Println(err)
	}
	return Cfg
}
