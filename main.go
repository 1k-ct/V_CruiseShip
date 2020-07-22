package main

import (
	"fmt"
	"log"

	glv "./getlivevideo"
	"github.com/gin-gonic/gin"
)

func startCruise(url string) func() (string, string) {
	dataLink := glv.GetLivingVideo(url) //動画をスクレイピングしてくる
	log.Println("スクレイピング出来たよ！")
	lenDataLink := len(dataLink) // 動画の本数
	//fmt.Println(lenDataLink)
	n := 0
	return func() (string, string) {
		n++
		if n == lenDataLink {
			return dataLink[0], "owari" //errors.New("終了")
		}
		return dataLink[n], "mada"
	}
}
func functionStartCruise(url string) {
	//sc := startCruise(url)

}
func main() {
	url := "https://virtual-youtuber.userlocal.jp/lives"
	dataLink := glv.GetLivingVideo(url)

	for _, v := range dataLink {
		fmt.Println(v)
	}
	fmt.Println(len(dataLink))

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	hello := "hello gin"
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"test":        hello,
			"dataLinkLen": len(dataLink),
		})
		//log.Println("http respons 200 GET")
	})

	sc := startCruise(url)

	router.GET("/new", func(c *gin.Context) {
		dataLink, owari := sc()
		if owari == "owari" { //"owari"
			sc = startCruise(url)
			//startCruise(url)
			c.HTML(200, "index.html", gin.H{"dataLink": dataLink[0]})
		}
		log.Println(dataLink)
		c.HTML(200, "index.html", gin.H{"dataLink": dataLink})
		//c.Redirect(302, "/")
	})

	router.GET("/stoppoint", func(c *gin.Context) {
		c.HTML(200, "stoppoint.html", gin.H{})
	})

	router.Run()
}
