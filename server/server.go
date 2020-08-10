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

	u := r.Group("/startPP")
	{
		ctrl := c.Controller{}
		u.GET("/", ctrl.Home)
		u.GET("/new", ctrl.VideoStart)
		u.GET("/ggnew", ctrl.Interim)
		u.GET("/stoppoint", ctrl.Stop)
	}
	return r
}
