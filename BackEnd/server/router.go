package server

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"sushi/utils/config"
	"sushi/utils/ratelimit"
)

func NewRouter(server *Server, conf config.Config, socketserver *socketio.Server) *gin.Engine {
	gin.SetMode(server.config.GinMode())
	r := gin.Default()

	r.Use(ratelimit.GinMiddleware())
	r.Use(CORSMiddleware())
	r.Use(static.Serve("/", static.LocalFile("/app/Demo-UI", true)))

	//public
	r.GET("/ping", server.controller.HandlePing)
	r.GET("/toponym", server.controller.HandleToponym)
	r.GET("/unit", server.controller.HandleGetUnit)
	r.POST("/unit", server.controller.HandleRegisterUnit)
	r.GET("/record", server.controller.HandleTransRecord)
	v1 := r.Group("/v1")
	authorizedV1 := v1.Group("/")
	authorizedV1.Use(server.GetAuth())
	WithUserRoutes(authorizedV1, server, conf)
	return r
}

func WithTeamRoutes(r *gin.RouterGroup, server *Server) {
	//r.GET("/", server.controller.team.HandleTeamList)
}

func WithUserRoutes(r *gin.RouterGroup, server *Server, conf config.Config) {
	authorized := r
	authorized.POST("/users/profile")
}

func (server Server) GetAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		//AUTH
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Auth-Token, Authorization, Code, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT , PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
