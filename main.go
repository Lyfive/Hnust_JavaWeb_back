/**
 @author: 李映飞
 @date:   2022/3/31
 @note:
**/
package main

import (
	"javaweb/router"
)

func main() {
	// 路由初始化
	r := router.InitRouter()

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
