package main

import (
  "log"
  "github.com/SivanMehta/pokemon-go/battle"
  "github.com/SivanMehta/pokemon-go/pokemon"
)

func main() {
  battles := make(chan int, 1)
  battle.Battle(pokemon.Dummies[0], pokemon.Dummies[1], battles)

  result := <-battles

  log.Println(result)
}