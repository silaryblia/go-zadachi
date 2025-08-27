package main

import (
	"battleship/game"
)

func main() {
	// Создаем игру
	game := game.NewGame()

	// Запускаем игру
	game.Start()

	// Останавливаем игру
	defer game.Stop()
}

/*
	// Create player
	p := player.NewPlayer("Player 1")

	// Размещаем корабли случайным образом
	p.PlaceShipsRandomly()

	// Выводим поле с кораблями
	fmt.Printf("Поле игрока: %s\n", p.Name)
	p.Render()
*/
