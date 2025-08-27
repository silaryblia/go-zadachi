package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Cell представляет клетку поля
type Cell struct {
	X, Y    int
	HasShip bool
}

// Ship представляет корабль
type Ship struct {
	Size  int
	Cells []*Cell
	Sunk  bool
}

// Board представляет игровое поле
type Board struct {
	Cells [10][10]*Cell
	Ships []*Ship
}

// Player представляет игрока
type Player struct {
	Name  string
	Board *Board
}

// Renderable интерфейс для отрисовки
type Renderable interface {
	Render()
}

// ShipPlacer интерфейс для размещения кораблей
type ShipPlacer interface {
	PlaceShipsRandomly()
}

// NewCell создает новую клетку
func NewCell(x, y int) *Cell {
	return &Cell{X: x, Y: y, HasShip: false}
}

// NewBoard создает новое поле
func NewBoard() *Board {
	board := &Board{}

	// Инициализируем клетки
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			board.Cells[i][j] = NewCell(i, j)
		}
	}
	return board
}

// NewPlayer создает нового игрока
func NewPlayer(name string) *Player {
	return &Player{
		Name:  name,
		Board: NewBoard(),
	}
}

// Render отрисовывает поле игрока
func (b *Board) Render() {
	fmt.Println("  A B C D E F G H I J")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i+1)
		for j := 0; j < 10; j++ {
			cell := b.Cells[i][j]
			if cell.HasShip {
				fmt.Print("■ ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// PlaceShipsRandomly размещает все корабли случайным образом
func (p *Player) PlaceShipsRandomly() {
	ships := []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1} // Размеры кораблей

	for _, size := range ships {
		p.placeShip(size)
	}
}

// placeShip размещает один корабль заданного размера
func (p *Player) placeShip(size int) {
	rand.Seed(time.Now().UnixNano())

	for {
		// Случайное направление (0-горизонт, 1 - вертикаль)
		isHorizontal := rand.Intn(2) == 0

		// Случайная начальная позиция
		var x, y int
		if isHorizontal {
			x = rand.Intn(10)
			y = rand.Intn(11 - size)
		} else {
			x = rand.Intn(11 - size)
			y = rand.Intn(10)
		}

		// Проверяем, можно ли разместить корабль
		if p.canPlaceShip(x, y, size, isHorizontal) {
			// размещаем корабль
			ship := &Ship{Size: size}
			for i := 0; i < size; i++ {
				var cell *Cell
				if isHorizontal {
					cell = p.Board.Cells[x][y+i]
				} else {
					cell = p.Board.Cells[x+i][y]
				}
				cell.HasShip = true
				ship.Cells = append(ship.Cells, cell)
			}
			p.Board.Ships = append(p.Board.Ships, ship)
			break
		}
	}
}

// canPlaceShip проверяет, можно ли разместить корабль
func (p *Player) canPlaceShip(x, y, size int, isHorizontal bool) bool {
	for i := -1; i <= size; i++ {
		for j := -1; j <= 1; j++ {
			var checkX, checkY int

			if isHorizontal {
				checkX = x + j
				checkY = y + i
			} else {
				checkX = x + i
				checkY = y + j
			}

			if checkX < 0 || checkX >= 10 || checkY < 0 || checkY >= 10 {
				continue
			}

			if p.Board.Cells[checkX][checkY].HasShip {
				return false
			}
		}
	}
	return true
}

func main() {
	// Create player
	player := NewPlayer("Player 1")

	// Размещаем корабли случайным образом
	player.PlaceShipsRandomly()

	// Выводим поле с кораблями
	fmt.Printf("Поле игрока: %s\n", player.Name)
	player.Board.Render()
}
