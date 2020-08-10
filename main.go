package main

import (
	"log"

	c "github.com/1k-ct/V_CruiseShip/controller"
	glv "github.com/1k-ct/V_CruiseShip/getlivevideo"
	"github.com/gin-gonic/gin"
)

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

func functionStartCruise(url string) {
	//sc := startCruise(url)

}

/*
func main() {
	url := "https://virtual-youtuber.userlocal.jp/lives"

	//	dataLink := glv.GetLivingVideo(url)
	//		for _, v := range dataLink {
	//			fmt.Println(v)
	//		}
	//		fmt.Println(len(dataLink))
	//
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "start.html", gin.H{})
		//log.Println("http respons 200 GET")
	})

	sc := startCruise(url)

	router.GET("/new", func(c *gin.Context) {
		dataLink, ok := sc()
		if ok {
			//log.Println(dataLink)
			c.HTML(200, "index.html", gin.H{"dataLink": dataLink})
			//c.Redirect(302, "/")
		} else if !ok { //"owari"
			sc = startCruise(url)
			//startCruise(url)
			//c.HTML(200, "index.html", gin.H{"dataLink": dataLink})
			c.Redirect(302, "/ggnew")
		}
	})

	router.GET("/ggnew", func(c *gin.Context) {
		c.Redirect(302, "/new")
	})

	router.GET("/stoppoint", func(c *gin.Context) {
		c.HTML(200, "stoppoint.html", gin.H{})
	})

	router.Run()
}
*/
var url string = "https://virtual-youtuber.userlocal.jp/lives"

func main() {
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
	r.Run()
}

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
	//r.GET("/new", ctrl.VideoStart)
	r.GET("/ggnew", ctrl.Interim)
	r.GET("/stoppoint", ctrl.Stop)

	return r
}
