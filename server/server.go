package server

import (
	"log"

	c "github.com/1k-ct/V_CruiseShip/controller"
	glv "github.com/1k-ct/V_CruiseShip/getlivevideo"
	"github.com/gin-gonic/gin"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()

}

var url string = "https://virtual-youtuber.userlocal.jp/lives"

func router() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	ctrl := c.Controller{}
	r.GET("/", ctrl.Home)
	//r.GET("/new", ctrl.VideoStart)
	r.GET("/ggnew", ctrl.Interim)
	r.GET("/stoppoint", ctrl.Stop)

	sc := startCruise(url)
	r.GET("/new", func(c *gin.Context) {
		dataLink, ok := sc()
		if ok {
			c.HTML(200, "index.html", gin.H{"dataLink": dataLink})
		} else if !ok {
			sc = startCruise(url)
			c.Redirect(302, "/ggnew")
		}
	})
	return r
}
func startCruise(url string) func() (string, bool) {
	dataLink := glv.GetLivingVideo(url) //動画をスクレイピングしてくる
	log.Println("スクレイピング出来たよ！")
	lenDataLink := len(dataLink) // 動画の本数
	//fmt.Println(lenDataLink)
	n := -1
	return func() (string, bool) {
		n++
		if n == lenDataLink {
			return dataLink[0], false //errors.New("終了")
		}
		return dataLink[n], true //, "mada"
	}
}
