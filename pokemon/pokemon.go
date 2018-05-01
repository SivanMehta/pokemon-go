package pokemon

type PokeType struct {
  name string
  weaknesses []*PokeType
  resistances []*PokeType
}

func (b PokeType) String() string {
  return b.name
}

var (
  Water PokeType
  Fire PokeType
  Grass PokeType
  Steel PokeType
  Bug PokeType
  Fairy PokeType
)

func init() {
  Water.name = "Water"
  Fire.name = "Fire"
  Grass.name = "Grass"
  Steel.name = "Steel"
  Bug.name = "Bug"
  Fairy.name = "Fairy"

  Water.weaknesses = []*PokeType{ &Grass }
  Fire.weaknesses = []*PokeType{ &Water }
  Grass.weaknesses = []*PokeType{ &Fire, &Bug }
  Steel.weaknesses = []*PokeType{ &Fire }
  Bug.weaknesses = []*PokeType{ &Fire }
  Fairy.weaknesses = []*PokeType{ &Steel }

  Water.resistances = []*PokeType{ &Water, &Fire }
  Fire.resistances = []*PokeType{ &Fire, &Grass }
  Grass.resistances = []*PokeType{ &Grass, &Water }
  Steel.resistances = []*PokeType{ &Steel, &Bug, &Fairy }
  Bug.resistances = []*PokeType{ &Grass }
  Fairy.resistances = []*PokeType{ &Bug }
}