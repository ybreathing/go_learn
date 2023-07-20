package main

import (
	"zzy/go-learn/common"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//mysql.DeregisterLocalFile()
	// 初始化数据库
	db := common.InitDB()
	// 延迟关闭db
	defer db.Close()

	r := gin.Default()
	r = CollectRouter(r)

	panic(r.Run()) // listen and serve on 0.0.0.0:8080

}
