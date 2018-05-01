package battle

import (
  "github.com/SivanMehta/pokemon-go/pokemon"
)

const basePower = 100

// maximizing damage output considering the stats and typing
// of the attacker and the defender
// returns and int representing the highest possible damange
func optimalAttack(attacking *pokemon.Pokemon, defending *pokemon.Pokemon) int {
  return 10

  // use the type that will do the most damage (remember that all moves are STAB)
  // use the physical/special option that will do the most damage
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
      // break in the middle of the turn because B has fainted
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