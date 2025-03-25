package scenes

import (
	"platformer/animations"
	"platformer/components"

	"github.com/TheBitDrifter/blueprint"
	"github.com/TheBitDrifter/blueprint/vector"
	"github.com/TheBitDrifter/warehouse"

	blueprintclient "github.com/TheBitDrifter/blueprint/client"
	blueprintinput "github.com/TheBitDrifter/blueprint/input"
	blueprintmotion "github.com/TheBitDrifter/blueprint/motion"
	blueprintspatial "github.com/TheBitDrifter/blueprint/spatial"
)

type Scene struct {
	Name          string
	Plan          blueprint.Plan
	Width, Height int
}

func NewPlayer(sto warehouse.Storage, x, y float64) error {
	// Get or create the archetype
	playerArchetype, err := sto.NewOrExistingArchetype(
		blueprintspatial.Components.Position,
		blueprintspatial.Components.Position,
		blueprintspatial.Components.Shape,
		blueprintspatial.Components.Direction,
		blueprintmotion.Components.Dynamics,
		blueprintinput.Components.InputBuffer,
		blueprintclient.Components.CameraIndex,
		blueprintclient.Components.SpriteBundle,
		blueprintclient.Components.SoundBundle,
	)

	// Position state
	playerPos := blueprintspatial.NewPosition(x, y)
	// Hitbox state
	playerHitbox := blueprintspatial.NewRectangle(18, 58)
	// Physics state
	playerDynamics := blueprintmotion.NewDynamics(10)
	// Basic Direction State
	playerDirection := blueprintspatial.NewDirectionRight()
	// Input state
	playerInputBuffer := blueprintinput.InputBuffer{ReceiverIndex: 0}
	// Camera Reference
	playerCameraIndex := blueprintclient.CameraIndex(0)
	// Sprite Reference
	playerSprites := blueprintclient.NewSpriteBundle().
		AddSprite("characters/box_man_sheet.png", true).
		WithAnimations(animations.IdleAnimation, animations.RunAnimation, animations.FallAnimation, animations.JumpAnimation).
		SetActiveAnimation(animations.IdleAnimation).
		WithOffset(vector.Two{X: -72, Y: -59}).
		WithPriority(20)

	// Generate the player
	err = playerArchetype.Generate(1,
		playerPos,
		playerHitbox,
		playerDynamics,
		playerDirection,
		playerInputBuffer,
		playerCameraIndex,
		playerSprites,
	)
	if err != nil {
		return err
	}
	return nil
}

func NewFloor(sto warehouse.Storage, y float64) error {
	terrainArchetype, err := sto.NewOrExistingArchetype(
		blueprintclient.Components.SpriteBundle,
		components.BlockTerrainTag,
		blueprintspatial.Components.Shape,
		blueprintspatial.Components.Position,
		blueprintmotion.Components.Dynamics,
	)
	if err != nil {
		return err
	}
	return terrainArchetype.Generate(1,
		blueprintspatial.NewPosition(1500, y),
		blueprintspatial.NewRectangle(4000, 50),
		blueprintclient.NewSpriteBundle().
			AddSprite("terrain/floor.png", true).
			WithOffset(vector.Two{X: -1500, Y: -25}),
	)
}

func NewInvisibleWalls(sto warehouse.Storage, width, height int) error {
	terrainArchetype, err := sto.NewOrExistingArchetype(
		blueprintclient.Components.SpriteBundle,
		components.BlockTerrainTag,
		blueprintspatial.Components.Shape,
		blueprintspatial.Components.Position,
		blueprintmotion.Components.Dynamics,
	)
	if err != nil {
		return err
	}

	// Wall left (invisible)
	err = terrainArchetype.Generate(1,
		blueprintspatial.NewRectangle(10, float64(height+300)),
		blueprintspatial.NewPosition(0, 0),
	)
	if err != nil {
		return err
	}

	// Wall right (invisible)
	return terrainArchetype.Generate(1,
		blueprintspatial.NewRectangle(10, float64(height+300)),
		blueprintspatial.NewPosition(float64(width), 0),
	)
}

func NewBlock(sto warehouse.Storage, x, y float64) error {
	terrainArchetype, err := sto.NewOrExistingArchetype(
		blueprintclient.Components.SpriteBundle,
		components.BlockTerrainTag,
		blueprintspatial.Components.Shape,
		blueprintspatial.Components.Position,
		blueprintmotion.Components.Dynamics,
	)
	if err != nil {
		return err
	}
	return terrainArchetype.Generate(1,
		blueprintspatial.NewPosition(x, y),
		blueprintspatial.NewRectangle(64, 75),
		blueprintclient.NewSpriteBundle().
			AddSprite("terrain/block.png", true).
			WithOffset(vector.Two{X: -33, Y: -38}),
	)
}
