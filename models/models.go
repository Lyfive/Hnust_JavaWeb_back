/**
 @author: 李映飞
 @date:   2022/4/3
 @note:
**/
package models

import (
	"gorm.io/gorm"
	"log"
)
import "gorm.io/driver/mysql"

var db *gorm.DB
var err error

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/javaweb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("models : %v\n", err)
	}
}
