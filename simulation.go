package main

import (
  "log"
  "fmt"
  // "math/rand"
  "sort"
  "github.com/SivanMehta/pokemon-go/battle"
  "github.com/SivanMehta/pokemon-go/pokemon"
)

type sortable struct {
  Origin int
  Fitness float64
}

func (s sortable) String() string {
  return fmt.Sprintf("(%d %2f)", s.Origin, s.Fitness)
}

type sorter []*sortable

func (s sorter) Len() int {
    return len(s)
}
func (s sorter) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s sorter) Less(i, j int) bool {
    return s[i].Fitness > s[j].Fitness
}

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

  results := make([]*sortable, census)

  for i, _ := range pokemon.Population {
    fitness := make(chan float64)
    go battle.Fitness(i, fitness)
    results[i] = &sortable{ Origin: i, Fitness: <-fitness }
  }

  sort.Sort(sorter(results))

  log.Println(results)
  return nextGeneration
}

func main() {
  generation()
}