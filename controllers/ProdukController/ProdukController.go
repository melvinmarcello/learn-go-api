package ProdukController

import (
	"go-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var product []models.Product
	result := models.DB.Debug().Find(&product)

	if result.Error != nil {
		println("No Product found")
		c.JSON(http.StatusNotFound, gin.H{"error": "No Product found"})
		return
	} else {
		println("Product fetched successfully")
		c.JSON(http.StatusOK, gin.H{"product": product})
	}
}

func Show(c *gin.Context) {
	var news models.Product

	id := c.Param("id")

	if err := models.DB.First(&news, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"news": news})
}
func Update(c *gin.Context) {

}
func Create(c *gin.Context) {

}
func Delete(c *gin.Context) {

}
