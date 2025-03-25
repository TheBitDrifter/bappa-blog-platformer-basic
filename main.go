package main

import (
	"embed"
	"log"
	"platformer/actions"
	"platformer/clientsystems"
	"platformer/coresystems"
	"platformer/scenes" // Import our scenes package

	"github.com/TheBitDrifter/coldbrew"
	coldbrew_clientsystems "github.com/TheBitDrifter/coldbrew/clientsystems"
	coldbrew_rendersystems "github.com/TheBitDrifter/coldbrew/rendersystems"
	"github.com/hajimehoshi/ebiten/v2"
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

	// Configure client settings
	client.SetTitle("Platformer")
	client.SetResizable(true)
	client.SetMinimumLoadTime(30)

	// Register scene One
	err := client.RegisterScene(
		scenes.SceneOne.Name,
		scenes.SceneOne.Width,
		scenes.SceneOne.Height,
		scenes.SceneOne.Plan,
		[]coldbrew.RenderSystem{},
		clientsystems.DefaultClientSystems,
		coresystems.DefaultCoreSystems,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Register global systems
	client.RegisterGlobalRenderSystem(
		coldbrew_rendersystems.GlobalRenderer{},
		&coldbrew_rendersystems.DebugRenderer{},
	)

	// Activate the camera
	client.ActivateCamera()

	// Register receiver/actions
	receiver1, _ := client.ActivateReceiver()
	receiver1.RegisterKey(ebiten.KeySpace, actions.Jump)
	receiver1.RegisterKey(ebiten.KeyW, actions.Jump)
	receiver1.RegisterKey(ebiten.KeyA, actions.Left)
	receiver1.RegisterKey(ebiten.KeyD, actions.Right)
	receiver1.RegisterKey(ebiten.KeyS, actions.Down)

	// Default client systems for camera mapping and receiver mapping
	client.RegisterGlobalClientSystem(
		coldbrew_clientsystems.InputBufferSystem{},
		&coldbrew_clientsystems.CameraSceneAssignerSystem{},
	)

	// Run the client
	if err := client.Start(); err != nil {
		log.Fatal(err)
	}
}
