package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"schedule/config"
	"schedule/database"
)

func Login(c *gin.Context) {
	account := c.PostForm("account")
	passwd := c.PostForm("passwd")
	fmt.Println(account)
	fmt.Println(passwd)
	if account == config.Account && passwd == config.PassWord {
		c.SetCookie("name", account, 1000, "/", "localhost", false, true)
		fmt.Println(true)
		c.JSON(200, true)
		return
	}

	c.JSON(200, false)
	return
}
func AddTeacher(c *gin.Context) {
	name := c.PostForm("name")
	sex := c.PostForm("sex")
	workID := c.PostForm("workID")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	fmt.Println(name, sex, workID, phone, email)
	judge := database.TeacherAdd(name, sex, workID, phone, email)
	if judge {
		c.JSON(200, true)
		return
	}
	c.JSON(200, false)
	return
}
func AllTeacher(c *gin.Context) {
	i := database.AllTeacher()

	c.JSON(200, i)
}
func ChangeTeachers(c *gin.Context) {
	fmt.Println(c.PostForm("Id"), c.PostForm("Name"), c.PostForm("Sex"), c.PostForm("Phone"), c.PostForm("Email"), c.PostForm("WorkId"))
	err := database.ChangeTeacher(c.PostForm("Id"), c.PostForm("Name"), c.PostForm("Sex"), c.PostForm("Phone"), c.PostForm("Email"), c.PostForm("WorkId"))
	if err != nil {
		c.JSON(200, "修改失败")
		return
	}
	c.JSON(200, "修改成功")
}
