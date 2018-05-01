package main

import (
  "log"
  // "math/rand"
  "github.com/SivanMehta/pokemon-go/battle"
  "github.com/SivanMehta/pokemon-go/pokemon"
)

// First, Have every pokemon battle every other one
// and accumulate the results into their fitness attribute
// each battle can be run inside a goroutine because battle.Battle
// takes a channel that "returns" the result of each battle
// that way we can avoid using a mutex locking the counter appropriately
//
// Then, we we sort each pokemon by their fitness and drop the bottom 50%
//
// Finally, we randomly breed the remaining pokemon to fill out the remaining population
func generation() []*pokemon.Pokemon {
  return make([]*pokemon.Pokemon, 10)
}

func main() {
  battles := make(chan int, 1)
  battle.Battle(pokemon.Dummies[0], pokemon.Dummies[1], battles)

  result := <-battles

  log.Println(result)
}