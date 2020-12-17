package api

import (
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/controllers"
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/procedures"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strings"
)

type Handler struct {
	Router *gin.Engine
	*controllers.ApiController
}

type HandlerClosure func(*Handler)

func NewApiHandler(db *procedures.DB) *Handler {
	router := gin.Default()
	// Cors Config
	// Only use CORS config in STANDALONE server mode e.g. server is NOT behind proxy handling CORS
	corsConfig := cors.DefaultConfig()
	if os.Getenv("CORS_ORIGINS") != "" {
		corsAllowedOrigins := strings.Split(os.Getenv("CORS_ORIGINS"), ",")
		log.Printf("Allowed Origins: %s", corsAllowedOrigins)
		corsConfig.AllowOrigins = corsAllowedOrigins
		corsConfig.AllowCredentials = true
		corsConfig.AllowWebSockets = true
		router.Use(cors.New(corsConfig))
	}

	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	apiController := controllers.NewApiController(db)

	return &Handler{
		Router:        router,
		ApiController: apiController,
	}
}

func (h *Handler) CreateEndpoints() {
	v1 := h.Router.Group("/api/v1")
	{
		// Public group
		public := v1.Group("/public")
		{
			public.GET("/version", func(context *gin.Context) {
				context.JSON(200, gin.H{"version": "v0.0.1"})
			})
		}

		// private group (protected with auth)
		private := v1.Group("/private")
		{
			camera := private.Group("/camera")
			{
				camera.POST("/model", h.CameraModelCreate)
				camera.GET("/model", h.CameraModelRead)

				camera.POST("", h.CameraCreate)
				camera.GET("", h.CameraRead)
			}

			network := private.Group("/network")
			{
				network.POST("/proxy", h.NetworkProxyCreate)
			}
		}
	}
}
