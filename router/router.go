/**
 @author: 李映飞
 @date:   2022/4/3
 @note:
**/
package router

import (
	"github.com/gin-gonic/gin"
	"javaweb/router/api"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())

	// 管理员面板
	r.GET("/",api.Admin)
	// 学生列表组
	StudentList := r.Group("/student")
	{
		// 查询学生信息
		StudentList.GET("/find", api.Find)
		// 更新信息
		StudentList.PUT("/update", api.Update)
		// 添加信息
		StudentList.POST("/add", api.Add)
		// 删除信息
		StudentList.DELETE("/delete", api.Delete)
	}
	// 学生账号组
	Account := r.Group("/account")
	{
		// 注册
		Account.POST("/register", api.Register)
	}
	return r
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
