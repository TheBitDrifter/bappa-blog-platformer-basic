package main

import (
	"embed"
	"log"

	"github.com/TheBitDrifter/coldbrew"
)

//go:embed assets/*
var assets embed.FS

const (
	RESOLUTION_X       = 640
	RESOLUTION_Y       = 360
	MAX_SPRITES_CACHED = 100
	MAX_SOUNDS_CACHED  = 100
	MAX_SCENES_CACHED  = 12
)

func main() {
	// Create the client
	client := coldbrew.NewClient(
		RESOLUTION_X,
		RESOLUTION_Y,
		MAX_SPRITES_CACHED,
		MAX_SOUNDS_CACHED,
		MAX_SCENES_CACHED,
		assets,
	)

	// Settings
	client.SetTitle("Platformer")
	client.SetResizable(true)
	client.SetMinimumLoadTime(30)

	// Run the client
	if err := client.Start(); err != nil {
		log.Fatal(err)
	}
}
