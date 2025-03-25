package coresystems

import (
	"platformer/components"

	"github.com/TheBitDrifter/blueprint"
	"github.com/TheBitDrifter/warehouse"
)

type OnGroundClearingSystem struct{}

func (OnGroundClearingSystem) Run(scene blueprint.Scene, dt float64) error {
	const expirationTicks = 15

	onGroundQuery := warehouse.Factory.NewQuery().And(components.OnGroundComponent)
	onGroundCursor := scene.NewCursor(onGroundQuery)

	// Iterate through matched entities
	for range onGroundCursor.Next() {
		onGround := components.OnGroundComponent.GetFromCursor(onGroundCursor)

		// If it's expired, remove it
		if scene.CurrentTick()-onGround.LastTouch > expirationTicks {
			groundedEntity, _ := onGroundCursor.CurrentEntity()

			// We can't mutate while iterating so we enqueue the changes instead
			err := groundedEntity.EnqueueRemoveComponent(components.OnGroundComponent)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
