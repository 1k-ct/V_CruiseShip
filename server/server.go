package server

import (
	c "github.com/1k-ct/V_CruiseShip/controller"
	"github.com/gin-gonic/gin"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()

}

func router() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	
	ctrl := c.Controller{}
	r.GET("/", ctrl.Home)
	r.GET("/new", ctrl.VideoStart)
	r.GET("/ggnew", ctrl.Interim)
	r.GET("/stoppoint", ctrl.Stop)
	
	return r
}
