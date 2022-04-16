package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/// 此应用放在sgx 运行
/// todo 需要做好数据冗余处理

func init() {
	// 初始化数据库dbsoucename
	dbsoucename := "xuperchain"
	db, err := InitDB(dbsoucename)
	if err != nil {
		panic("init db error")
	}
	GDB = db
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})

	// 创建地址
	r.POST("/create", func(c *gin.Context) {
		// 创建地址
		// todo
		c.String(http.StatusOK, "hello word")
	})

	// 签名
	r.POST("/sign", func(c *gin.Context) {
		// todo
		c.String(http.StatusOK, "hello word")
	})

	//监听端口默认为8080
	r.Run(":8080")
}
