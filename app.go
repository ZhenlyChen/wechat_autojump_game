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
	filename = "tmp_jump.png"
	// SleepMills 中场休息时间
	SleepMills time.Duration = 1000
)

// Run 开始运行
func Run() {
	for {
		time.Sleep(SleepMills * time.Millisecond)
		saveScreenShot(filename)
		distance := getLocation(filename)
		jump(distance)
		// 是否开启自动jump
		/* getLocation(filename)
		inputStr := ""
		fmt.Scanln(&inputStr)
		if inputStr == "exit" {
			os.Exit(0)
		} */
	}
}

func getLocation(name string) float64 {
	f1, err := os.Open(name)
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
	distance := getDistance(start[0], start[1], end[0], end[1])
	fmt.Printf("Current Location: %v, %v\nNext Location: %v,  %v\nThe Distance: %v\n", start[0], start[1], end[0], end[1], distance)
	return distance
}

func getDistance(currentX, currentY, nextX, nextY int) float64 {
	return math.Sqrt(math.Pow(float64(currentX-nextX), 2) + math.Pow(float64(currentY-nextY), 2))
}
