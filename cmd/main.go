package main

import (
	"flag"
	"fmt"
	"time"

	g "github.com/ZhenlyChen/wechat_autojump_game"
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
	if mills != 0 {
		g.SleepMills = time.Duration(mills)
	}
	fmt.Println("请输入跳跃的倍率（参考值1.5 - 2.5）：")
	fmt.Scanf("%f", speed)
	if speed != 0 {
		g.Speed = speed
	}
	g.Run()
}
