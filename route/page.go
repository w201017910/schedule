package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"schedule/database"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "page-login.html", gin.H{})
}
func SchoolPage(c *gin.Context) {
	c.HTML(http.StatusOK, "school-class-information.html", gin.H{})
}
func TablePage(c *gin.Context) {

	c.HTML(http.StatusOK, "table-jsgrid.html", gin.H{"teacher": database.AllTeacher()})
}
func RoomPage(c *gin.Context) {
	c.HTML(http.StatusOK, "addroom.html", gin.H{"room": database.AllCollege()})
}
func Room(c *gin.Context) {
	c.HTML(http.StatusOK, "room.html", gin.H{"room": database.AllRoom(), "college": database.AllCollege()})
}
func ClassPage(c *gin.Context) {
	c.HTML(http.StatusOK, "addclass.html", gin.H{"college": database.AllCollege()})
}
func ClassMange(c *gin.Context) {
	c.HTML(http.StatusOK, "class.html", gin.H{"class": database.AllClass(), "room": database.AllRoom(), "college": database.AllCollege()})
}
func CurriculumPage(c *gin.Context) {
	c.HTML(http.StatusOK, "curriculum.html", gin.H{"teacher": database.AllTeacher(), "room": database.AllRoom()})
}
func CoursePage(c *gin.Context) {
	Type := c.Query("type")
	id := c.Query("id")
	year := c.Query("year")
	var course [16][20]database.Course
	fmt.Println(Type)
	fmt.Println(id)
	fmt.Println(year)
	if Type == "教室" {
		course = database.CreateByRoom(id, year)
		fmt.Println("教室")
	} else if Type == "教师" {
		course = database.CreateByTeacher(id, year)
		fmt.Println("教师")
	} else if Type == "班级" {
		course = database.CreateByClass(id, year)
		fmt.Println("班级")
	}

	c.HTML(http.StatusOK, "course.html", gin.H{"course": course})
}
func SelectPage(c *gin.Context) {
	c.HTML(http.StatusOK, "select.html", gin.H{})
}
