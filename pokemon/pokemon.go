package pokemon

import (
  "math/rand"
  "fmt"
)

const population = 20

type PokeType struct {
  Name string
  Weaknesses []*PokeType
  Resistances []*PokeType
}

func (b PokeType) String() string {
  return b.Name
}

type Stats struct {
  HP int
  Atk int
  Def int
  SpAtk int
  SpDef int
  Speed int
}

type Pokemon struct {
  Primary *PokeType
  Secondary *PokeType
  Stats Stats
}

func (p Pokemon) String() string {
  return fmt.Sprintf("%s, %s", p.Primary.Name, p.Secondary.Name)
}

var (
  Water PokeType
  Fire PokeType
  Grass PokeType
  Steel PokeType
  Bug PokeType
  Fairy PokeType

  Dummies [population]*Pokemon
)

func generatePokemon() *Pokemon {
  stats := Stats{ 
    HP: rand.Intn(255),
    Atk: rand.Intn(255),
    Def: rand.Intn(255),
    SpAtk: rand.Intn(255),
    SpDef: rand.Intn(255),
    Speed: rand.Intn(255),
  }

  types := [...]*PokeType{
    &Fire, &Water, &Grass,
    &Steel, &Bug, &Fairy,
  }

  primary := types[rand.Intn(len(types))]
  secondary := types[rand.Intn(len(types))]

  return &Pokemon{ Primary: primary, Secondary: secondary, Stats: stats }
}

func init() {
  Water.Name = "Water"
  Fire.Name = "Fire"
  Grass.Name = "Grass"
  Steel.Name = "Steel"
  Bug.Name = "Bug"
  Fairy.Name = "Fairy"

  Water.Weaknesses = []*PokeType{ &Grass }
  Fire.Weaknesses = []*PokeType{ &Water }
  Grass.Weaknesses = []*PokeType{ &Fire, &Bug }
  Steel.Weaknesses = []*PokeType{ &Fire }
  Bug.Weaknesses = []*PokeType{ &Fire }
  Fairy.Weaknesses = []*PokeType{ &Steel }

  Water.Resistances = []*PokeType{ &Water, &Fire }
  Fire.Resistances = []*PokeType{ &Fire, &Grass }
  Grass.Resistances = []*PokeType{ &Grass, &Water }
  Steel.Resistances = []*PokeType{ &Steel, &Bug, &Fairy }
  Bug.Resistances = []*PokeType{ &Grass }
  Fairy.Resistances = []*PokeType{ &Bug }

  for i := 0; i < population; i++ {
    Dummies[i] = generatePokemon()
  } 
}