package wechat_autojump_game

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"time"

	"image/png"
)

var (
	Filename = "tmp_jump.png"
	// SleepMills 中场休息时间
	SleepMills time.Duration = 1000
	// NextStep 下一步
	NextStep = 0.0
	// AutoJump 是否自动跳
	AutoJump = false
	// LStart 开始坐标
	LStart = [2]int{0, 0}
	// LEnd 结束坐标
	LEnd = [2]int{0, 0}
)

// Run 开始运行
func Run() {
	for AutoJump {
		time.Sleep(SleepMills * time.Millisecond)
		SaveScreenShot()
		start, end := GetLocation()
		distance := GetDistance(start, end)
		NextStep = distance
		fmt.Printf("Current Location: %v, %v\nNext Location: %v,  %v\nThe Distance: %v\n", start[0], start[1], end[0], end[1], distance)
		Jump(distance)
	}
	// 是否开启自动jump
	/* GetLocation(filename)
	inputStr := ""
	fmt.Scanln(&inputStr)
	if inputStr == "exit" {
		os.Exit(0)
	} */
}

// GetLocation 获取当前地址
func GetLocation() ([]int, []int) {
	f1, err := os.Open(Filename)
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	data, _ := ioutil.ReadAll(f1)
	img, _ := png.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	start, end := Find(img)
	LStart[0] = start[0]
	LStart[1] = start[1]
	LEnd[0] = end[0]
	LEnd[1] = end[1]
	return start, end
}

// GetDistance 获取距离
func GetDistance(current, next []int) float64 {
	return math.Sqrt(math.Pow(float64(current[0]-next[0]), 2) + math.Pow(float64(current[1]-next[1]), 2))
}
