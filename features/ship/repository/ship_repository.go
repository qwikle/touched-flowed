package repository

import (
	MapEntity "touchedFlowed/features/map/entities"
	"touchedFlowed/features/ship/entities"
)

type ShipRepository interface {
	SetShipPosition(ship *entities.Ship, position MapEntity.Position) (*entities.Ship, error)
	GetFleet() []entities.Fleet
}
