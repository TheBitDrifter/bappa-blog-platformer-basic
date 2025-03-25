package coresystems

import (
	"platformer/actions"
	"platformer/components"

	"github.com/TheBitDrifter/blueprint"
	blueprintinput "github.com/TheBitDrifter/blueprint/input"
	blueprintmotion "github.com/TheBitDrifter/blueprint/motion"
	blueprintspatial "github.com/TheBitDrifter/blueprint/spatial"
)

const (
	speed     = 120.0
	jumpforce = 220.0
)

type PlayerMovementSystem struct{}

func (sys PlayerMovementSystem) Run(scene blueprint.Scene, dt float64) error {
	// Query all entities with input buffers (players)
	cursor := scene.NewCursor(blueprint.Queries.InputBuffer)

	for range cursor.Next() {
		dyn := blueprintmotion.Components.Dynamics.GetFromCursor(cursor)
		incomingInputs := blueprintinput.Components.InputBuffer.GetFromCursor(cursor)
		direction := blueprintspatial.Components.Direction.GetFromCursor(cursor)
		isGrounded := components.OnGroundComponent.CheckCursor(cursor)

		_, pressedLeft := incomingInputs.ConsumeInput(actions.Left)
		if pressedLeft {
			direction.SetLeft()
			dyn.Vel.X = -speed
		}

		_, pressedRight := incomingInputs.ConsumeInput(actions.Right)
		if pressedRight {
			direction.SetRight()

			dyn.Vel.X = speed
		}
		_, pressedUp := incomingInputs.ConsumeInput(actions.Jump)
		if pressedUp && isGrounded {
			dyn.Vel.Y = -jumpforce
		}

		// Add down handling here:
		_, pressedDown := incomingInputs.ConsumeInput(actions.Down)
		if pressedDown && !pressedUp { // <- you cant drop and jump same tick
			playerEntity, _ := cursor.CurrentEntity()
			err := playerEntity.EnqueueAddComponent(components.IgnorePlatformComponent)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
