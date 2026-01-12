package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/alecthomas/kong"
	"github.com/sevlyar/go-daemon"
)

var CLI struct {
	Play struct {
		URL string `optional:"" arg:"" name:"url" help:"URL to play."`
	} `cmd:"" help:"Play audio."`

	Stop struct {
	} `cmd:"" help:"Stop audio."`

	Switch struct {
		URL string `arg:"" name:"url" help:"URL to switch to."`
	} `cmd:"" help:"Switch audio url."`
}

func play(url string) {
	ctx := &daemon.Context{
		PidFileName: "sample.pid",
		PidFilePerm: 0644,
		LogFileName: "sample.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
	}

	d, err := ctx.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer ctx.Release()

	log.Print("- - - - - - - - - - - - - - -")
	log.Print("daemon started")

	exec.Command("mplayer", url).Run()

	log.Print("daemon terminated")
}

func stop() {
	pidBytes, err := os.ReadFile("sample.pid")
	if err != nil {
		log.Print("Error reading pid file: ", err)
		return
	}
	pid, err := strconv.Atoi(strings.TrimSpace(string(pidBytes)))
	if err != nil {
		log.Print("Error parsing pid: ", err)
		return
	}
	syscall.Kill(-pid, syscall.SIGKILL)
}

func main() {
	ctx := kong.Parse(&CLI)

	switch ctx.Command() {
	case "play <url>":
		play(ctx.Args[1])
	case "play":
		play("http://orf-live.ors-shoutcast.at/oe3-q2a")
	case "stop":
		stop()
	case "switch <url>":
		stop()
		time.Sleep(time.Millisecond * 500)
		play(ctx.Args[1])
	default:
		panic(ctx.Command())
	}
}
