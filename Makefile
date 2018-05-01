build:
	go build -o pokemon.out simulation.go

dev: build
	./pokemon.out

test:
	go test -v ./battle/
