package routes

import (
	"go-api/controllers/ProdukController"

	"go-api/controllers/NewsController"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	router.GET("/api/produk", ProdukController.Index)
	router.GET("/api/produk/:id", ProdukController.Show)
	router.POST("/api/produk", ProdukController.Create)
	router.PUT("/api/produk/:id", ProdukController.Update)
	router.DELETE("/produk/", ProdukController.Delete)

	router.GET("/api/news", NewsController.Index)
	router.GET("/api/news/:id", NewsController.Show)
	router.POST("/api/news", NewsController.Create)
	router.PUT("/api/news/:id", NewsController.Update)
	router.DELETE("/news/", NewsController.Delete)

	router.Run("127.0.0.1:8080")
}
