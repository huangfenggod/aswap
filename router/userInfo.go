package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"pugg/language"
	"pugg/service"
	"strconv"
)

func getUserInfo(ctx *gin.Context)  {
	lang := ctx.GetHeader("accept-language")
	l :=language.GetLangPackage(lang)

	address := ctx.Query("address")
	invite := ctx.Query("invite")
		if len(address)==0 {
			ctx.JSON(http.StatusOK,ResponseUtil{Status: true,Code: 201,Msg: l.GetParaWrong()})
			return
		}
		getAddr := service.GetUserInfo(address,invite)
		ctx.JSON(http.StatusOK,ResponseUtil{Status: true,Code:200,Msg: l.GetSuccess(),Data: getAddr})
		return
}


func getFans(ctx *gin.Context)  {
	lang := ctx.GetHeader("accept-language")
	l :=language.GetLangPackage(lang)
	address := ctx.Query("address")
	pageNumStr := ctx.Query("pageNum")
	pageSizeStr := ctx.Query("pageSize")
	pageNum, err1 := strconv.Atoi(pageNumStr)
	pageSize, err2 := strconv.Atoi(pageSizeStr)
	if len(address)==0||err1!=nil||err2!=nil {
		ctx.JSON(http.StatusOK,ResponseUtil{Status: true,Code: 201,Msg: l.GetParaWrong()})
		return
	}
	fans, total := service.GetFans(address, pageNum, pageSize)
	ctx.JSON(http.StatusOK,ResponseUtil{Status: true,Code:200,Msg: l.GetSuccess(),Data: gin.H{"total":total,"data":fans}})

}

func notice(ctx *gin.Context)  {
	lang := ctx.GetHeader("accept-language")

	l :=language.GetLangPackage(lang)
	id := ctx.Query("id")
	i, _ := strconv.Atoi(id)
	notices := service.GetNotices(lang,i)
	ctx.JSON(http.StatusOK,ResponseUtil{Status: true,Code: 200,Msg: l.GetSuccess(),Data: notices})
	return
}

func images(c *gin.Context)  {
	fileName := c.Query("fileName")
	//打开文件
	_, errByOpenFile := os.Open("/usr/pic/" + fileName)
	//非空处理
	if  len(fileName)<1|| errByOpenFile != nil {
		/*c.JSON(http.StatusOK, gin.H{
		    "success": false,
		    "message": "失败",
		    "error":   "资源不存在",
		})*/
		c.Redirect(http.StatusFound, "/404")
		return
	}
	//c.Header("Content-Type", "application/octet-stream")
	//c.Header("Content-Disposition", "attachment; filename="+fileName)
	//c.Header("Content-Transfer-Encoding", "binary")
	c.File( "/usr/pic/" + fileName)
	//file, _ := ioutil.ReadFile("/usr/pic/" + fileName)
	//c.Writer.Write(file)
	return
}

func banner(c *gin.Context)  {
	var s [3]string
	s[0]="http://144.34.167.208:8084/api/images?fileName=dof1.jpg"
	s[1]="http://144.34.167.208:8084/api/images?fileName=dof2.jpg"
	//s[2]="https://api.dreamspaceos.com/api/images?fileName=dof3.jpg"
	s[2]="http://144.34.167.208:8084/api/images?fileName=dof4.jpg"
	c.JSON(http.StatusOK,ResponseUtil{Status: true,Code: 200,Msg: "success",Data: s})

}

func getPhoto(c *gin.Context)  {
	header :=c.GetHeader("accept-language")
	s :="http://144.34.167.208:8084/api/images?fileName=dof1.jpg"

	if header=="zh-cn" {
		c.JSON(http.StatusOK,ResponseUtil{Status: true,Code: 200,Msg: "success",Data: s})
	}else {
		c.JSON(http.StatusOK,ResponseUtil{Status: true,Code: 200,Msg: "success",Data: s})
	}
}
