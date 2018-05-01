build:
	go build -o pokemon.out simulation.go

dev: build
	./pokemon.out

simulation:
	./pokemon.out
