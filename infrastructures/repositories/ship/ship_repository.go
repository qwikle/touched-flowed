package ship

import (
	"errors"
	MapEntity "touchedFlowed/features/map/entities"
	"touchedFlowed/features/ship/entities"
	"touchedFlowed/features/ship/repository"
)

type Repository struct {
	repository *repository.ShipRepository
}

func (r Repository) SetShipPosition(ship *entities.Ship, position MapEntity.Position) (*entities.Ship, error) {
	if position.X > MapEntity.MaxX || position.Y > MapEntity.MaxY || position.X < MapEntity.MinX || position.Y < MapEntity.MinY {
		return nil, errors.New("position out of map")
	}
	ship.Position = position
	return ship, nil
}

func (r Repository) GetFleet() []entities.Fleet {
	panic("implement me")
}

func NewRedisShipRepository() repository.ShipRepository {
	return &Repository{}
}
