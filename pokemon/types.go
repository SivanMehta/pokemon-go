package pokemon

import (
  "time"
  "math/rand"
)

type PokeTypes []*PokeType

type PokeType struct {
  Name string
  Weaknesses PokeTypes
  Resistances PokeTypes
  Immunities PokeTypes
}

func (b PokeType) String() string {
  return b.Name
}

var (
  Water PokeType
  Fire PokeType
  Grass PokeType
  Steel PokeType
  Bug PokeType
  Fairy PokeType
  Dark PokeType
  Ghost PokeType
  Electric PokeType
  Flying PokeType
  Ice PokeType
  Ground PokeType
  Dragon PokeType
  Rock PokeType
  Fighting PokeType
  Poison PokeType
  Psychic PokeType
  Normal PokeType

  PossibleTypes [18]*PokeType
)

// collapse a list basic list of types and return the unpacked version
func unpack(types ...*PokeType) PokeTypes {
  return types
}

func init() {
  rand.Seed(time.Now().UTC().UnixNano())
  Water.Name = "Water"
  Fire.Name = "Fire"
  Grass.Name = "Grass"
  Steel.Name = "Steel"
  Bug.Name = "Bug"
  Fairy.Name = "Fairy"
  Dark.Name = "Dark"
  Ghost.Name = "Ghost"
  Electric.Name = "Electric"
  Flying.Name = "Flying"
  Ice.Name = "Ice"
  Ground.Name = "Ground"
  Dragon.Name = "Dragon"
  Rock.Name = "Rock"
  Fighting.Name = "Fighting"
  Poison.Name = "Poison"
  Psychic.Name = "Psychic"
  Normal.Name = "Normal"

  Water.Weaknesses = unpack(&Grass, &Electric)
  Fire.Weaknesses = unpack(&Water, &Ground, &Rock)
  Grass.Weaknesses = unpack(&Fire, &Bug, &Flying) 
  Steel.Weaknesses = unpack(&Fire, &Ground, &Fighting)
  Bug.Weaknesses = unpack(&Fire, &Flying, &Rock)
  Fairy.Weaknesses = unpack(&Steel, &Poison)
  Dark.Weaknesses = unpack(&Bug, &Fairy)
  Ghost.Weaknesses = unpack(&Ghost, &Dark)
  Electric.Weaknesses = unpack(&Ground)
  Flying.Weaknesses = unpack(&Electric, &Ice, &Rock)
  Ice.Weaknesses = unpack(&Fire, &Steel, &Rock, &Fighting)
  Ground.Weaknesses = unpack(&Water, &Grass, &Ice)
  Dragon.Weaknesses = unpack(&Dragon, &Fairy, &Ice)
  Rock.Weaknesses = unpack(&Grass, &Ground, &Steel)
  Fighting.Weaknesses = unpack(&Flying, &Fairy, &Psychic)
  Poison.Weaknesses = unpack(&Ground, &Psychic)
  Psychic.Weaknesses = unpack(&Dark, &Ghost, &Bug)
  Normal.Weaknesses = unpack(&Fighting)

  Water.Resistances = unpack(&Water, &Fire, &Ice)
  Fire.Resistances = unpack(&Fire, &Grass, &Steel, &Fairy)
  Grass.Resistances = unpack(&Grass, &Water, &Electric, &Ground)
  Steel.Resistances = unpack(&Steel, &Bug, &Fairy, &Flying, &Ice, &Dragon, &Psychic, &Normal)
  Bug.Resistances = unpack(&Grass, &Fighting)
  Fairy.Resistances = unpack(&Bug, &Fighting, &Dark)
  Dark.Resistances = unpack(&Bug, &Dark)
  Ghost.Resistances = unpack(&Bug, &Poison)
  Electric.Resistances = unpack(&Electric, &Flying)
  Flying.Resistances = unpack(&Grass, &Bug, &Fighting)
  Ice.Resistances = unpack(&Ice)
  Ground.Resistances = unpack(&Rock, &Poison)
  Dragon.Resistances = unpack(&Water, &Fire, &Grass, &Electric)
  Rock.Resistances = unpack(&Flying, &Fire, &Poison, &Normal)
  Fighting.Resistances = unpack(&Bug, &Rock)
  Poison.Resistances = unpack(&Grass, &Fighting, &Poison, &Bug)
  Psychic.Resistances = unpack(&Fighting, &Psychic)

  Steel.Immunities = unpack(&Poison)
  Fairy.Immunities = unpack(&Dragon)
  Dark.Immunities = unpack(&Psychic)
  Ghost.Immunities = unpack(&Fighting, &Normal)
  Flying.Immunities = unpack(&Ground)
  Ground.Immunities = unpack(&Electric)
  Normal.Immunities = unpack(&Ghost)

  PossibleTypes = [...]*PokeType{
    &Fire, &Water, &Grass,
    &Steel, &Bug, &Fairy,
    &Dark, &Ghost, &Electric,
    &Flying, &Ice, &Ground,
    &Dragon, &Rock, &Fighting,
    &Poison, &Psychic, &Normal,
  }

  for i := 0; i < population; i++ {
    Population[i] = generatePokemon()
  }
}