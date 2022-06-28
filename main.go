package main

import (
	"github.com/gin-gonic/gin"
	"schedule/database"
	"schedule/route"
)

func main() {
	database.Init()
	Start("localhost:8080")
}
func Start(addr string) (err error) {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("./template/*")
	r.GET("/page-login.html", route.LoginPage)
	r.GET("/", route.SchoolPage)
	r.GET("/table-jsgrid.html", route.TablePage)
	r.GET("/addroom.html", route.RoomPage)
	r.GET("/room.html", route.Room)
	r.GET("/addclass.html", route.ClassPage)
	r.POST("/login", route.Login)
	r.POST("/addTeacher", route.AddTeacher)
	r.POST("/allTeacher", route.AllTeacher)
	r.POST("/changeTeacher", route.ChangeTeachers)
	r.POST("/addRoom", route.AddRoom)
	r.POST("/changeRoom", route.ChangeRoom)
	r.POST("/delRoom", route.DelRoom)
	err = r.Run(addr)
	return err
}
