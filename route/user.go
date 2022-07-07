package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"schedule/config"
	"schedule/database"
	"strings"
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

	err := database.ChangeTeacher(c.PostForm("Id"), c.PostForm("Name"), c.PostForm("Sex"), c.PostForm("Phone"), c.PostForm("Email"), c.PostForm("WorkId"))
	if err != nil {
		c.JSON(200, "修改失败")
		return
	}
	c.JSON(200, "修改成功")
}
func AddRoom(c *gin.Context) {
	name := c.PostForm("name")
	college := c.PostForm("college")
	number := c.PostForm("number")
	collegeId := strings.Split(college, "-")[0]
	err := database.AddRoom(name, number, collegeId)
	if err != nil {
		c.JSON(200, false)
		return
	}
	c.JSON(200, true)
}
func ChangeRoom(c *gin.Context) {
	id := c.PostForm("Id")
	name := c.PostForm("Name")
	number := c.PostForm("Number")
	CollegeName := c.PostForm("CollegeName")
	collegeId := strings.Split(CollegeName, "-")[0]
	err := database.ChangeRoom(id, name, number, collegeId)
	if err != nil {
		c.JSON(200, false)
		return
	}
	c.JSON(200, true)
}
func DelRoom(c *gin.Context) {
	id := c.PostForm("Id")
	err := database.DelRoom(id)
	if err != nil {
		c.JSON(200, false)
		return
	}
	c.JSON(200, true)
}
func AddClass(c *gin.Context) {
	name := c.PostForm("name")
	number := c.PostForm("number")
	college := c.PostForm("college")

	collegeId := strings.Split(college, "-")[0]

	err := database.AddClass(name, number, collegeId)
	if err != nil {
		c.JSON(200, false)
		return
	}
	c.JSON(200, true)
}
func ChangeClass(c *gin.Context) {
	id := c.PostForm("Id")
	name := c.PostForm("Name")
	number := c.PostForm("Number")
	college := c.PostForm("college")

	collegeId := strings.Split(college, "-")[0]

	err := database.ChangeClass(id, name, number, collegeId)
	if err != nil {
		c.JSON(200, false)
		return
	}
	c.JSON(200, true)
}
func DelClass(c *gin.Context) {
	id := c.PostForm("Id")
	err := database.DelClass(id)
	if err != nil {
		c.JSON(200, false)
		return
	}
	c.JSON(200, true)
}
func AddCurriculum(c *gin.Context) {
	ban := c.PostForm("ban")
	className := c.PostForm("className")
	name := c.PostForm("name")
	teacher := c.PostForm("teacher")
	timeLong := c.PostForm("timeLong")
	totalTime := c.PostForm("totalTime")
	year := c.PostForm("year")
	origin := c.PostForm("origin")
	room := c.PostForm("room")
	roomId := strings.Split(room, "-")[0]
	err := database.AddCurriculum(name, teacher, timeLong, totalTime, year, className, ban, origin, roomId)
	if err != nil {
		c.JSON(200, false)
		return
	}
	c.JSON(200, true)
}
