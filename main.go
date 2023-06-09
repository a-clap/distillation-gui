package main

import (
	"embed"
	"flag"
	"log"
	"os"

	"github.com/a-clap/distillation-gui/backend"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

var (
	mock = flag.Bool("mock", false, "use mocks")
	// addr = flag.String("addr", "bananapi-zero.local:50002", "the distillation port")
	addr = flag.String("addr", "localhost:50002", "the distillation port")
)

func main() {
	flag.Parse()

	var opts []backend.Option
	if *mock {
		opts = mockClients()
	} else {
		opts = getopts(*addr)
	}

	p, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	opts = append(opts, backend.WithLoadSaver(&saver{path: p}))

	// Create backend
	back, err := backend.New(
		opts...,
	)
	if err != nil {
		log.Fatalln(err)
	}

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "gui",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        back.Startup,
		Bind: []interface{}{
			back,
			&backend.Events{},
			&backend.Models{},
		},
	})

	if err != nil {
		log.Fatalln("Error:", err.Error())
	}
}
