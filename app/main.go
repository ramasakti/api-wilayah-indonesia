package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "indonesia-api/docs"
)

// @title Indonesia Wilayah API
// @version 1.0
// @description API Wilayah Indonesia (Provinsi, Kabupaten, Kecamatan, Desa).
// @contact.name API Support
// @contact.url https://github.com/ramasakti/api-wilayah-indonesia
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api

func main() {
	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		dataDir = "../data"
	}

	store := NewDataStore(dataDir)
	h := &Handlers{Store: store}

	// Set Gin to release mode if not specified
	if os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Enable CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := r.Group("/api")
	{
		api.GET("/provinces.json", h.GetProvinces)
		api.GET("/provinces", h.GetProvinces)

		api.GET("/regencies/:id", h.GetRegencies)
		api.GET("/districts/:id", h.GetDistricts)
		api.GET("/villages/:id", h.GetVillages)
		api.GET("/province/:id", h.GetProvince)
		api.GET("/regency/:id", h.GetRegency)
		api.GET("/district/:id", h.GetDistrict)
		api.GET("/village/:id", h.GetVillage)
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
