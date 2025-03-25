package coresystems

import (
	"platformer/components"

	"github.com/TheBitDrifter/blueprint"
	blueprintmotion "github.com/TheBitDrifter/blueprint/motion"
	blueprintspatial "github.com/TheBitDrifter/blueprint/spatial"
	"github.com/TheBitDrifter/blueprint/vector"
	"github.com/TheBitDrifter/tteokbokki/motion"
	"github.com/TheBitDrifter/tteokbokki/spatial"
	"github.com/TheBitDrifter/warehouse"
)

type PlayerPlatformCollisionSystem struct {
	playerLastPositions []vector.Two
	maxPositionsToTrack int
}

func NewPlayerPlatformCollisionSystem() *PlayerPlatformCollisionSystem {
	trackCount := 15 // higher count == more tunneling protection == higher cost
	return &PlayerPlatformCollisionSystem{
		playerLastPositions: make([]vector.Two, 0, trackCount),
		maxPositionsToTrack: trackCount,
	}
}

func (s *PlayerPlatformCollisionSystem) Run(scene blueprint.Scene, dt float64) error {
	platformTerrainQuery := warehouse.Factory.NewQuery().And(components.PlatformTag)
	platformCursor := scene.NewCursor(platformTerrainQuery)
	playerCursor := scene.NewCursor(blueprint.Queries.InputBuffer)

	for range platformCursor.Next() {
		for range playerCursor.Next() {
			err := s.resolve(scene, platformCursor, playerCursor)
			if err != nil {
				return err
			}
			playerPos := blueprintspatial.Components.Position.GetFromCursor(playerCursor)
			s.trackPosition(playerPos.Two)
		}
	}
	return nil
}

func (s *PlayerPlatformCollisionSystem) resolve(scene blueprint.Scene, platformCursor, playerCursor *warehouse.Cursor) error {
	// Get the player state
	playerShape := blueprintspatial.Components.Shape.GetFromCursor(playerCursor)
	playerPosition := blueprintspatial.Components.Position.GetFromCursor(playerCursor)
	playerDynamics := blueprintmotion.Components.Dynamics.GetFromCursor(playerCursor)

	// Get the platform state
	platformShape := blueprintspatial.Components.Shape.GetFromCursor(platformCursor)
	platformPosition := blueprintspatial.Components.Position.GetFromCursor(platformCursor)
	platformDynamics := blueprintmotion.Components.Dynamics.GetFromCursor(platformCursor)

	// Check for collision
	if ok, collisionResult := spatial.Detector.Check(
		*playerShape, *platformShape, playerPosition.Two, platformPosition.Two,
	); ok {

		// Check if any of the past player positions indicate the player was above the platform
		platformTop := platformShape.Polygon.WorldVertices[0].Y

		playerWasAbove := s.checkAnyPlayerPositionWasAbove(platformTop, playerShape.LocalAAB.Height)

		// We only want to resolve collisions when:
		// 1. The player is falling (vel.Y > 0)
		// 2. The collision is with the top of the platform
		// 3. The player was above the platform at some point (within n ticks)
		if playerDynamics.Vel.Y > 0 && collisionResult.IsTopB() && playerWasAbove {

			motion.Resolver.Resolve(
				&playerPosition.Two,
				&platformPosition.Two,
				playerDynamics,
				platformDynamics,
				collisionResult,
			)

			// Standard onGround handling
			currentTick := scene.CurrentTick()

			// If not grounded, enqueue onGround with values
			playerAlreadyGrounded, onGround := components.OnGroundComponent.GetFromCursorSafe(playerCursor)

			if !playerAlreadyGrounded {
				playerEntity, _ := playerCursor.CurrentEntity()
				err := playerEntity.EnqueueAddComponentWithValue(
					components.OnGroundComponent,
					components.OnGround{LastTouch: currentTick, Landed: currentTick},
				)
				if err != nil {
					return err
				}
			} else {
				onGround.LastTouch = currentTick
			}
		}
	}
	return nil
}

// trackPosition adds a position to the history and ensures only the last N are kept
func (s *PlayerPlatformCollisionSystem) trackPosition(pos vector.Two) {
	// Add the new position
	s.playerLastPositions = append(s.playerLastPositions, pos)

	// If we've exceeded our max, remove the oldest position
	if len(s.playerLastPositions) > s.maxPositionsToTrack {
		s.playerLastPositions = s.playerLastPositions[1:]
	}
}

// checkAnyPlayerPositionWasAbove checks if the player was above a non-rotated platform in any historical position
func (s *PlayerPlatformCollisionSystem) checkAnyPlayerPositionWasAbove(platformTop float64, playerHeight float64) bool {
	if len(s.playerLastPositions) == 0 {
		return false
	}

	// Check all stored positions to see if the player was above in any of them
	for _, pos := range s.playerLastPositions {
		playerBottom := pos.Y + playerHeight/2
		if playerBottom <= platformTop {
			return true // Found at least one position where player was above
		}
	}

	return false
}
