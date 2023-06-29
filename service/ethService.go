package service

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"pugg/config"
	"pugg/token"
	"pugg/utils"
	"strconv"
)

var Conn  *ethclient.Client
var EthConn *ethclient.Client
var TronConn *ethclient.Client

var UsdtToken *token.Usdt
var DofToken * token.DOF
func InitDial()  {
	dial, err := ethclient.Dial(config.Cfg.Ethereum.Network)
	if err != nil{
		log.Printf("bsc connect eth fail because:%s", err)
	}
	Conn = dial

	etDial, err1 := ethclient.Dial(config.Cfg.Ethereum.EthNetwork)
	if err1 != nil{
		log.Printf("eth connect eth fail because:%s", err)
	}
	EthConn = etDial

	TronDial, err2 := ethclient.Dial(config.Cfg.Ethereum.EthNetwork)
	if err2 != nil{
		log.Printf("tron connect eth fail because:%s", err)
	}
	TronConn = TronDial

	usdt, _ := token.NewUsdt(common.HexToAddress(config.Cfg.Ethereum.UsdtAddress), dial)
	UsdtToken = usdt

	dof,_:=token.NewDOF(common.HexToAddress(config.Cfg.Ethereum.DofAddress),dial)
	DofToken = dof
}
//砍去10位
func bigIntToInt10(bigint *big.Int) (int64,error){
	var in int64
	s := bigint.String()
	if len(s)<=10 {
		return int64(0),nil
	}
	s1 := s[0 : len(s)-10]
	atoi, err := strconv.ParseInt(s1,10,64)
	if err !=nil {
		return in,err
	}
	return atoi,nil
}

//砍去
func BigIntToFloat(bigint *big.Int) (float32,error){
	s := bigint.String()
	if len(s)<=10 {
		return float32(0),nil
	}
	s1 := s[0 : len(s)-10]
	atoi, err := strconv.ParseInt(s1,10,64)
	if err !=nil {
		return float32(0),err
	}
	return  float32(atoi)/100000000 ,nil
}
func BigIntToFloat6(bigint *big.Int) (float32,error) {
	s:=bigint.String()
	atoi, err := strconv.ParseInt(s,10,64)
	if err!=nil {
		return 0,err
	}
	return float32(atoi)/1000000,nil
}


//直接是int
func bigIntToInt(bigint *big.Int) (int64,error){
	var in int64
	s := bigint.String()
	atoi, err := strconv.ParseInt(s,10,64)
	if err !=nil {
		return in,err
	}
	return atoi,nil
}
//直接是砍去18为int类型
func bigIntToInt18(bigint *big.Int) (int64,error){
	var in int64
	s := bigint.String()
	if len(s)<=18 {
		return int64(0),nil
	}
	s1 := s[0 : len(s)-18]
	atoi, err := strconv.ParseInt(s1,10,64)
	if err !=nil {
		return in,err
	}
	return atoi,nil
}

//生成钱包公私钥
func GenerateKey()  {
	key, err := crypto.GenerateKey()
	if err!=nil {
		return
	}
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	fmt.Printf("钱包地址是:[%v]\n",address)
	toString := hex.EncodeToString(key.D.Bytes())
	fmt.Printf("私钥是：[%v]\n",toString)
}



//这是把int64末尾添加8位0转成bigint
func intTobigInt10(in int64) *big.Int {


	inString := strconv.FormatInt(in, 10)
	newString := inString+"0000000000"
	bigint, _ := new(big.Int).SetString(newString, 10)
	return bigint
}
//这是把int64转成bigint
func intTobigInt18(in int64) *big.Int {

	inString := strconv.FormatInt(in, 10)
	newString := inString+"000000000000000000"
	bigint, _ := new(big.Int).SetString(newString, 10)
	return bigint
}
//这是把float32转成bigint
func floatTobigInt(in float32) *big.Int {
	i := in * float32(100000000)
	return intTobigInt10(int64(i))
}



