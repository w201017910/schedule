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
