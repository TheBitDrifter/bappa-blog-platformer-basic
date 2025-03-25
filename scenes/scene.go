package scenes

import "github.com/TheBitDrifter/blueprint"

type Scene struct {
	Name          string
	Plan          blueprint.Plan
	Width, Height int
}
