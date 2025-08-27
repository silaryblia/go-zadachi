package main

import (
	"fmt"
)

// Game представляет игру
type Game struct {
	Player  *Player
	Running bool
}

type GameImpls interface {
	Start()
	Stop()
}

// NewGame создает новую игру
func NewGame() *Game {
	return &Game{
		Player: Player.NewPlayer("Player 1"),
	}
}

// Start начинает игру
func (g *Game) Start() {
	fmt.Println("🚀 Игра начинается!")
	g.Running = true

	// Размещаем корабли
	fmt.Println("🛳️  Размещаем корабли...")
	g.Player.PlaceShipsRandomly()

	// Проверяем корректность размещения
	if g.Player.Board.IsValidPlacement() {
		fmt.Println("✅ Корабли размещены корректно!")
		fmt.Printf("📊 Количество кораблей: %d\n", g.Player.Board.GetShipsCount())
		fmt.Printf("📊 Количество палуб: %d\n", g.Player.Board.GetShipCells())
	} else {
		fmt.Println("❌ Ошибка размещения кораблей!")
		g.Stop()
		return
	}

	// Показываем поле
	fmt.Printf("\n🎯 Поле игрока:\n")
	g.Player.B()

	fmt.Println("🎮 Игра готова к началу!")
}

// Stop останавливает игру
func (g *Game) Stop() {
	fmt.Println("🛑 Игра завершена!")
	g.Running = false
}
