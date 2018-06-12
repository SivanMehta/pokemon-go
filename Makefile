pokemon.out:
	go build -o pokemon.out simulation.go

.PHONY: simulation
simulation: pokemon.out
	./pokemon.out
