package battle

import (
  // "log"
  "math/rand"
  "github.com/SivanMehta/pokemon-go/pokemon"
)

const basePower = 100

// maximizing damage output considering the stats and typing
// of the attacker and the defender
//
// returns and int representing the highest possible damange
func optimalAttack(attacker *pokemon.Pokemon, defender *pokemon.Pokemon) float64 {
  var base float64

  // use the type that will do the most damage (remember that all moves are STAB)
  primary := defender.Multiplier(attacker.Primary, basePower)
  secondary := defender.Multiplier(attacker.Secondary, basePower)

  if primary > secondary {
    base = primary
  } else {
    base = secondary
  }

  atk := pokemon.Stat(attacker.Stats.Atk)
  def := pokemon.Stat(defender.Stats.Atk)
  spAtk := pokemon.Stat(attacker.Stats.SpAtk)
  spDef := pokemon.Stat(defender.Stats.SpDef)

  // use the actual battle damage formula to calculate damage
  // based on the physical / special split
  physicalDamage := (.84 * base * atk / def) + 2.0
  specialDamage := (.84 * base * spAtk / spDef) + 2.0

  if physicalDamage > specialDamage {
    return physicalDamage
  } else {
    return specialDamage
  }
}

// takes 2 pokemon and returns the difference in hp
// at the end of the battle. Opponents will always use the
// highest damage possible and just take turns using that
// most damaging attack
//
// if the result > 0, a has won, if it is < 0, b has won
func Battle(a *pokemon.Pokemon, b *pokemon.Pokemon, done chan<- float64) {
  // log.Println("Battling", a, b)
  hpA := pokemon.HpStat(a.Stats.HP)
  hpB := pokemon.HpStat(b.Stats.HP)
  atkA := optimalAttack(a, b)
  atkB := optimalAttack(b, a)
  aFaster := a.Stats.Speed > b.Stats.Speed

  for hpA > 0 && hpB > 0 {
    multiplier := (217 + rand.Float64() * 38.0) / 255.0
    if rand.Intn(16) < 1 {
      multiplier *= 1.5
    }

    if(aFaster) {
      hpB -= atkA * multiplier
      if(hpB <= 0) {
      // break in the middle of the turn because b has fainted
        break;
      }
      hpA -= atkB * multiplier
    } else {
      hpA -= atkB * multiplier
      hpB -= atkA * multiplier
    }
  }

  done <- hpA - hpB
}