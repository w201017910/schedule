package route

import (
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
	c.HTML(http.StatusOK, "addclass.html", gin.H{"college": database.AllCollege(), "room": database.AllRoom()})
}
func ClassMange(c *gin.Context) {
	c.HTML(http.StatusOK, "class.html", gin.H{"class": database.AllClass(), "room": database.AllRoom(), "college": database.AllCollege()})
}
