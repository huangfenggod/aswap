package service

import (
	"errors"
	"log"
	"pugg/config"
	"pugg/sql"
	"strconv"
	"time"
)
//chain 1是bsc，2是eth
func CheckBscTransfer(chain int)  {
	var r Request
	if chain==1 {
		i :=sql.GetBscBlock()+1
		r.Url="https://api.bscscan.com/api?module=logs&action=getLogs&fromBlock="+strconv.Itoa(i)+"&toBlock=latest&address=0x55d398326f99059ff775485246999027b3197955&topic0=0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef&topic0_2_opr=and&topic2=0x000000000000000000000000aDcD71352e8802a8dA38a7fF331a6D18dC1b518D&apikey=SPT2S73P7YVTRMPS3CHCAYSI94HXQ4B7VV"
		log.Println("执行bsc链")
	}

	if  chain==2{
		i :=sql.GetEthBlock()+1
		r.Url="https://api.etherscan.io/api?module=logs&action=getLogs&fromBlock="+strconv.Itoa(i)+"&toBlock=latest&address=0xdac17f958d2ee523a2206206994597c13d831ec7&topic0=0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef&topic0_2_opr=and&topic2=0x000000000000000000000000063E2Bc497999EB65Dd66725D9237fe8D241581b&page=1&offset=1000&apikey=H54IJ6ST42QW9GAAFBUZPM7D2P3BRUXR6D"
		//log.Println("执行eth链")
	}
	//r.Url="https://api.bscscan.com/api?module=logs&action=getLogs&fromBlock="+strconv.Itoa(i)+"&toBlock=latest&address=0x55d398326f99059ff775485246999027b3197955&topic0=0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef&topic0_2_opr=and&topic2=0x000000000000000000000000FC9b8b3de9E3BD79caB33236e6b45BAf218Bd9aB&apikey=SPT2S73P7YVTRMPS3CHCAYSI94HXQ4B7VV"
	//r.Url="https://api.etherscan.io/api?module=logs&action=getLogs&fromBlock="+strconv.Itoa(i)+"&toBlock=latest&address=0xdac17f958d2ee523a2206206994597c13d831ec7&topic0=0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef&topic0_2_opr=and&topic2=0x000000000000000000000000dAC17F958D2ee523a2206206994597C13D831ec7&page=1&offset=1000&apikey=H54IJ6ST42QW9GAAFBUZPM7D2P3BRUXR6D"

	get := r.BscGet()
	for k,v:=range get{
		if k=="blockNumber"{
				sql.UpdateBscBlock(v.(int64),chain)
		}
		//map[""]
		if k=="data" {
			for k1,v1 :=range v.(map[string]interface{}){
				transaction := sql.ExistsTransaction(k1)
				//transaction不存在,充值
				if !transaction {
					for k2,v2 :=range v1.(map[string]float32){
						//查询k2地址是否存在
						if v2>=100 {
							user := sql.ExistsUser(k2)
							if user {
								sql.UpdateUserAmount(k2,v2)

							}else {
								sql.InsertUserAmount(k2,v2)
							}
							sql.InsertRechargeRecord(k2,v2,k1)
							log.Printf("查链充值:金额:%f  地址:%s  transaction:%s   链:%d",v2,k2,k1,chain)
						}
					}
				}
			}
		}
	}

}

func CheckTronTransfer()  {
	var r Request
	r.Url="https://apilist.tronscanapi.com/api/new/token_trc20/transfers?limit=100&start=0&sort=-timestamp&count=true&filterTokenValue=1&relatedAddress=TQv9PszjG5HrXnCbwZWuS6RG9ignAaUV4M&contract_address=TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	get := r.TronGet()
	log.Println("执行tron链")
	for k,v :=range get{
		transaction := sql.ExistsTransaction(k)
		if !transaction {

			for k1,v1 :=range v.(map[string]interface{}){
				user := sql.ExistsUser(k1)
				if user {
					sql.UpdateUserAmount(k1,v1.(float32))

				}else {
					sql.InsertUserAmount(k1,v1.(float32))
				}
				sql.InsertRechargeRecord(k1,v1.(float32),k)
			}
		}
	}
}


func Charge(address string,amount float32,transaction string,chainId  string)(bool,error){
	go func() {
		charge(address,amount,transaction,chainId)
	}()
	return true,nil
}

func charge(address string,amount float32,transaction string,chainId  string)(bool,error)  {
	chain :=1
	wait :=3
	if chainId=="56"{
		chain=1
		wait=3
	}
	if chainId=="1"{
		chain=2
		wait=14
	}

	time.Sleep(time.Second *time.Duration(wait))
	//这里dbtc确认转账
	log.Printf("address:%s,amount:%f,transaction:%s",address,amount,transaction)
	receive := CheckReceiveU(address,amount,transaction,chain)
	i:=0
	if !receive {
		time.Sleep(time.Second * time.Duration(wait))
		receive = CheckReceiveU(address,amount,transaction,chain)
		i++
		if i>=2 && !receive {
			log.Printf("not receiver u address:%s",address)
			return false,errors.New("not receive usdt")
		}
	}
	user := sql.ExistsUser(address)
	if !user {
		sql.InsertUserAmount(address,amount)
	}else {
		sql.UpdateUserAmount(address,amount)
		sql.InsertRechargeRecord(address,amount,transaction)
		log.Printf("直接充值:金额:%f  地址:%s  transaction:%s   链:%d",amount,address,transaction,chain)
	}
	return true,nil
}
func CheckReceiveU(address string,amount float32,transaction string,chain int) bool {
	//查询是否有transaction\
	hash := CheckTransferUBySqlAndEth(transaction,address,amount,chain)
	return hash
}

func CheckTransferUBySqlAndEth(transaction string,to string,amount float32,chain int) bool {
	exists := sql.ExistsTransaction(transaction)
	if exists {
		return false
	}
	var  hash  bool
	if chain==1{
		hash, _ = CheckTransferU(transaction,config.Cfg.Ethereum.KeyPath,amount)
	}
	if chain==2{
		hash, _ = EthCheckTransferU(transaction,config.Cfg.Ethereum.KeyPath,amount)
	}

	return hash
}
