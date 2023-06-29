package pojo

import "time"

//用户表
type AAUser struct {
	Address string `db:"address"`
	FAddress string `db:"f_address"`
	Code string `db:"code"`
	Amount float32 `db:"amount"`
	Nft int `db:"nft"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}




type UserInfo struct {
	IsExists  bool
	User      AAUser
	InviteUrl string
	Fans int
	Teams int
}
