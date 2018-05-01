package battle

import (
  "github.com/SivanMehta/pokemon-go/pokemon"
)

const basePower = 12

// maximizing damage output considering the stats and typing
// of the attacker and the defender
//
// returns and int representing the highest possible damange
func optimalAttack(attacker *pokemon.Pokemon, defender *pokemon.Pokemon) int {
  var base int

  // use the type that will do the most damage (remember that all moves are STAB)
  primary := defender.Multiplier(attacker.Primary, basePower)
  secondary := defender.Multiplier(attacker.Secondary, basePower)

  if primary > secondary {
    base = primary
  } else {
    base = secondary
  }

  // use the actual battle damage formula to calculate damage
  // based on the physical / special split
  // TODO: change 1 to .84 and all these ints to floats
  physicalDamage := (1 * base * attacker.Stats.Atk / defender.Stats.Def) + 2
  specialDamage := (1 * base * attacker.Stats.SpAtk / defender.Stats.SpDef) + 2

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
func Battle(a *pokemon.Pokemon, b *pokemon.Pokemon, done chan int) {
  hpA := a.Stats.HP
  hpB := b.Stats.HP
  atkA := optimalAttack(a, b)
  atkB := optimalAttack(b, a)
  aFaster := a.Stats.Speed > b.Stats.Speed

  for hpA > 0 && hpB > 0 {
    if(aFaster) {
      hpB -= atkA
      if(hpB <= 0) {
      // break in the middle of the turn because b has fainted
        break;
      }
      hpA -= atkB
    } else {
      hpA -= atkB
      hpB -= atkA
    }
  }

  done <- hpA - hpB
}