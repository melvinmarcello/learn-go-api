package NewsController

import (
	"go-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var news []models.News
	result := models.DB.Debug().Find(&news)

	if result.Error != nil {
		println("No news found")
		c.JSON(http.StatusNotFound, gin.H{"error": "No news found"})
		return
	} else {
		println("news fetched successfully")
		c.JSON(http.StatusOK, gin.H{"news": news})
	}
}

func Show(c *gin.Context) {
	var news models.News

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
