package main

import (
	"flag"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"

	g "github.com/ZhenlyChen/wechat_autojump_game"
	"github.com/gin-gonic/gin"
)

var (
	mills int64
	speed float64
)

func init() {
	flag.Int64Var(&mills, "m", 0, "millseconds sleep after each jump")
	flag.Float64Var(&speed, "s", 0, fmt.Sprintf("speed value, default to , %.4f", g.Speed))
	flag.Parse()
}
func main() {
	r := gin.Default()
	r.Static("/console", "./public")  // 静态文件
	r.GET("/", func(c *gin.Context) { // 跳转到主页
		c.Redirect(http.StatusMovedPermanently, "/console/")
	})
	r.StaticFile("/image.png", "./jump.720.png") // 截图
	r.GET("/api/adbStatus", func(c *gin.Context) {
		str := g.ExecAdb("devices")
		c.String(http.StatusOK, str)
	})
	r.POST("/api/restartAdb", func(c *gin.Context) {
		str := g.ExecAdb("kill-server")
		str += g.ExecAdb("start-server")
		c.String(http.StatusOK, str)
	})
	r.POST("/api/setPort", func(c *gin.Context) {
		str := g.ExecAdb("tcpip", "23333")
		c.String(http.StatusOK, str)
	})
	r.POST("/api/connect", func(c *gin.Context) {
		str := g.ExecAdb("connect", c.Query("ip")+":23333")
		c.String(http.StatusOK, str)
	})
	r.GET("/api/getAutoStatus", func(c *gin.Context) {
		c.String(http.StatusOK, strconv.FormatBool(g.AutoJump))
	})
	r.GET("/api/getXY", func(c *gin.Context) {
		var msg struct {
			StartX int     `json:"StartX"`
			StartY int     `json:"StartY"`
			EndX   int     `json:"EndX"`
			EndY   int     `json:"EndY"`
			Dis    float64 `json:"Dis"`
			Speed  float64 `json:"Speed"`
		}
		msg.StartX = g.LStart[0]
		msg.StartY = g.LStart[1]
		msg.EndX = g.LEnd[0]
		msg.EndY = g.LEnd[1]
		msg.Dis = g.NextStep
		msg.Speed = g.Speed
		c.JSON(http.StatusOK, msg)
	})
	r.POST("/api/jump", func(c *gin.Context) {
		g.SaveScreenShot()
		start, end := g.GetLocation()
		dis := g.GetDistance(start, end)
		g.NextStep = dis
		g.Jump(g.NextStep)
		c.String(http.StatusOK, "OK")
	})
	r.POST("/api/autoJump", func(c *gin.Context) {
		if g.AutoJump {
			c.String(http.StatusBadRequest, "已经在运行")
		} else {
			g.AutoJump = true
			go g.Run()
			c.String(http.StatusOK, "OK")
		}
	})
	r.POST("/api/setSpeed", func(c *gin.Context) {
		speed, err := strconv.ParseFloat(c.Query("speed"), 64)
		g.Speed = speed
		if err == nil {
			c.String(http.StatusOK, "OK")
		} else {
			c.String(http.StatusOK, "参数无效")
		}
	})
	r.POST("/api/stopJump", func(c *gin.Context) {
		if !g.AutoJump {
			c.String(http.StatusBadRequest, "没有运行")
		} else {
			g.AutoJump = false
			c.String(http.StatusOK, "OK")
		}
	})
	cmd := exec.Command("cmd", "/C", "start", "http://localhost:8080") // 启动浏览器
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	http.ListenAndServe(":8080", r)
	/* for {
		g.Run()
	} */
}
