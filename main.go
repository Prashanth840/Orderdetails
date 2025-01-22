package main

import (
	"orderdetails/data"
	"orderdetails/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	data.DbConnect()
	r := gin.Default()
	routes.Routes(r)
	r.Run(":9000")
}
