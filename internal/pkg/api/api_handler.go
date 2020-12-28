package api

import (
	"github.com/ReolinkCameraAPI/noctilucago/config"
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/api/responses"
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/controllers"
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/procedures"
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/service"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strings"
	"time"
)

type Handler struct {
	Router        *gin.Engine
	CameraService *service.CameraService
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

	cameraService := service.NewCameraService()

	return &Handler{
		Router:        router,
		CameraService: cameraService,
		ApiController: apiController,
	}
}

func (h *Handler) CreateEndpoints() error {

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         config.NlConfig.Auth.JWT.Issuer,
		Key:           []byte(config.NlConfig.Auth.JWT.Key),
		Timeout:       time.Second * time.Duration(config.NlConfig.Auth.JWT.Timeout),
		MaxRefresh:    time.Second * time.Duration(config.NlConfig.Auth.JWT.Refresh),
		Authenticator: h.Login,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, responses.GenericResponse{
				Status:  "error",
				Message: message,
			})
		},
	})

	if err != nil {
		return err
	}

	err = authMiddleware.MiddlewareInit()

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

			auth := public.Group("/auth")
			{
				// swagger:route POST /public/auth/login auth AuthLogin
				//
				// User Login
				//
				// Log in with a username and password
				//
				// Consumes:
				// - application/json
				// Produces:
				// - application/json
				// Schemes: http, https
				// Deprecated: false
				// Responses:
				//	default: genericResponse
				//  403: genericResponse
				auth.POST("/login", authMiddleware.LoginHandler)
			}

		}

		// private group (protected with auth)
		private := v1.Group("/private", gin.BasicAuth(gin.Accounts{
			"admin": "admin",
		}))
		private.Use(authMiddleware.MiddlewareFunc())
		{

			model := private.Group("/model")
			{
				// swagger:route POST /private/model model CreateModel
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
				//	default: genericResponse
				//	200: cameraModelResponse
				//	500: genericResponse
				model.POST("", h.CameraModelCreate)

				// swagger:route GET /private/model model GetAllModel
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
				//	default: genericResponse
				//	200: cameraModelArrayResponse
				//	500: genericResponse
				model.GET("", h.CameraModelRead)
			}

			camera := private.Group("/camera")
			{

				// swagger:route POST /private/camera/:model camera CreateCamera
				//
				// Create a new Camera
				//
				// Pass the model uuid as a parameter and the rest of the information in the body.
				// A new camera will be returned.
				//
				// Consumes:
				// - application/json
				// Produces:
				// - application/json
				// Schemes: http, https
				// Deprecated: false
				// Responses:
				//	default: genericResponse
				//	200: cameraResponse
				//	500: genericResponse
				camera.POST("/:model", h.CameraCreate)

				// swagger:route GET /private/camera camera GetAllCameras
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
				//	default: genericResponse
				//	200: cameraModelArrayResponse
				//	500: genericResponse
				camera.GET("", h.CameraRead)

				// swagger:route DELETE /private/camera camera DeleteCamera
				//
				// Delete specified camera
				//
				// Use the cameras' UUID to delete it
				//
				// Consumes:
				// - application/json
				// Produces:
				// - application/json
				// Schemes: http, https
				// Deprecated: false
				// Responses:
				//	default: genericResponse
				//	200: genericResponse
				//	500: genericResponse
				camera.DELETE("", h.CameraDelete)

				// swagger:route PUT /private/camera camera UpdateCamera
				//
				// Update the specified camera
				//
				// Update an existing cameras' settings
				//
				// Consumes:
				// - application/json
				// Produces:
				// - application/json
				// Schemes: http, https
				// Deprecated: false
				// Responses:
				//	default: genericResponse
				//	200: cameraResponse
				//	500: genericResponse
				camera.PUT("", h.CameraUpdate)
			}

			network := private.Group("/network")
			{

				// swagger:route POST /private/network/proxy proxy CreateProxy
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
				//	default: genericResponse
				//	200: networkProxyResponse
				//	500: genericResponse
				network.POST("/proxy", h.NetworkProxyCreate)

				// swagger:route PUT /private/network/proxy/:uuid proxy UpdateProxy
				//
				// Update a proxy setting
				//
				// Pass the proxy's UUID with the updated proxy information
				//
				// Consumes:
				// - application/json
				// Produces:
				// - application/json
				// Schemes: http, https
				// Deprecated: false
				// Responses:
				//	default: genericResponse
				//	200: networkProxyResponse
				//	500: genericResponse
				network.PUT("/proxy/:uuid", h.NetworkProxyUpdate)

				// swagger:route GET /private/network/proxy/:uuid proxy GetOneProxy
				//
				// Get a singular proxy using its UUID
				//
				// Get the proxy settings object using its UUID
				//
				// Consumes:
				// - application/json
				// Produces:
				// - application/json
				// Schemes: http, https
				// Deprecated: false
				// Responses:
				//	default: genericResponse
				//	200: networkProxyResponse
				//	500: genericResponse
				network.GET("/proxy/:uuid", h.NetworkProxyReadUUID)

				// swagger:route GET /private/network/proxies proxy GetAllProxies
				//
				// Get all the proxies
				//
				// Get all the created proxies
				//
				// Consumes:
				// - application/json
				// Produces:
				// - application/json
				// Schemes: http, https
				// Deprecated: false
				// Responses:
				//	default: genericResponse
				//	200: networkProxyArrayResponse
				//	500: genericResponse
				network.GET("/proxies", h.NetworkProxyRead)

				// swagger:route GET /private/network/proxies/schemes schemes GetAllSchemes
				//
				// Get all the proxy schemes accepted by the server
				//
				// The scheme can be HTTP, HTTPS or SOCKS5
				//
				// Consumes:
				// - application/json
				// Produces:
				// - application/json
				// Schemes: http, https
				// Deprecated: false
				// Responses:
				//	default: genericResponse
				//	200: networkProxySchemeResponse
				network.GET("/proxies/schemes", h.NetworkProxyReadScheme)

				// swagger:route GET /private/network/protocols protocols GetAllProtocols
				//
				// Get all the protocols accepted by the server
				//
				// A protocol can be UDP or TCP
				//
				// Consumes:
				// - application/json
				// Produces:
				// - application/json
				// Schemes: http, https
				// Deprecated: false
				// Responses:
				//	default: genericResponse
				//	200: networkProtocolResponse
				network.GET("/protocols", h.NetworkReadProtocol)
			}

			user := private.Group("/user")
			{

				// swagger:route GET /private/user user GetAllUsers
				//
				// Get all the User Accounts
				//
				//
				// Consumes:
				// - application/json
				// Produces:
				// - application/json
				// Schemes: http, https
				// Deprecated: false
				// Responses:
				//	default: genericResponse
				//	200: userArrayResponse
				//	500: genericResponse
				user.GET("", h.UserRead)

				// swagger:route POST /private/user user CreateUser
				//
				// Create a new User account
				//
				// Create a new User account for managing cameras
				//
				// Consumes:
				// - application/json
				// Produces:
				// - application/json
				// Schemes: http, https
				// Deprecated: false
				// Responses:
				//	default: genericResponse
				//	200: userResponse
				//	500: genericResponse
				user.POST("", h.UserCreate)

				// swagger:route PUT /private/user/:uuid user UpdateUser
				//
				// Update User
				//
				// Update an existing user's credentials
				//
				// Consumes:
				// - application/json
				// Produces:
				// - application/json
				// Schemes: http, https
				// Deprecated: false
				// Responses:
				//	default: genericResponse
				//	200: userResponse
				//	500: genericResponse
				user.PUT("/:uuid", h.UserUpdate)

				// swagger:route DELETE /private/user user DeleteUser
				//
				// Delete User
				//
				// Delete an existing user account
				//
				// Consumes:
				// - application/json
				// Produces:
				// - application/json
				// Schemes: http, https
				// Deprecated: false
				// Responses:
				//	default: genericResponse
				//	200: genericResponse
				//	500: genericResponse
				user.DELETE("/:uuid", h.UserDelete)
			}

			auth := private.Group("/auth")
			{
				// swagger:route GET /private/auth/refresh auth AuthRefresh
				//
				// Refresh the JWT token
				//
				// The refresh token is set according to the `noctiluca` config `Refresh` option.
				//
				// Consumes:
				// - application/json
				// Produces:
				// - application/json
				// Schemes: http, https
				// Deprecated: false
				// Responses:
				//	default: genericResponse
				//	500: genericResponse
				auth.GET("refresh", authMiddleware.RefreshHandler)
			}
		}

	}

	return nil
}
