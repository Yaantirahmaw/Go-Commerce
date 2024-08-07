package main

import (
	"go-commerce/configs"
	"go-commerce/handlers"
	"go-commerce/middlewares"
	"go-commerce/migrations"

	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"

	"net/http"
	_ "net/http/pprof"
)

func main() {
	db := configs.ConnectDB()
	migrations.Migrate(db)
	
	router := gin.Default()

	// Rute CRUD Produk
	router.GET("/products", middlewares.AuthMiddleware(), handlers.ListProducts(db))
	router.GET("/products/:id", handlers.GetProduct(db))
	router.POST("/products", handlers.CreateProduct(db))
	router.PUT("/products/:id", handlers.UpdateProduct(db))
	router.DELETE("/products/:id", handlers.DeleteProduct(db))

	// Rute CRUD Product Category
	router.GET("/product-categories", handlers.ListProductCategories(db))
	router.GET("/product-categories/:id", handlers.GetProductCategory(db))
	router.POST("/product-categories", handlers.CreateProductCategory(db))
	router.PUT("/product-categories/:id", handlers.UpdateProductCategory(db))
	router.DELETE("/product-categories/:id", handlers.DeleteProductCategory(db))

	// Rute Transaksi
	router.POST("/transactions", handlers.CreateTransaction(db))
	router.GET("/transactions/:id", handlers.GetTransactionWithItems(db))

	// Rute Autentikasi
	router.POST("/login", handlers.Login(db))
	router.POST("/register", handlers.Register(db))

	// Rute Debug Profiling
	router.GET("/debug/pprof/*pprof", gin.WrapH(http.DefaultServeMux))

	router.Run(":5000")
}
