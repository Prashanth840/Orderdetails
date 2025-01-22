package controller

import (
	"log"
	"net/http"
	"orderdetails/csvfile"
	"orderdetails/models"
	"orderdetails/repository"

	"github.com/gin-gonic/gin"
)

func TotalRevenue(c *gin.Context) {
	var input models.Input
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := repository.TotalRevenue(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"TotalRevenue": result})
}

func TotalRevenuebyproduct(c *gin.Context) {
	var input models.Input
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := repository.TotalRevenuebyproduct(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"TotalRevenuebyproduct": result})
}

func TotalRevenueByCategory(c *gin.Context) {
	var input models.Input
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := repository.TotalRevenueByCategory(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"TotalRevenuebyproduct": result})
}

func RefreshData(c *gin.Context) {
	log.Println("Starting data refresh...")
	csvfile.Loadcsvdata("sales_data.csv")
	c.JSON(http.StatusOK, gin.H{"message": "Data refresh completed successfully."})
}
