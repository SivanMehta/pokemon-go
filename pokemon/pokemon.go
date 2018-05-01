package pokemon

import (
  "math/rand"
  "fmt"
  "time"
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

// prints out primary and secondary type of a pokemon
func (p Pokemon) String() string {
  return fmt.Sprintf("%s, %s", p.Primary.Name, p.Secondary.Name)
}

// outputs damage multiplier of an attack, given a basepower
func (p Pokemon) Multiplier(attack *PokeType, basepower int) int {
  base := basepower
  for i := 0; i < len(p.Primary.Weaknesses); i ++ {
    if(attack == p.Primary.Weaknesses[i]) {
      base *= 2
    }
  }

  for i := 0; i < len(p.Secondary.Weaknesses); i ++ {
    if(attack == p.Secondary.Weaknesses[i]) {
      base *= 2
    }
  }

  for i := 0; i < len(p.Primary.Resistances); i ++ {
    if(attack == p.Primary.Resistances[i]) {
      base /= 2
    }
  }

  for i := 0; i < len(p.Secondary.Resistances); i ++ {
    if(attack == p.Secondary.Resistances[i]) {
      base /= 2
    }
  }

  return base
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

// for our purposes evs and ivs are all always maxed out
// and we always have a neutral nature at level 100
func hpStat() int {
  base := rand.Intn(255)
  return ((base + 31) + 4) + 100 + 10
}

func stat() int {
  base := rand.Intn(255)
  return ((base + 31) + 4) + 5
}

func generatePokemon() *Pokemon {
  stats := Stats{ 
    HP: hpStat(),
    Atk: stat(),
    Def: stat(),
    SpAtk: stat(),
    SpDef: stat(),
    Speed: stat(),
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
  rand.Seed(time.Now().UTC().UnixNano())
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