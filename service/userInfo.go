package service

import (
	"pugg/pojo"
	"pugg/sql"
	"strings"
)


func GetUserInfo(address string,invite string)pojo.UserInfo  {
	var ui pojo.UserInfo

	existsUser := sql.ExistsUser(address)

	if !existsUser {
		if len(invite)>0 {
			split := strings.Split(invite, "#/")
			sql.InsertUserWithInvite(address,split[0])
		}else {
			ui.IsExists=false

			return ui
			//sql.InsertUser(address)
		}
	}

	user := sql.GetUserInfo(address)
	//if len(user.Address)>0 {
	//	ui.IsExists=true
	//}
	ui.User = user

	//推广链接
	if len(user.Code)>1 {
		ui.InviteUrl = "http://144.34.167.208/?invite="+user.Code
	}
	ui.Fans=sql.GetFansNum(address)
	ui.Teams = sql.GetTeamNum(address)

	return ui
}


func GetFans(address string,pageNum int,pageSize int)([]pojo.AAUser,int)  {
	fans, i := sql.GetFans(address, pageNum, pageSize)
	return fans,i
}

func GetTeams(address string,pageNum int,pageSize int)([]pojo.AAUser,int) {
	teams, i := sql.GetTeams(address, pageNum, pageSize)
	return teams,i
}



//参与规则
func GetNotices(lang string,i int)string  {
	x :=1
	if lang=="en-us"{
		if i==1 {
			x=2
		}
		if i==2 {
			x=4
		}

	}

	if lang=="zh-cn"{
		if i==1 {
			x=1
		}
		if i==2 {
			x=3
		}

	}
	if lang=="jp"{
		if i==1 {
			x=5
		}
		if i==2 {
			x=6
		}

	}
	if lang=="fr"{
		if i==1 {
			x=7
		}
		if i==2 {
			x=8
		}

	}
	if lang=="gr"{
		if i==1 {
			x=9
		}
		if i==2 {
			x=10
		}

	}

	return string(x)
}
