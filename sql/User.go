package sql

import (
	"pugg/pojo"
	"pugg/utils"
)

func GetUserInfo(address string)pojo.AAUser  {
	var user pojo.AAUser
	DB.Raw("select * from aa_user where address =?",address).Scan(&user)
	return user
}

func ExistsUser(address string)bool  {
	var e exists
	DB.Raw("select 1 ex from aa_user where address =?",address).Scan(&e)
	if e.Ex==1 {
		return true
	}
	return false

}
func GetFansNum(address string) int {
	var e exists
	DB.Raw("select count(*)ex from aa_user where f_address=?",address).Scan(&e)
	return e.Ex
}

func GetTeamNum(address string)int  {
	var e exists
	DB.Raw("WITH recursive cte(address,f_address,amount,nft,create_time,update_time) " +
		"AS(SELECT address,f_address,amount,nft,create_time,update_time FROM aa_user WHERE address=? " +
		"UNION ALL SELECT t1.address,t1.f_address,t1.amount,t1.nft,t1.create_time,t1.update_time " +
		"FROM aa_user t1 INNER JOIN cte ON t1.f_address=cte.address) " +
		"SELECT count(*)ex  FROM cte WHERE address<>?",address,address).Scan(&e)
	return e.Ex
}

func UpdateBscBlock(block int64,chain int)  {
	DB.Exec("update aa_chain_block set number = ? where id = ?",block,chain)
}



func InsertUser(address string)  {
	code := GenerateCode()
	DB.Exec("insert into aa_user(address,invite) values(?,?)",address,code)
}

func InsertUserAmount(address string,amount float32)  {
	code := GenerateCode()
	DB.Exec("insert into aa_user(address,code,amount) values(?,?,?)",address,code,amount)
}

func InsertUserWithInvite(address string,invite string) bool {
	code := GenerateCode()
	var u AAUser
	DB.Raw("select * from aa_user where invite=?",invite).Scan(&u)
	if len(u.Address)>0 {
	DB.Exec("insert into aa_user(address,code,f_address) values(?,?,?)",address,code,u.Address)
	return true
	}else {
		//DB.Exec("insert into dof_user(address,invite) values(?,?)",address,code)
		return false
	}
}


type exists struct {
	Ex int `db:"ex"`
}


func ExistsTransaction(transaction string)bool  {
	var e exists
	DB.Raw("select 1 ex from aa_transaction where transaction=? limit 1",transaction).Scan(&e)
	if e.Ex==1 {
		return true
	}
	return false

}


func GetBscBlock() int {
	var e exists
	DB.Raw("select number ex from aa_chain_block where id=1").Scan(&e)
	return e.Ex
}
func GetEthBlock() int {
	var e exists
	DB.Raw("select number ex from aa_chain_block where id=2").Scan(&e)
	return e.Ex
}

func GetFans(address string,pageNum int,pageSize int)([]pojo.AAUser,int)  {

	var e exists
	DB.Raw("select count(*) ex from aa_user where f_address = ?",address).Scan(&e)

	var u []pojo.AAUser
	DB.Raw("select * from aa_user where f_address = ? order by id desc limit ? offset ? ",
		address,pageSize,(pageNum-1)*pageSize).Scan(&u)
	return u,e.Ex
}

func GetTeams(address string,pageNum int,pageSize int)([]pojo.AAUser,int)  {
	var e exists
	DB.Raw("WITH recursive cte(address,f_address,amount,nft,create_time,update_time) " +
		"AS(SELECT address,f_address,amount,nft,create_time,update_time FROM aa_user WHERE address=? " +
		"UNION ALL SELECT t1.address,t1.f_address,t1.amount,t1.nft,t1.create_time,t1.update_time " +
		"FROM aa_user t1 INNER JOIN cte ON t1.f_address=cte.address) " +
		"SELECT count(*)ex  FROM cte WHERE address<>?",address,address).Scan(&e)

	var u []pojo.AAUser
	DB.Raw("WITH recursive cte(address,f_address,amount,nft,create_time,update_time) " +
		"AS(SELECT address,f_address,amount,nft,create_time,update_time FROM aa_user WHERE address=? " +
		"UNION ALL SELECT t1.address,t1.f_address,t1.amount,t1.nft,t1.create_time,t1.update_time " +
		"FROM aa_user t1 INNER JOIN cte ON t1.f_address=cte.address) " +
		"SELECT cte.*  FROM cte WHERE address<>? limit ? offset ?",address,address,pageSize,(pageNum-1)*pageSize).Scan(&u)

	return u,e.Ex
}


func UpdateUserAmount(address string,amount float32)  {
	DB.Exec("update dof_user set amount = amount+? where address =?",amount,address)
}

func InsertRechargeRecord(address string,amount float32,transaction string)  {
	DB.Exec("insert into aa_transaction(address,amount,transaction) values(?,?,?)",address,amount,transaction)
}




//生成唯一码·
func GenerateCode() string {
	var randStr =""
	for true {
		randStr = utils.RandStr(6)
		var e exists
		DB.Raw("select 1 ex from aa_user where code = ? limit 1",randStr).Scan(&e)
		if e.Ex==0 {
			break
		}
	}
	return randStr
}
