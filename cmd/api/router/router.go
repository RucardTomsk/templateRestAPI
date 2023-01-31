package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"templateRestAPI/cmd/api/controller"
	"templateRestAPI/cmd/config"
	"templateRestAPI/internal/api/middleware"
	"templateRestAPI/internal/common"
	"templateRestAPI/internal/telemetry/metric"
)

type Router struct {
	config config.Config
}

func NewRouter(config config.Config) *Router {
	return &Router{
		config: config,
	}
}

func (h *Router) InitRoutes(
	logger *zap.Logger,
	controllerContainer *controller.Container,
) *gin.Engine {
	gin.SetMode(h.config.Server.GinMode)

	router := gin.Default()

	router.Use(gin.Recovery())
	router.Use(middleware.SetMetrics(*logger))
	router.Use(middleware.SetTracingContext(*logger))
	router.Use(middleware.SetRequestLogging(*logger))
	router.Use(middleware.SetRateLimiter(h.config.Server, *logger))
	router.Use(middleware.SetOperationName(h.config.Server, *logger))
	router.Use(middleware.SetAuthorizationCheck(h.config.Server, *logger))
	router.Use(middleware.SetAccessControl(h.config.Server, *logger))
	router.Use(cors.New(common.DefaultCorsConfig()))

	p := metric.NewPrometheus("gin")
	p.Use(router)

	router.GET("api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
