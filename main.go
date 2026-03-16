// main.go
package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"lab-cuantitativo/router"
)

func main() {
	// Modo producción cuando se despliega en Render
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := router.Setup()
	r.Run(":" + port)
}