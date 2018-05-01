package battle

import (
  "github.com/SivanMehta/pokemon-go/pokemon"
)

// takes 2 pokemon and returns the difference in hp
// at the end of the battle. Very basically
// if the result > 0, a has won, if it is < 0, b has won
func Battle(a *pokemon.Pokemon, b *pokemon.Pokemon, done chan int) {
  hpA := a.Stats.HP
  hpB := b.Stats.HP

  for hpA > 0 && hpB > 0 {
    hpB--;
  }

  done <- hpA - hpB
}