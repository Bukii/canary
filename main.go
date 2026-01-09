package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/alecthomas/kong"
	"github.com/sevlyar/go-daemon"
)

var CLI struct {
	Play struct {
		URL string `optional:"" arg:"" name:"url" help:"URL to play."`
	} `cmd:"" help:"Play audio."`

	Stop struct {
	} `cmd:"" help:"Stop audio."`
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
	defer exec.Command("pkill", "-x", "ffplay").Run()

	log.Print("- - - - - - - - - - - - - - -")
	log.Print("daemon started")

	exec.Command("ffplay", url, "-loglevel", "quiet", "-nodisp").Run()
}

func stop() {
	pid, err := os.ReadFile("sample.pid")
	if err != nil {
		log.Print("Error reading pid file: ", err)
		return
	}
	exec.Command("kill", string(pid)).Run()
	exec.Command("pkill", "-x", "ffplay").Run()
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
	default:
		panic(ctx.Command())
	}
}
