// router/router.go
package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"lab-cuantitativo/handlers"
)

// Setup configura todas las rutas y middlewares del servidor
func Setup() *gin.Engine {
	r := gin.Default()

	// ── Middleware CORS ──────────────────────────────────────
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{
			"http://localhost:5173",
			"http://localhost:3000",
			"https://mauricionoj.com",
		},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: false,
	}))

	// ── Health check ─────────────────────────────────────────
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "lab-cuantitativo-api",
		})
	})

	// ── Rutas del Lab Cuantitativo ───────────────────────────
	lab := r.Group("/api/lab")
	{
		lab.POST("/vpn-tir/",     handlers.VPNTIRHandler)
		lab.POST("/anualidades/", handlers.AnualidadesHandler)
		lab.POST("/simplex/",     handlers.SimplexHandler)
		lab.POST("/estadistica/", handlers.EstadisticaHandler)
		lab.POST("/newton/",      handlers.NewtonHandler)
		lab.POST("/integracion/", handlers.IntegracionHandler)
	}

	return r
}