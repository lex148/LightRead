package main

import (
	"bufio"
	"fmt"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
	"io"
	"os"
	"os/exec"
)

type SourceType float64

const (
	Clipboard = iota
	Selection = iota
)

func main() {
	X, _ := xgbutil.NewConn()
	keybind.Initialize(X)

	cb1 := keybind.KeyPressFun(
		func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			read(Selection)
		})
	cb1.Connect(X, X.RootWin(), "control-z", true)

	cb2 := keybind.KeyPressFun(
		func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			read(Clipboard)
		})
	cb2.Connect(X, X.RootWin(), "control-q", true)

	fmt.Printf("Listening for Crtl-z ...(Read Selected Text)\n")
	fmt.Printf("Listening for Crtl-q ...(Read Clipboard Text)\n")
	xevent.Main(X)

}

var player (exec.Cmd)
var tts_engine (exec.Cmd)
var xsel (exec.Cmd)

func read(source SourceType) {
	////conflicts with swift lang
	//_, err := exec.LookPath("swift")
	//if err == nil {
	//	swift_read(source)
	//} else {
	festival_read(source)
	//}
}

func kill_all_the_things() {
	if player.Process != nil {
		player.Process.Kill()
		player.Process.Wait()
	}
	if tts_engine.Process != nil {
		tts_engine.Process.Kill()
		tts_engine.Process.Wait()
	}
	if xsel.Process != nil {
		xsel.Process.Kill()
		xsel.Process.Wait()
	}
}

func festival_read(source SourceType) {
	kill_all_the_things()
	xsel = *build_xsel_command(source)
	tts_engine = *exec.Command("text2wave", "-o", "/dev/stdout", "/dev/stdin")
	player = *exec.Command("aplay", "-fS16_LE", "-r16000", "/dev/stdin")
	xsel_pipe, _ := xsel.StdoutPipe()
	tts_engine.Stdin = xsel_pipe
	tts_engine_pipe, _ := tts_engine.StdoutPipe()
	player.Stdin = tts_engine_pipe
	xsel.Start()
	tts_engine.Start()
	player.Start()
}

func swift_read(source SourceType) {
	if player.Process != nil {
		player.Process.Kill()
		player.Process.Wait()
	}
	write_sel_to_temp_file(source)
	tts_engine := exec.Command("swift", "-o", "/tmp/light_read.wav", "-f", "/tmp/light_read.txt")
	tts_engine.Run()
	os.Remove("/tmp/light_read.txt")
	player = *exec.Command("aplay", "/tmp/light_read.wav")
	player.Start()
}

func build_xsel_command(source SourceType) *exec.Cmd {

	switch {
	case source == Selection:
		return exec.Command("xsel")
	case source == Clipboard:
		return exec.Command("xsel", "-b")
	}
	return exec.Command("xsel")
}

func write_sel_to_temp_file(source SourceType) {
	text, _ := os.Create("/tmp/light_read.txt")
	w := bufio.NewWriter(text)
	buf := make([]byte, 1024)
	xsel := build_xsel_command(source)
	r, _ := xsel.StdoutPipe()
	xsel.Start()
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		if _, err := w.Write(buf[:n]); err != nil {
			panic(err)
		}
	}
	w.Flush()
	text.Close()
}
