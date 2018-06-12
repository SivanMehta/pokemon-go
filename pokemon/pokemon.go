package pokemon

import (
  "math"
  "math/rand"
  "fmt"
)

const (
  population = 100
  mutationRate = .1
  BST = 600
)

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
  return fmt.Sprintf("%s,%s,%d,%d,%d,%d,%d,%d", p.Primary.Name, p.Secondary.Name, p.Stats.HP, p.Stats.Atk, p.Stats.Def, p.Stats.SpAtk, p.Stats.SpDef, p.Stats.Speed)
}

func combine(a, b int) int {
  noise := int(math.Ceil(rand.NormFloat64() * 5) + .5)
  avg := noise + (a + b) / 2
  if avg < 0 {
    return 0
  } else {
    return avg
  }
}

func scale(stats Stats, bst int) Stats {
  total := stats.HP +
           stats.Atk +
           stats.Def +
           stats.SpAtk +
           stats.SpDef +
           stats.Speed

  return Stats{
    HP: stats.HP * bst / total,
    Atk: stats.Atk * bst / total,
    Def: stats.Def * bst / total,
    SpAtk: stats.SpAtk * bst / total,
    SpDef: stats.SpDef * bst / total,
    Speed: stats.Speed * bst / total,
  }
}

// Basic combination of two pokemon
//  - 2 out of the 4 types between the parents
//    - mutation might randomly pick a type
//  - avg of parents' stats + some noise so we don't regress
//
// For simplicity's sake, we're assuming that every pair is compatible
func (p Pokemon) Breed(partner *Pokemon) *Pokemon {
  possibleTypes := [4]*PokeType{ p.Primary, p.Secondary, partner.Primary, p.Secondary }
  var baby Pokemon

  // choose one of parents' types, potentially mutating
  if rand.Float64() > mutationRate {
    baby.Primary = possibleTypes[rand.Intn(len(possibleTypes))]
  } else {
    baby.Primary = PossibleTypes[rand.Intn(len(PossibleTypes))]
  }

  // choose one of parents' types, potentially mutating
  if rand.Float64() > mutationRate {
    baby.Secondary = possibleTypes[rand.Intn(len(possibleTypes))]
  } else {
    baby.Secondary = PossibleTypes[rand.Intn(len(PossibleTypes))]
  }

  combinedStats := Stats{
    HP: combine(p.Stats.HP, partner.Stats.HP),
    Atk: combine(p.Stats.Atk, partner.Stats.Atk),
    Def: combine(p.Stats.Def, partner.Stats.Def),
    SpAtk: combine(p.Stats.SpAtk, partner.Stats.SpAtk),
    SpDef: combine(p.Stats.SpDef, partner.Stats.SpDef),
    Speed: combine(p.Stats.Speed, partner.Stats.Speed),
  }
  baby.Stats = scale(combinedStats, BST)

  return &baby
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

  for i := 0; i < len(p.Primary.Immunities); i ++ {
    if(attack == p.Primary.Immunities[i]) {
      base = 0.0
    }
  }

  for i := 0; i < len(p.Secondary.Immunities); i ++ {
    if(attack == p.Secondary.Immunities[i]) {
      base = 0.0
    }
  }

  return base
}

var (
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

  // scale stats so everything is always BST = 600
  stats = scale(stats, BST)

  primary := PossibleTypes[rand.Intn(len(PossibleTypes))]
  secondary := PossibleTypes[rand.Intn(len(PossibleTypes))]

  return &Pokemon{ Primary: primary, Secondary: secondary, Stats: stats }
}
