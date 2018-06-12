.PHONY: make

pokemon.out:
	go build -o pokemon.out simulation.go

.PHONY: simulation
simulation: pokemon.out
	@./pokemon.out

data.csv: pokemon.out
	@./pokemon.out > data.csv

stat-progression.png: pokemon.out data.csv
	@Rscript plot.R
	@rm Rplots.pdf
