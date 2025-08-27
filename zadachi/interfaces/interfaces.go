package interfaces

// Renderable интерфейс для отрисовки
type Renderable interface {
	Render()
}

// ShipPlacer интерфейс для размещения кораблей
type ShipPlacer interface {
	PlaceShipsRandomly()
}
