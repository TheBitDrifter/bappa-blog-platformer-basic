package components

import "github.com/TheBitDrifter/warehouse"

type OnGround struct {
	Landed, LastTouch int
}

var OnGroundComponent = warehouse.FactoryNewComponent[OnGround]()

type IgnorePlatform struct {
	Items [5]struct {
		LastActive int
		EntityID   int
		Recycled   int
	}
}

var IgnorePlatformComponent = warehouse.FactoryNewComponent[IgnorePlatform]()
