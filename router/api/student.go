/**
 @author: 李映飞
 @date:   2022/4/3
 @note:
**/
package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"javaweb/models"
	"log"
	"net/http"
	"strconv"
)

func Admin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"user_name": "LyFive",
			"user_type": "院级管理员",
			"user_num":  1169442146,
			"college":   "计算机科学与工程学院",
			"isadmin":   true,
		},
		"message": gin.H{
			"student_count": 1245,
			"teacher_count": 1245,
			"course_count":  52,
			"courses_count": 62,
			"class_count":   124,
			"major_count":   10,
		},
	})
}
func Find(c *gin.Context) {
	var stu models.Student
	stu.Account = c.Query("account")
	stu.Name = c.Query("name")
	stu.Major = c.Query("major")
	stu.Class = c.Query("classname")

	ops := models.GetOperator(stu)
	c.JSON(http.StatusOK, gin.H{
		"data": ops,
	})
}
func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Fatalf("Update have err : %v\n", err)
	}
	fmt.Println(id)
	op := models.UpdateOperatorForId(uint(uint64(id)), c)
	c.JSON(http.StatusOK, gin.H{
		"data": op,
	})
}

func Add(c *gin.Context) {
	var err error
	var op models.Operator
	err = c.BindJSON(&op)
	fmt.Println(op)
	if err != nil {
		log.Fatalf("/add BindJSON err : %v\n", err)
	}
	models.AddOperator(op)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Success",
	})
}

func Delete(c *gin.Context) {
	ids := c.QueryArray("ids")
	fmt.Println(ids)
	msg := "Error"
	for _, id := range ids {
		x, _ := strconv.Atoi(id)
		msg = models.DeleteOperator(uint(uint64(x)))
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func Register(c *gin.Context) {
	var err error
	var stu models.Student
	err = c.BindJSON(&stu)
	fmt.Println(stu)
	if err != nil {
		log.Fatalf("/register BindJSON err : %v \n", err)
	}
	models.AddStudent(stu)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Success",
	})
}
