package nosql

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"pugg/config"
	"pugg/pojo"
	"time"
)

var redisCli *redis.Client

func InitRedis()  {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Cfg.RedisNetwork,
		Password: config.Cfg.RedisPassword, // no password set
		DB:       0,  // use default DB
	})
	redisCli = rdb
}
type Std struct {
	Name string `json:"name" redis:"name"`
	Age int `json:"age" redis:"age"`
}

func (std *Std)MarJson() []byte {
	marshal, err := json.Marshal(std)
	fmt.Println(err)
	return marshal
}

//设置2小时的缓存
func SetUserCache(address string, info pojo.UserInfo)  {
	marshal, _ := json.Marshal(info)
	redisCli.Set(context.Background(),"dof:"+address,marshal,time.Hour*2)

}
//获取缓存内容
func GetUserCache(address string) (pojo.UserInfo,bool){
	var info pojo.UserInfo
	exists := redisCli.Exists(context.Background(), "dof:"+address)
	if exists.Val()==0 {
	return info,false
	}

	get := redisCli.Get(context.Background(), "dof:"+address)

	err := json.Unmarshal([]byte(get.Val()), &info)
	if err!=nil {
		return info,false
	}
	return info,true
}
//清除缓存
func DelCache(address string)  {
	redisCli.Del(context.Background(),"dof:"+address)
}
func DelRank(s string )  {
	redisCli.Del(context.Background(),s)
}



func DelAllCache(address []string)  {
	//addr := make([]string,0)
	for _,v:=range address{
		v="dof:"+v
		//addr = append(addr, v)
		redisCli.Del(context.Background(),v)
	}
}


func SetInRedis(address string,name string,last time.Duration)  {
	ctx := context.Background()
	redisCli.Set(ctx, "dof"+name+":"+address, 1, last*time.Second )
	log.Printf("address:%s :访问加锁",address)
}


func GetRedis(address string,name string) int {
	ctx := context.Background()
	v := redisCli.Get(ctx,"dof"+name+":"+address)
	v1,_ := v.Int()
	return v1
}


func RemoveRedis(address string,name string)  {
	ctx := context.Background()
	redisCli.Del(ctx, "dbtc:"+name+":"+address)
	log.Printf("address:%s :访问解锁锁",address)
}


func Lock(address string,name string) (int,error) {
	getRedis := GetRedis(address,name)
	log.Printf("redis锁操作，返回结果：%d",getRedis)
	if getRedis==1 {
		log.Printf("address:%s :访问 %s 频繁操作",address,name)
		return 0,errors.New("访问过于频繁稍后再试")
	}
	SetInRedis(address,name,60)
	return 0,nil
}

func TimeLock(address string,name string) (int,error) {
	getRedis := GetRedis(address,name)
	log.Printf("rediss锁操作，返回结果：%d",getRedis)
	if getRedis==1 {
		log.Printf("address:%s :访问 %s 频繁操作",address,name)
		return 0,errors.New("访问过于频繁稍后再试")
	}
	SetInRedis(address,name,9)
	return 0,nil
}


func UnLock(address string,name string)  {
	RemoveRedis(address,name)
}

func SDel(i int)  {
	var s string
	switch i {
	case 1:
		s = "morning"
		break
	case 2:
		s = "afternoon"
		break
	default:
		s = "morning"
	}
	redisCli.Del(context.Background(),s)
}

func SAdd(address string,i int)  {
	var s string
	switch i {
	case 1:
		s = "morning"
		break
	case 2:
		s = "afternoon"
		break
	default:
		s = "morning"
	}
	add := redisCli.SAdd(context.Background(), s,address)
	fmt.Println(add.Err())
}


//有序集合
func ZAdd(address string,score float32,i int)  {
	var s string
	switch i {
	case 1:
		s = "morning"
		break
	case 2:
		s = "afternoon"
		break
	case 3:
		s = "morning1"
		break
	case 4:
		s="afternoon1"
		break
	case 5:
		s="morning2"
		break
	case 6:
		s="afternoon2"
		break
	}
	add := redisCli.ZAdd(context.Background(), s, &redis.Z{Score: float64(score), Member: address})
	fmt.Println(add.Err())
}

//有序集合查询是否是成员
func ZIsMember(address string,i int) bool {
	var s string
	switch i {
	case 1:
		s = "morning"
		break
	case 2:
		s = "afternoon"
		break
	case 3:
		s = "morning1"
		break
	case 4:
		s="afternoon1"
		break
	case 5:
		s="morning2"
		break
	case 6:
		s="afternoon2"
		break
	default:
		s = "morning"
	}

	score := redisCli.ZScore(context.Background(), s, address)
	if score.Err()!=nil {
		return false
	}
	return true

}

func GetZRankNumber(i int) int {
	var s string
	switch i {
	case 1:
		s = "morning"
		break
	case 2:
		s = "afternoon"
		break
	case 3:
		s = "morning1"
		break
	case 4:
		s="afternoon1"
		break
	case 5:
		s="morning2"
		break
	case 6:
		s="afternoon2"
		break
	}
	card := redisCli.ZCard(context.Background(), s)
	return int(card.Val())
}

func GetZRankHead(number int,i int)[]string  {
	var s string
	switch i {
	case 1:
		s = "morning"
		break
	case 2:
		s = "afternoon"
		break
	case 3:
		s = "morning1"
		break
	case 4:
		s="afternoon1"
		break
	case 5:
		s="morning2"
		break
	case 6:
		s="afternoon2"
		break
	}
	score := redisCli.ZRevRange(context.Background(),s,0,int64(number)-1)

	return score.Val()
}
func GetZRankTail(number int,i int)[]string  {
	var s string
	switch i {
	case 1:
		s = "morning"
		break
	case 2:
		s = "afternoon"
		break
	case 3:
		s = "morning1"
		break
	case 4:
		s="afternoon1"
		break
	case 5:
		s="morning2"
		break
	case 6:
		s="afternoon2"
		break
	}

	score := redisCli.ZRange(context.Background(),s,0,int64(number))

	return score.Val()
}



func SIsMember(address string,i int)bool  {
	var s string
	switch i {
	case 1:
		s = "morning"
		break
	case 2:
		s = "afternoon"
		break
	default:
		s = "morning"
	}

	member := redisCli.SIsMember(context.Background(), s, address)
	return member.Val()
}

func GetMember(i int) []string {
	var s string
	switch i {
	case 1:
		s = "morning"
		break
	case 2:
		s = "afternoon"
		break
	case 3:
		s = "morning1"
		break
	case 4:
		s="afternoon1"
		break
	case 5:
		s="morning2"
		break
	case 6:
		s="afternoon2"
		break
	default:
		s = "morning"
	}

	fmt.Println(i)
	fmt.Println(s)
	members := redisCli.ZRange(context.Background(), s,0,-1)
	result, err := members.Result()
	fmt.Println(err)
	return result
}
