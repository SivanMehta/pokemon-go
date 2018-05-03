# Building

```sh
# generate a binary that contains the entire simulation
make build
```

# Running

```sh
# run the built binary
make simulation

# or you can run this to build and then immediately run the simulation
make dev
```

# Implementation Specific Details

While most of the genetic-algorithm specific code, is in `simulation.go`, everything specific this domain is defined in this section. According to [Wikipedia], I need the to define the following attributes to have an accurate genetic algorithm:

## Genetic Representation

Pokemon in this simulation have 8 "genes", 2 [types] and 6 [base stats]. These stats represent the "genome". In the future I could deal with [IVs] or even [movesets], but those are not currently needed for this rudimentary implementation.

## Initialization

I randomly generated a fixed number of Pokemon to match the above genome. The types are randomly assigned, and the base stats are relatively random generated with the following formula in `Go`,

```go
int(math.Ceil(rand.NormFloat64() * 10 + 100) + .5)
```

## Fitness

To define the fitness of each pokemon, we implement a simulate a simplified [battle] between each pair and aggregate the results. A simplified battle goes as follows:

1. Each pokemon determines the attack that will do the most damage to its opponents
1. The Pokemon with the higher speed will always move first
1. Pokemon take turns attacking until one has "fainted" (HP is less than or equal to 0)
1. The "fitness" of each Pokemon is the leftover HP at the end of the battle
  - therefore the "winner" would have positive HP while the loser has negative HP

After each one has its fitness evaluated, a set percentage of the population is dropped off, and the rest of the population randomly breeds to fill out the rest of the population

## Genetic Operators / Crossover

We "crossover" two genomes by simply averaging their stats and adding a little noise. For the child's type, we randomly pick between the 4 types of the parents, which a chance "mutation" of a random type showing up.

## Termination

Not yet decided, but we can just run the simulation many generations and try and visualize some of the trends that may occur.

# Todo

- crossover / breeding
- multiple generations
- CLI-ize the simulation
- fancy stream / websocket / event emitter to "listen" to progress

[Wikipedia]: https://en.wikipedia.org/wiki/Genetic_algorithm#Optimization_problems
[types]: https://bulbapedia.bulbagarden.net/wiki/Type
[base stats]: https://bulbapedia.bulbagarden.net/wiki/Base_stats
[IVs]: https://bulbapedia.bulbagarden.net/wiki/Individual_values
[moveset]: https://bulbapedia.bulbagarden.net/wiki/Move
[battle]: https://bulbapedia.bulbagarden.net/wiki/Pok%C3%A9mon_battle
