package coresystems

import (
	"github.com/TheBitDrifter/blueprint"
	tteo_coresystems "github.com/TheBitDrifter/tteokbokki/coresystems"
)

var DefaultCoreSystems = []blueprint.CoreSystem{
	GravitySystem{},
	FrictionSystem{},
	PlayerMovementSystem{},
	tteo_coresystems.IntegrationSystem{},
	tteo_coresystems.TransformSystem{},
	PlayerBlockCollisionSystem{},
	NewPlayerPlatformCollisionSystem(),
	OnGroundClearingSystem{},
	IgnorePlatformClearingSystem{},
}
