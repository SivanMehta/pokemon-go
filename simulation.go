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
  census := len(pokemon.Population)
  nextGeneration := make([]*pokemon.Pokemon, census)

  for _, poke := range pokemon.Population {
    total := float64(0)

    for _, otherPoke := range pokemon.Population {
      channel := make(chan float64)
      go battle.Battle(poke, otherPoke, channel)
      total = <- channel
    }

    fitness := total / float64(census)
    log.Println(fitness)
  }
  return nextGeneration
}

func main() {
  generation()
}