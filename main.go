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
	"github.com/itchyny/volume-go"
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

	Volume struct {
		Down struct {
		} `cmd:"" help:"Decrease volume."`
		Up struct {
		} `cmd:"" help:"Increase volume."`
		Set struct {
			Level int `arg:"" name:"level" help:"Set volume level (0-100)."`
		} `cmd:"" help:"Set volume level."`
	} `cmd:"" help:"Volume control."`
}

func play(url string) {
	ctx := &daemon.Context{
		PidFileName: "/tmp/sample.pid",
		PidFilePerm: 0644,
		LogFileName: "/tmp/sample.log",
		LogFilePerm: 0640,
		WorkDir:     "/tmp/",
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

	err = exec.Command("mplayer", url, "-softvol").Run()
	if err != nil {
		log.Print("Error starting mplayer: ", err)
	}

	log.Print("daemon terminated")
}

func stop() {
	pidBytes, err := os.ReadFile("/tmp/sample.pid")
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
		time.Sleep(time.Millisecond * 100)
		play(ctx.Args[1])
	case "volume up":
		volume.IncreaseVolume(10)
	case "volume down":
		volume.IncreaseVolume(-10)
	case "volume set <level>":
		level, err := strconv.Atoi(ctx.Args[2])
		if err != nil {
			log.Print("Error parsing volume level: ", err)
			return
		}
		if level < 0 {
			level = 0
		}
		if level > 100 {
			level = 100
		}
		volume.SetVolume(level)
	default:
		panic(ctx.Command())
	}
}