//func getTransactOpts(storePath string) (*bind.TransactOpts ,error) {
//	privateKey, err := crypto.HexToECDSA(storePath)
//	if err !=nil {
//		log.Fatalln(err)
//	}
//	publicKey := privateKey.Public()
//	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//	if !ok {
//		log.Println(err)
//		return nil,errors.New("has got treasure")
//	}
//	crypto.PubkeyToAddress(*publicKeyECDSA)
//	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(config.Cfg.Ethereum.ChainId)))
//	if err !=nil {
//		log.Println(err)
//		return nil,errors.New("privatekey wrong")
//	}
//	return auth ,nil
//}
func getTransactOpts(keyStore string ) (*bind.TransactOpts ,error) {
	privateKey, err := crypto.HexToECDSA(keyStore)
	if err !=nil {
		log.Fatalln(err)
		return nil,err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	crypto.PubkeyToAddress(*publicKeyECDSA)
	if !ok {
		log.Println(err)
		return nil,errors.New("has got treasure")
	}
	crypto.PubkeyToAddress(*publicKeyECDSA)

	auth := bind.NewKeyedTransactor(privateKey)
	if err !=nil {
		log.Println(err)
		return nil,errors.New("privatekey wrong")
	}
	return auth ,nil
}

//查询是否转账

func CheckHash(tradHash string) (bool,error)  {
	hash, _, err := Conn.TransactionByHash(context.Background(), common.HexToHash(tradHash))
	if err !=nil {
		return false,err
	}
	if hash.To()!=nil {
		log.Printf("checkHash:%s", hash.To())
		//log.Printf("type:%d",hash.Type())
		//fmt.Println(hash.AccessList())
		fmt.Println(hash.Data())
		toString := hex.EncodeToString(hash.Data())
		log.Printf("input data:%s", toString)
	}
	return true,nil
}
//查询是否转账U
const usdtMethodId = "a9059cbb"
const usdtAddress = "0x55d398326f99059ff775485246999027b3197955"
func CheckTransferU(tradHash string,toAddress string,amount float32) (bool,error)  {
	hash, _, err := Conn.TransactionByHash(context.Background(), common.HexToHash(tradHash))
	if err !=nil {
		log.Printf("err:%s",err.Error())
		return false,err
	}
	data := hex.EncodeToString(hash.Data())
	log.Printf("input data:%s",data)
	log.Println(data)

	if hash.To()!=nil && hash.Data()!=nil && len(data)>=136{
		log.Printf("checkHash:%s",hash.To().String())
		//data := hex.EncodeToString(hash.Data())
		method := data[0:8]

		to := data[8:72]
		a1 := data[72:136]
		trueString := "0x0"+utils.GetTrueString(to)

		setString, b := new(big.Int).SetString(utils.GetTrueString(a1), 16)
		if !b {
			return false,errors.New("trans err")
		}
		log.Println(common.HexToAddress(trueString)==common.HexToAddress(toAddress))
		log.Println(common.HexToAddress(trueString))
		log.Println(common.HexToAddress(toAddress))
		if method == usdtMethodId && common.HexToAddress(trueString)==common.HexToAddress(toAddress) &&  setString.String()== floatTobigInt(amount).String(){
			return true,nil
		}
		log.Printf("input data:%s",data)
	}
	return false,errors.New("err")
}


const ethUsdtMethodId = "a9059cbb"
const ethUsdtAddress = "0xdac17f958d2ee523a2206206994597c13d831ec7"
func EthCheckTransferU(tradHash string,toAddress string,amount float32) (bool,error)  {
	hash, _, err := EthConn.TransactionByHash(context.Background(), common.HexToHash(tradHash))
	if err !=nil {
		log.Printf("err:%s",err.Error())
		return false,err
	}
	data := hex.EncodeToString(hash.Data())
	log.Printf("input data:%s",data)
	if hash.To()!=nil && hash.Data()!=nil && len(data)>=136{
		log.Printf("checkHash:%s",hash.To())
		//data := hex.EncodeToString(hash.Data())
		method := data[0:8]

		to := data[8:72]
		a1 := data[72:136]
		trueString := "0x0"+utils.GetTrueString(to)

		setString, b := new(big.Int).SetString(utils.GetTrueString(a1), 16)
		if !b {
			return false,errors.New("trans err")
		}
		if method == ethUsdtMethodId && common.HexToAddress(trueString)==common.HexToAddress(toAddress) &&  setString.String()==new(big.Int).SetInt64(int64(amount*1000000)).String(){
			return true,nil
		}
		log.Printf("input data:%s",data)
	}
	return false,errors.New("err")
}


func GetNonce()int64  {
	at, _ := Conn.PendingNonceAt(context.Background(), common.HexToAddress(config.Cfg.Ethereum.KeyPath))
	return int64(at)
}

func UsdtTransferAmount(address string,amount float32) bool {
	auth, err := getTransactOpts(config.Cfg.Ethereum.KeyStore)
	if err !=nil{
		return false
	}
	once, err1 := UsdtToken.Transfer(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer},common.HexToAddress(address),floatTobigInt(amount))

	log.Printf("transfer begin address:%s,amount:%f",address,amount)

	if err1 !=nil {
		log.Printf("err1%s",err1.Error())
		return false
	}
	_, err2 := bind.WaitMined(context.Background(), Conn, once)
	if err2!=nil {
		log.Printf("err1%s",err2.Error())
		return false
	}
	log.Printf("transfer address:%s,amount:%f",address,amount)
	return true
}


func UsdtTransferAmountSync(address string,amount float32,nonce int64) bool {
	auth, err := getTransactOpts(config.Cfg.Ethereum.KeyStore)
	if err !=nil{
		return false
	}
	once, err1 := UsdtToken.Transfer(&bind.TransactOpts{
		From:   auth.From,
		Nonce: new(big.Int).SetInt64(nonce),
		Signer: auth.Signer},common.HexToAddress(address),floatTobigInt(amount))
	log.Printf("transfer begin address:%s,amount:%f",address,amount)
	if err1 !=nil {
		log.Printf("err1%s",err1.Error())
		return false
	}
	_, err2 := bind.WaitMined(context.Background(), Conn, once)
	if err2!=nil {
		log.Printf("err1%s",err2.Error())
		return false
	}
	log.Printf("transfer address:%s,amount:%f",address,amount)
	return true

}

func DofMintTo(address string,amount float32)bool  {
	auth, err := getTransactOpts(config.Cfg.Ethereum.DofKey)
	if err !=nil{
		return false
	}
	once, err1 := DofToken.Mint(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer},common.HexToAddress(address),floatTobigInt(amount))
	log.Printf("dof mint begin address:%s,amount:%f",address,amount)
	if err1 !=nil {
		log.Printf("dof mint err1%s",err1.Error())
		return false
	}
	_, err2 := bind.WaitMined(context.Background(), Conn, once)
	if err2!=nil {
		log.Printf("dof mint err1%s",err2.Error())
		return false
	}
	log.Printf("dof mint address:%s,amount:%f",address,amount)
	return true
}
