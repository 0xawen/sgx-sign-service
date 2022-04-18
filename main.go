package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/// 此应用放在sgx 运行
/// 需要做好数据冗余处理

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
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "ok",
			"data": "sign serve ok",
		})
	})

	// 创建地址
	r.GET("/create", func(c *gin.Context) {
		// 创建地址
		addr, err := CreateXuperAccount()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "error",
				"data": nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg":  "ok",
			"data": addr,
		})
	})

	// 签名
	r.POST("/sign", func(c *gin.Context) {
		// 传入需要地址和数据
		paramters := struct {
			Address string `json:"address"`
			Msg     string `json:"msg"`
		}{}
		err := c.ShouldBindJSON(&paramters)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "paramters error",
				"data": nil,
			})
			return
		}
		// 校验参数
		if paramters.Address == "" || paramters.Msg == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "paramters error",
				"data": nil,
			})
			return
		}
		// 签名
		signserve := NewXuperchainAccount(paramters.Address)
		sign, err := signserve.Sign([]byte(paramters.Msg))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg":  "sign error",
				"data": nil,
			})
			return
		}

		// todo 处理 sign
		c.JSON(http.StatusOK, gin.H{
			"msg":  "sign success",
			"data": sign,
		})
	})

	// 验证
	r.POST("/verify", func(c *gin.Context) {
		// sign  and  msg
		paramters := struct {
			Address string `json:"address"`
			Sign    []byte `json:"sign"`
			Msg     string `json:"msg"`
		}{}
		err := c.ShouldBindJSON(&paramters)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "paramters error",
				"data": nil,
			})
			return
		}
		if paramters.Address == "" || paramters.Sign == nil || paramters.Msg == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "paramters error",
				"data": nil,
			})
			return
		}
		// 校验
		signserve := NewXuperchainAccount(paramters.Address)
		result, err := signserve.verify(paramters.Sign, []byte(paramters.Msg))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg":  "verify error",
				"data": nil,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":  "verify success",
			"data": result,
		})
	})

	//监听端口默认为8080
	r.Run(":8080")
}
