package main

import (
  "log"
  "fmt"
  "math/rand"
  "sort"
  "github.com/SivanMehta/pokemon-go/battle"
  "github.com/SivanMehta/pokemon-go/pokemon"
)

const (
  generations = 10
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

// Randomly breed pokemon until the population is filled out
func Breed(nextGeneration []*pokemon.Pokemon) {
  census := len(nextGeneration)
  population := (census / 2)
  left := population

  for left < census {
    father := nextGeneration[rand.Intn(population - 1)]
    mother := nextGeneration[rand.Intn(population - 1)]
    baby := father.Breed(mother)
    nextGeneration[left] = baby

    left += 1
  }

  // place new pokemon into pokemon.Population
  // essentially copying nextGeneration into pokemon.Population

  for i, poke := range nextGeneration {
    pokemon.Population[i] = poke
  }
}

// does one generation of a genetic algorithm, modifying pokemon.Pokemon in place
// population: a group of a valid Pokemon
// fitness: sum total of battle scores against all other pokemon
// cutoff: arbitrarily cut off 1/2 of population
func Generation() {
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

  // breed to fill out the rest
  Breed(nextGeneration)
}

func main() {
  for i := 0; i < generations; i++ {
    Generation()
    log.Println(pokemon.Population)
  }
}
