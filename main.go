package main

import "github.com/gin-gonic/gin"

func main() {
	Start("localhost:8080")
}
func Start(addr string) (err error) {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("./template/*")
	err = r.Run(addr)
	return err
}
