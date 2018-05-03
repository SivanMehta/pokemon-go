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

While most of the genetic-algorithm specific code is in `simulation.go`, everything specific this domain is defined in this section. According to [Wikipedia], I need to define the following attributes to have an accurate genetic algorithm:

## Genetic Representation

Pokemon in this simulation have 8 "genes", 2 [types] and 6 [base stats]. These stats represent the "genome". In the future I could deal with [IVs] or even [movesets], but those are not currently needed for this rudimentary implementation.

## Initialization

I randomly generate a fixed number of Pokemon to match the above genome. The types are randomly assigned, and the base stats are relatively random generated. The generated stats follow a Normal distribution with mean of 100 and standard deviation of 10.

## Fitness

To define the fitness of each pokemon, we implement a simulate a simplified [battle] between each pair and aggregate the results (everyone battles everyone). Each simplified battle goes as follows:

- Each pokemon determines the attack that will do the most damage to its opponent
- The Pokemon with the higher speed will always move first:
  - Pokemon take turns attacking until one has "fainted" (HP is less than or equal to 0)
- The "fitness" of each Pokemon is the leftover HP at the end of the battle
  - Therefore the "winner" would have positive HP while the loser has negative HP

After each one has its fitness evaluated, a set percentage of the population is killed off, and the rest of the population randomly breeds to fill out the rest of the population (see next section as to how).

## Genetic Operators / Crossover

We "crossover" two genomes by simply averaging their stats and adding a little noise. For the child's types, we randomly pick between the 4 types of the parents, which a chance "mutation" of a random type showing up.

## Termination

Not yet decided, but we can just run the simulation over many generations and try and visualize some of the trends that may occur.

# Todo

- crossover / breeding
- multiple generations

# Future Work / Nice to Haves

- parameterization of the simulation (generations, kill %age, noise in crossover, etc.)
- fancy stream / websocket / event emitters to "listen" to progress
- graphs of generated data

[Wikipedia]: https://en.wikipedia.org/wiki/Genetic_algorithm#Optimization_problems
[types]: https://bulbapedia.bulbagarden.net/wiki/Type
[base stats]: https://bulbapedia.bulbagarden.net/wiki/Base_stats
[IVs]: https://bulbapedia.bulbagarden.net/wiki/Individual_values
[movesets]: https://bulbapedia.bulbagarden.net/wiki/Move
[battle]: https://bulbapedia.bulbagarden.net/wiki/Pok%C3%A9mon_battle
