package components

import "github.com/TheBitDrifter/warehouse"

type OnGround struct {
	Landed, LastTouch int
}

var OnGroundComponent = warehouse.FactoryNewComponent[OnGround]()
