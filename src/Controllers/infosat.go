package controllers

import (
	"net/http"

	config "../Config"
	models "../Models"
	"github.com/gin-gonic/gin"
)

//Benefits Get all benefits to user
func Benefits(c *gin.Context) {
	var benefits []models.Deducciones
	var db = config.GetDB()
	db.Find(&benefits)
	c.JSON(http.StatusOK, &benefits)
}

//ObInfo get a list of obligations for the user
func ObInfo(c *gin.Context) {
	var obligations []models.Obligaciones
	var db = config.GetDB()
	db.Find(&obligations)
	c.JSON(http.StatusOK, &obligations)
}

//CreateBenefit for add new deduction in to db
func CreateBenefit(c *gin.Context) {
	var benefit models.Deducciones
	var db = config.GetDB()
	err := c.BindJSON(&benefit)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&benefit)
	c.JSON(http.StatusOK, &benefit)
}

// CreateOb for add new obligation in to db
func CreateOb(c *gin.Context) {
	var obligation models.Obligaciones
	var db = config.GetDB()
	err := c.BindJSON(&obligation)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&obligation)
	c.JSON(http.StatusOK, &obligation)
}

// UpdateBenefit update deducciones
func UpdateBenefit(c *gin.Context) {
	id := c.Param("id")
	var benefit models.Deducciones
	var db = config.GetDB()
	err := db.Where("id = ?", id).First(&benefit).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&benefit)
	db.Save(&benefit)
	c.JSON(http.StatusOK, &benefit)
}

//UpdateOb for Obligaciones
func UpdateOb(c *gin.Context) {
	id := c.Param("id")
	var obligation models.Obligaciones
	var db = config.GetDB()
	err := db.Where("id = ?", id).First(&obligation).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&obligation)
	db.Save(&obligation)
	c.JSON(http.StatusOK, &obligation)
}

//DeleteBenefit for deleting deducciones
func DeleteBenefit(c *gin.Context) {
	id := c.Param("id")
	var benefit models.Deducciones
	var db = config.GetDB()

	if err := db.Where("id = ?", id).First(&benefit).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	db.Delete(&benefit)
}

//DeleteOb for deleting Obligaciones
func DeleteOb(c *gin.Context) {
	id := c.Param("id")
	var obligation models.Obligaciones
	db := config.GetDB()
	if err := db.Where("id = ?", id).First(&obligation).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	db.Delete(&obligation)
}
