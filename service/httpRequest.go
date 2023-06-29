package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"pugg/utils"
	"strconv"
	"strings"
)

type Request struct {
	Url string
}

func (r *Request)BscGet() map[string]interface{}{
	resp, err := http.Get(r.Url)
	if err != nil {
		log.Printf("request err:%s",err.Error())
		return nil
	}
	contains := strings.Contains(r.Url, "etherscan")
	chain :=0 //0是bsc ，1是eth
	if contains{
		chain=1
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var value interface{}
	err1 := json.Unmarshal(body, &value)
	if err1!=nil {
		log.Printf("request err:%s",err1.Error())
		return nil
	}
	data := value.(map[string]interface{})
	//fmt.Println(data)
	if data["status"]=="0" {
		return nil
	}

	for k,v := range data {
		if k=="result" {
			blockNumber :=int64(0)
			res := make(map[string]interface{},2)
			tmp := make(map[string]interface{},len(v.([]interface{})))
			for _,j:=range v.([]interface{}){
				m := j.(map[string]interface{})
				replace := strings.Replace(m["data"].(string), "0x", "", 1)
				trueString := utils.GetTrueString(replace)
				if len(trueString)>0 {
					//d, err2 := strconv.ParseInt(trueString, 16, 64)
					//fmt.Println(m["transactionHash"])
					//fmt.Println(err2)
					//fmt.Println(d)
					var b big.Int
					setString, b2 := b.SetString(trueString, 16)
					if !b2 {
						return nil
					}

					//block, _ := b.SetString(strings.Replace(m["blockNumber"].(string), "0x", "", 1), 16)
					block,_ :=strconv.ParseInt(strings.Replace(m["blockNumber"].(string), "0x", "", 1),16,64)

					if block>blockNumber{
					blockNumber = block
					}
					//var transferData float32
					transferData :=float32(0)
					if chain==0{
						transferData, _ = BigIntToFloat(setString)
					}
					if chain==1 {
						transferData, _ = BigIntToFloat6(setString)
					}

					topics := m["topics"].([]interface{})
					sendAddress :="0x"+utils.GetTrueString(strings.Replace(topics[1].(string), "0x", "", 1))

					data1 :=map[string]float32{sendAddress:transferData}
					tmp[m["transactionHash"].(string)] = data1
				}
			}
			res["blockNumber"]=blockNumber
			res["data"] = tmp
			return res
		}
	}
	return nil
}

func (r *Request)TronGet()map[string]interface{} {
	resp, err := http.Get(r.Url)
	if err != nil {
		log.Printf("request err:%s",err.Error())
		return nil
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var value interface{}
	err1 := json.Unmarshal(body, &value)
	if err1!=nil {
		log.Printf("request err:%s",err1.Error())
		return nil
	}
	data := value.(map[string]interface{})
	//fmt.Println(data)
	if data["total"]=="0" {
		return nil
	}
	//var mp map[string]interface{}
	mp := make(map[string]interface{},100)

	tokenTransfers :=data["token_transfers"].([]interface{})
	for _,tf := range tokenTransfers{
		tf1 := tf.(map[string]interface{})
		if tf1["to_address"]!="TWYAFBF7ChJyGJGF8PcTUtVRdSuyXUgSmM"||tf1["quant"]==0||
			tf1["contract_address"]!="TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"||tf1["confirmed"]==false{
			continue
		}
		//var m map[string]interface{}
		m := make(map[string]interface{},1)
		quant, _ := strconv.ParseFloat(tf1["quant"].(string), 64)

		m[tf1["from_address"].(string)] = float32(quant)/float32(1000000)
		mp[tf1["transaction_id"].(string)] = m
	}
	return mp
}
