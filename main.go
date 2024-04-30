package main

import (
	"fmt"

	"github.com/Alieksieiev0/space-invaders/internal/game"
)

func main() {
	fmt.Println(game.NewGameManager().Start())
}
