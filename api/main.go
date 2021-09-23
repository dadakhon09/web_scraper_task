package api

import (
	_ "github.com/dadakhon09/web_scraper_task/api/docs"
	"github.com/dadakhon09/web_scraper_task/api/handlers"
	"github.com/dadakhon09/web_scraper_task/config"
	"github.com/dadakhon09/web_scraper_task/pkg/logger"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterOptions struct {
	Log logger.Logger
	Cfg *config.Config
}

//@securityDefinitions.apikey ApiKeyAuth
//@in header
//@name Authorization
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(CORSMiddleware())

	config := cors.DefaultConfig()

	config.AllowAllOrigins = true
	//config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")
	config.AllowMethods = append(config.AllowMethods, "OPTIONS")

	//router.Use(cors.New(config))

	handlerV1 := handlers.New(&handlers.HandlerV1Options{
		Log: opt.Log,
		Cfg: opt.Cfg,
	})

	router.POST("/task", handlerV1.TaskHandler)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,upload-offset, upload-metadata, upload-length, tus-resumable, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, HEAD, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
