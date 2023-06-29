package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pugg/language"
	"pugg/service"
	"strconv"
)


type param struct {
	Address string `json:"address"`
	Amount string `json:"amount"`
	Transaction string `json:"transaction"`
	Invite string `json:"invite"`
	ChainId string `json:"chainId"`
	District int `json:"district"`
}
type param1 struct {
	Address string `json:"address"`
	//Amount string `json:"amount"`
	Id int `json:"id"`
}
type param2 struct {
	Address string `json:"address"`
	Amount string `json:"amount"`
	Transaction string `json:"transcation"`
}




func charge(ctx *gin.Context)  {
	lang := ctx.GetHeader("accept-language")
	l :=language.GetLangPackage(lang)
	var para param
	ctx.BindJSON(&para)
	amount, err1:= strconv.ParseFloat(para.Amount,64)
	if len(para.Address)==0||amount<=0||len(para.Transaction)<1||err1!=nil {
		ctx.JSON(http.StatusOK,ResponseUtil{Status: false,Code:201,Msg: l.GetParaWrong()})
		return
	}
	_,err := service.Charge(para.Address, float32(amount),para.Transaction,para.ChainId)
	if err==nil {
		ctx.JSON(http.StatusOK,ResponseUtil{Status: true,Code:200,Msg: l.GetSuccess()})
	}else {
		ctx.JSON(http.StatusOK,ResponseUtil{Status: false,Code:201,Msg: err.Error()})
	}
	return
}












////获取粉丝请求
//func getFans(ctx *gin.Context)  {
//	address := ctx.Query("address")
//	pageNumStr := ctx.Query("pageNum")
//	pageSizeStr := ctx.Query("pageSize")
//	pageNum, err1 := strconv.Atoi(pageNumStr)
//	pageSize, err2 := strconv.Atoi(pageSizeStr)
//	if len(address)==0||err1!=nil||err2!=nil {
//		ctx.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.ENGLISH_PARAMS_WRONG})
//		return
//	}
//	fans, total := service.GetFans(address, pageNum, pageSize)
//	ctx.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "success",Data: gin.H{"total":total,"data":fans}})
//}
