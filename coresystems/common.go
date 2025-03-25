package coresystems

import (
	"github.com/TheBitDrifter/blueprint"
	tteo_coresystems "github.com/TheBitDrifter/tteokbokki/coresystems"
)

var DefaultCoreSystems = []blueprint.CoreSystem{
	FrictionSystem{},
	PlayerMovementSystem{},
	tteo_coresystems.IntegrationSystem{}, // Update velocities and positions
	tteo_coresystems.TransformSystem{},   // Update collision shapes
}
