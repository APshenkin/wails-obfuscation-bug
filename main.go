package main

import (
	"embed"

	"changeme/api"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	first := api.NewFirst()
	second := api.NewSecond()
	third := api.NewThird()
	fourth := api.NewFourth()
	fifth := api.NewFifth()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "wails-obfuscation-bug",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			first,
			second,
			third,
			fourth,
			fifth,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
