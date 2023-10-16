package entities

import "touchedFlowed/features/ship/entities"

const (
	OrientationX = "x"
	OrientationY = "y"
	MaxX         = 9
	MaxY         = 9
	MinX         = 0
	MinY         = 0
)

type Map struct {
	fleet    entities.Fleet
	missiles []Position
}
