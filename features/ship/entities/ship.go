package entities

import "touchedFlowed/features/map/entities"

const (
	submarine  = 1
	destroyer  = 2
	cruiser    = 3
	battleship = 4
)

type Ship struct {
	Position     entities.Position
	Orientation  string
	HitPositions []entities.Position
	IsSunk       bool
	Case         int
}

type Fleet struct {
	Ships []Ship
}

func NewFleet() *Fleet {
	return &Fleet{
		Ships: []Ship{
			{
				Position: entities.Position{
					X: 0,
					Y: 0,
				},
				Orientation:  entities.OrientationX,
				HitPositions: []entities.Position{},
				IsSunk:       false,
				Case:         submarine,
			},
			{
				Position: entities.Position{
					X: 0,
					Y: 0,
				},
				Orientation:  entities.OrientationX,
				HitPositions: []entities.Position{},
				IsSunk:       false,
				Case:         destroyer,
			},
			{
				Position: entities.Position{
					X: 0,
					Y: 0,
				},
				Orientation:  entities.OrientationX,
				HitPositions: []entities.Position{},
				IsSunk:       false,
				Case:         cruiser,
			},
			{
				Position: entities.Position{
					X: 0,
					Y: 0,
				},
				Orientation:  entities.OrientationX,
				HitPositions: []entities.Position{},
				IsSunk:       false,
				Case:         battleship,
			},
		},
	}
}
