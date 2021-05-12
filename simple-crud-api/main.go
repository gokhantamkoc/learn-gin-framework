package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", index)
	r.POST("/company/create", CreateCompany)
	r.POST("/company/update", UpdateCompany)
	r.GET("/company/list", ListCompanies)
	r.GET("/company/:id", GetCompany)
	r.Run()
}

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"appName": "Learning Gin Framework",
		"version": "v0.0.1",
		"message": "Hello Gin!",
	})
}
