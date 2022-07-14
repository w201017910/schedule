package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"schedule/database"
	"schedule/route"
)

func main() {
	database.Init()
	Start("localhost:8080")
}
func Start(addr string) (err error) {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"name": name,
		"add":  add,
	})
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("./template/*")

	r.GET("/page-login.html", route.LoginPage)
	r.GET("/", route.SchoolPage)
	r.GET("/table-jsgrid.html", route.TablePage)
	r.GET("/addroom.html", route.RoomPage)
	r.GET("/room.html", route.Room)
	r.GET("/addclass.html", route.ClassPage)
	r.GET("/class.html", route.ClassMange)
	r.GET("/curriculum.html", route.CurriculumPage)
	r.GET("course.html", route.CoursePage)
	r.POST("/login", route.Login)
	r.POST("/addTeacher", route.AddTeacher)
	r.POST("/allTeacher", route.AllTeacher)
	r.POST("/changeTeacher", route.ChangeTeachers)
	r.POST("/addRoom", route.AddRoom)
	r.POST("/changeRoom", route.ChangeRoom)
	r.POST("/delRoom", route.DelRoom)
	r.POST("/addClass", route.AddClass)
	r.POST("/changeClass", route.ChangeClass)
	r.POST("/delClass", route.DelClass)
	r.POST("/addCurriculum", route.AddCurriculum)
	err = r.Run(addr)
	return err
}
func name(href [20]database.Course, i int) string {
	if href[i].Exist {

		return href[i].Name + "  " + database.QueryTeacherName(href[i].TeacherId) + "  " + database.QueryRoomName(href[i].RoomId) + "  " + database.QueryClassName(href[i].ClassId)
	}
	return "无"
}
func add(a int, b int) int {
	return a + b
}
