package main

import (
	"embed"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	log"github.com/sirupsen/logrus"

	"ThothNotes/backend"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	userDir, err := os.UserHomeDir()
	if err != nil {
		log.Error(err)
		userDir = ""
	}
	// Create an instance of the app structure
	store := backend.NewStore(userDir + "/notes.db")
	app := backend.NewNotesHandlers(store)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "Thoth Notes",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.Error("Error:", err.Error())
	}
}
