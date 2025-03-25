package clientsystems

import (
	"github.com/TheBitDrifter/coldbrew"
	coldbrew_clientsystems "github.com/TheBitDrifter/coldbrew/clientsystems"
)

var DefaultClientSystems = []coldbrew.ClientSystem{
	&CameraFollowerSystem{},
	&coldbrew_clientsystems.BackgroundScrollSystem{},
}
