package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"webtry/common"
	"webtry/routes"
)

/**
 * @Author: lbh
 * @Date: 2021/4/9
 * @Description:
 */


func main(){

	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	routes.CollectRoute(r)

	r.Run()


}

