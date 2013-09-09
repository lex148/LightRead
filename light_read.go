package main

import (
	"fmt"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
	"os/exec"
)

func main() {

	X, _ := xgbutil.NewConn()
	keybind.Initialize(X)

	cb1 := keybind.KeyPressFun(
		func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			read_sel()
		})
	cb1.Connect(X, X.RootWin(), "control-z", true)

	cb2 := keybind.KeyPressFun(
		func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			read_clip()
		})
	cb2.Connect(X, X.RootWin(), "control-q", true)

	fmt.Printf("Listening for Crtl-z ...(Read Selected Text)\n")
	fmt.Printf("Listening for Crtl-q ...(Read Clipboard Text)\n")
	xevent.Main(X)

}

var player (exec.Cmd)

func read_clip() {
	if player.Process != nil {
		player.Process.Kill()
		player.Process.Wait()
	}
	xsel := exec.Command("xclip", "-o")
	festival := exec.Command("text2wave", "-o", "/dev/stdout", "/dev/stdin")
	player = *exec.Command("aplay", "/dev/stdin")
	xsel_pipe, _ := xsel.StdoutPipe()
	festival.Stdin = xsel_pipe
	festival_pipe, _ := festival.StdoutPipe()
	player.Stdin = festival_pipe
	xsel.Start()
	festival.Start()
	player.Start()
}

func read_sel() {
	if player.Process != nil {
		player.Process.Kill()
		player.Process.Wait()
	}
	xsel := exec.Command("xsel")
	festival := exec.Command("text2wave", "-o", "/dev/stdout", "/dev/stdin")
	player = *exec.Command("aplay", "/dev/stdin")
	xsel_pipe, _ := xsel.StdoutPipe()
	festival.Stdin = xsel_pipe
	festival_pipe, _ := festival.StdoutPipe()
	player.Stdin = festival_pipe
	xsel.Start()
	festival.Start()
	player.Start()
}
