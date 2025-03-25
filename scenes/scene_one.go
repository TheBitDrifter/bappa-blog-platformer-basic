package scenes

import (
	"github.com/TheBitDrifter/blueprint"
	"github.com/TheBitDrifter/warehouse"
)

const SCENE_ONE_NAME = "scene one"

var SceneOne = Scene{
	Name:   SCENE_ONE_NAME,
	Plan:   sceneOnePlan,
	Width:  1600,
	Height: 500,
}

func sceneOnePlan(height, width int, sto warehouse.Storage) error {
	err := blueprint.NewParallaxBackgroundBuilder(sto).
		AddLayer("backgrounds/city/sky.png", 0.025, 0.025).
		AddLayer("backgrounds/city/far.png", 0.025, 0.05).
		AddLayer("backgrounds/city/mid.png", 0.1, 0.1).
		AddLayer("backgrounds/city/near.png", 0.2, 0.2).
		Build()

	err = NewPlayer(sto, 100, 100)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	return nil
}
