package pokemon

import (
  "math"
  "math/rand"
  "fmt"
  "time"
)

const population = 10

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
  fitness float64
}

// prints out primary and secondary type of a pokemon
func (p Pokemon) String() string {
  return fmt.Sprintf("%s, %s: %d %d %d %d %d %d", p.Primary.Name, p.Secondary.Name, 
    p.Stats.HP, p.Stats.Atk, p.Stats.Def, p.Stats.SpAtk, p.Stats.SpDef, p.Stats.Speed)
}

// outputs damage multiplier of an attack, given a basepower
func (p Pokemon) Multiplier(attack *PokeType, basepower float64) float64 {
  base := basepower
  for i := 0; i < len(p.Primary.Weaknesses); i ++ {
    if(attack == p.Primary.Weaknesses[i]) {
      base *= 2.0
    }
  }

  for i := 0; i < len(p.Secondary.Weaknesses); i ++ {
    if(attack == p.Secondary.Weaknesses[i]) {
      base *= 2.0
    }
  }

  for i := 0; i < len(p.Primary.Resistances); i ++ {
    if(attack == p.Primary.Resistances[i]) {
      base /= 2.0
    }
  }

  for i := 0; i < len(p.Secondary.Resistances); i ++ {
    if(attack == p.Secondary.Resistances[i]) {
      base /= 2.0
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

  Population [population]*Pokemon
)

// calculate the HP stat given a base stat
// for our purposes evs and ivs are all always maxed out
// and we always have a neutral nature at level 100
func HpStat(base int) float64 {
  return float64(((base + 31) * 2 + 4) + 100 + 10)
}

// calculate the actual stat given a base stat
// for our purposes evs and ivs are all always maxed out
// and we always have a neutral nature at level 100
func Stat(base int) float64 {
  return float64(((base + 31) * 2 + 4) + 5)
}

func stat() int {
  return int(math.Ceil(rand.NormFloat64() * 10 + 100) + .5)
}

func generatePokemon() *Pokemon {
  stats := Stats{ 
    HP: stat(),
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
    Population[i] = generatePokemon()
  } 
}