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
			// swagger:route GET /public/version server version
			//
			// Get the current server version
			//
			// The version follows the semVer structure e.g. v0.0.1
			//
			// Consumes:
			// - application/json
			// Produces:
			// - application/json
			// Schemes: http, https
			// Deprecated: false
			// Responses:
			//	default: versionResponse
			//	200: versionResponse
			public.GET("/version", h.Version)
		}

		// private group (protected with auth)
		private := v1.Group("/private")
		{
			camera := private.Group("/camera")
			{
				// swagger:route POST /private/camera/model create a camera model
				//
				// Create a Camera Model
				//
				// Create a Camera Model such as RLC-411WS.
				// This will also return the newly created model.
				//
				// Consumes:
				// - application/json
				// Produces:
				// - application/json
				// Schemes: http, https
				// Deprecated: false
				// Responses:
				//	default: generalResponse
				//	200: CameraModel
				//	500: generalResponse
				camera.POST("/model", h.CameraModelCreate)

				// swagger:route GET /private/camera/model array of models
				//
				// Get all camera models
				//
				// Get an array of created models, such as [RLC-411WS, RLC-510, ...].
				//
				// Consumes:
				// - application/json
				// Produces:
				// - application/json
				// Schemes: http, https
				// Deprecated: false
				// Responses:
				//	default: generalResponse
				//	200: []CameraModel
				//	500: generalResponse
				camera.GET("/model", h.CameraModelRead)

				// swagger:route POST /private/camera create camera
				//
				// Create a new Camera
				//
				// A new camera will be returned.
				//
				// Consumes:
				// - application/json
				// Produces:
				// - application/json
				// Schemes: http, https
				// Deprecated: false
				// Responses:
				//	default: generalResponse
				//	200: Camera
				//	500: generalResponse
				camera.POST("", h.CameraCreate)

				// swagger:route GET /private/camera array of cameras
				//
				// Get all cameras created
				//
				// Get an array of created cameras
				//
				// Consumes:
				// - application/json
				// Produces:
				// - application/json
				// Schemes: http, https
				// Deprecated: false
				// Responses:
				//	default: generalResponse
				//	200: []Camera
				//	500: generalResponse
				camera.GET("", h.CameraRead)
			}

			network := private.Group("/network")
			{
				// swagger:route POST /private/network/proxy create proxy
				//
				// Create a new proxy setting
				//
				// The proxy setting is a reusable setting that can be
				// re-applied to created cameras.
				//
				// Consumes:
				// - application/json
				// Produces:
				// - application/json
				// Schemes: http, https
				// Deprecated: false
				// Responses:
				//	default: generalResponse
				//	200: Proxy
				//	500: generalResponse
				network.POST("/proxy", h.NetworkProxyCreate)
			}
		}
	}
}
