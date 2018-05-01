package main

import (
  "log"
  "github.com/SivanMehta/pokemon-go/battle"
  "github.com/SivanMehta/pokemon-go/pokemon"
)

func main() {
  result := battle.Add(1, 2)
  log.Println(result)

  log.Println(pokemon.Dummies)
}