package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Player представляет игрока
type Player struct {
	Name     string
	Board    *Board
	Opponent *Player
}

type PlayerImpl interface {
	GiveUp()
	DoMove(x int, y int) (hit bool, sunk bool, message string)
	TakeMove(x int, y int) (hit bool, sunk bool, message string)
}

func (p *Player) GiveUp() {
	fmt.Printf("Игрок %s сдался!\n", p.Name)
}

func (p *Player) DoMove(x int, y int) (hit bool, sunk bool, message string) {
	for _, ship := range p.Opponent.Board.Ships {
		if ship.HandleShoot(x, y) {
			hit = true
			if ship.Sunk {
				sunk = true
				return true, true, "Потопил!"
			}
			return true, false, "Попал!"
		}
	}
	return false, false, "Мимо!"
}

func (p *Player) TakeMove(x int, y int) (hit bool, sunk bool, message string) {
	// Просто проверяем попадание в наши корабли
	for _, ship := range p.Board.Ships {
		if ship.HandleShoot(x, y) {
			hit = true
			if ship.Sunk {
				sunk = true
				return true, true, "Потопили ваш корабль!"
			}
			return true, false, "Попали в ваш кораль!"
		}
	}
	return false, false, "Промахнулись!"
}

// NewPlayer создает нового игрока
func NewPlayer(name string) *Player {
	return &Player{
		Name:  name,
		Board: *Board.NewBoard(),
	}
}

// placeShip размещает один корабль заданного размера
func (p *Player) placeShip(size int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		isHorizontal := r.Intn(2) == 0

		var x, y int
		if isHorizontal {
			x = r.Intn(10)
			y = r.Intn(11 - size)
		} else {
			x = r.Intn(11 - size)
			y = r.Intn(10)
		}

		if p.Board.CanPlaceShip(x, y, size, orientation) {
			p.Board.AddShip(x, y, size, orientation)
			break
		}
	}
}
