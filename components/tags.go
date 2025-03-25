package components

import "github.com/TheBitDrifter/warehouse"

var (
	BlockTerrainTag = warehouse.FactoryNewComponent[struct{}]()
	PlatformTag     = warehouse.FactoryNewComponent[struct{}]()
)
