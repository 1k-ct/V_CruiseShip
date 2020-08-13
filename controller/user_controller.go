package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	glv "github.com/1k-ct/V_CruiseShip/getlivevideo"
)

// Controller is user controlller
type Controller struct{}

//動画の情報を取ってくるURL
var url string = "https://virtual-youtuber.userlocal.jp/lives"

func startCruise(url string) ([]string, bool) {
	dataLink := glv.GetLivingVideo(url) //動画をスクレイピングしてくる
	log.Println("スクレイピング出来たよ！")
	return dataLink, true //, "mada"

}

//Home 一番初めのページ "/ggnew"
func (pc Controller) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "start.html", gin.H{})
}

/*
//VideoStart 視聴スタート
func (pc Controller) VideoStart(c *gin.Context) {
	dataLink, ok := startCruise(url)
	n := 0
	if ok {
		c.HTML(200, "index.html", gin.H{"dataLink": dataLink})
	} else if !ok {
		sc = startCruise(url)
		c.Redirect(302, "/ggnew")
	}
}
*/
//Interim VideoStart が、最後まで行われるとInterimが実行される
func (pc Controller) Interim(c *gin.Context) {
	c.Redirect(302, "/new")
}

// Stop 緊急時のストップポイント
func (pc Controller) Stop(c *gin.Context) {
	c.HTML(http.StatusOK, "stoppoint.html", gin.H{})
}
