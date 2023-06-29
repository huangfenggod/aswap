package router

func ApiGroup() {
	group := GIN.Group("/api")
	group.GET("/getUserInfo", getUserInfo)  //获取用户信息,参数address，invite,是否弹出用户上一次查看记录
	group.GET("/getFans",getFans)  //查询直推粉丝 ,参数address，pageNum.pageSize
	group.GET("/banner",banner)  //获取轮播图
	group.GET("/getPhoto",getPhoto)  //获取尾图  accept-language zh-cn   en-us

	group.GET("/images",images) //获取图片位置  fileName=1.jpg   2.jpg   3.jpg
	group.POST("/charge",charge)  // address amount  transaction,chainId  为升级进行充值
	group.GET("/notice",notice)  //header  accept-language "en-US"zh-CN" 项目介绍，id=1.参与规则id=2

	//group.GET("/incomeRecord",incomeRecord)  //收益记录 参数address,pageNum,pageSize
	//group.POST("/cash",cash)  //提现dbtc, 参数address，amount(dbtc数量) ，id（1是提现直推，2团队奖励，3质押利息）
	//group.GET("/ystData",ystData)  //获取昨日数据
	//group.GET("/cashRecord",cashRecord)  //提现记录
}
