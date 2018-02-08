package wechat_autojump_game

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"strings"
)

var (
	// Speed 跳跃距离的倍率，不同手机可能要手动调
	Speed float64 = 2.1
)

func Jump(distance float64) {
	pressTime := distance * Speed
	x, y := randomPosition()
	runAdb("shell", fmt.Sprintf("input swipe %d %d %d %d %d", x, y, x, y, int(pressTime)))
}

func SaveScreenShot() {
	filePath := "/sdcard/" + Filename
	runAdb("shell", "screencap -p "+filePath)
	runAdb("pull", filePath, ".")
}

func ExecAdb(args ...string) string {
	return runAdb(args...)
}

func runAdb(args ...string) string {
	var b bytes.Buffer
	cmd := exec.Command(".\\lib\\adb.exe", args...)
	cmd.Stdout = &b
	cmd.Stderr = &b
	err := cmd.Run()
	if cmd.Process != nil {
		cmd.Process.Kill()
	}
	if err != nil {
		fmt.Println(err)
		log.Fatalf("adb %s: %v", strings.Join(args, " "), err.Error())
	}
	return b.String()
}

//x : 350 - 450
//y : 850 - 950
func randomPosition() (x int, y int) {
	x = 350 + rand.Intn(100)
	y = 850 + rand.Intn(100)
	return
}
