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

// does one generation of a genetic algorithm, modifying pokemon.Pokemon in place
// population: a group of a valid Pokemon
// fitness: sum total of battle scores against all other pokemon
// cutoff: arbitrarily cut off 1/2 of population
// crossover / breeding:
//  - 2 out of the 4 types between the parents
//  - avg of parents' stats + some noise so we don't regress
func generation() {
  // setup general variable
  census := len(pokemon.Population)
  nextGeneration := make([]*pokemon.Pokemon, census)

  // concurrently gather fitness for each pokemon
  fitnessResults := make([]*sortable, census)
  for i, _ := range pokemon.Population {
    fitness := make(chan float64)
    go battle.Fitness(i, fitness)
    fitnessResults[i] = &sortable{ Origin: i, Fitness: <-fitness }
  }

  // cull half of the population
  sort.Sort(sorter(fitnessResults))
  for i, result := range fitnessResults {
    if i > census / 2 {
      break
    }
    nextGeneration[i] = pokemon.Population[result.Origin]
  }
  log.Println(nextGeneration)

  // breed to fill out the rest
  // replace pokemon.Pokemon IN PLACE
}

func main() {
  generation()
}