package main

import (
  "log"
  "github.com/SivanMehta/pokemon-go/battle"
)

func main() {
  result := battle.Add(1, 2)
  log.Println(result)
}