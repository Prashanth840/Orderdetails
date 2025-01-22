package main

import (
	"orderdetails/csvfile"
	"orderdetails/data"
	"orderdetails/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	data.DbConnect()
	csvfile.StartCronJob()
	r := gin.Default()
	routes.Routes(r)
	r.Run(":9000")
}
