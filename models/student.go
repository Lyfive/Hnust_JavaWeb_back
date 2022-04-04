/**
 @author: 李映飞
 @date:   2022/4/3
 @note:
**/
package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Student struct {
	// 检索集
	Account string `gorm:"primaryKey" json:"account"'`
	Name    string `json:"name"`
	Major   string `json:"major"`
	Class   string `json:"classname"`
}

type Operator struct {
	// 结果集
	ID       uint   `gorm:"primaryKey" json:"id"`
	Account  string `json:"account"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Action   string `json:"operation"`
	School   string `json:"school"`
}

func init() {
	err = db.AutoMigrate(&Student{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&Operator{})
	if err != nil {
		return
	}
}
// 学生信息操作

func AddOperator(op Operator) {
	db.Create(&op)
}

func DeleteOperator(id uint) string {
	res := db.Delete(&Operator{}, id)
	if res.Error != nil {
		return "Note Found"
	} else {
		return "Success"
	}
}

func GetOperator(student Student) (operators []Operator) {
	if student.Account == "" {
		student.Account = "%"
	}
	if student.Name == "" {
		student.Name = "%"
	}
	if student.Class == "" {
		student.Class = "%"
	}
	if student.Major == "" {
		student.Major = "%"
	}
	db.Where("account in (?)",
		db.Table("students").Select("account").
			Where("account like ? AND name like ? AND class like ? AND major like ?", student.Account, student.Name, student.Class, student.Major)).
		Find(&operators)
	fmt.Println(operators)
	return
}

func UpdateOperatorForId(id uint, c *gin.Context) (op Operator) {
	db.First(&op, id)
	if c.Query("account") != "" {
		op.Account = c.Query("account")
	}
	if c.Query("name") != "" {
		op.Name = c.Query("name")
	}
	if c.Query("password") != "" {
		op.Password = c.Query("password")
	}
	if c.Query("operation") != "" {
		op.Action = c.Query("operation")
	}
	if c.Query("school") != "" {
		op.School = c.Query("school")
	}
	db.Save(&op)
	return
}

// 学生账号操作

func AddStudent(stu Student){
	db.Create(&stu)
}
