package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	firstNames := []string{"Alice", "Bob", "Charlie", "Diana"}
	lastNames := []string{"Smith", "Johnson", "Williams", "Brown"}

	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]

	fmt.Printf("Nome aleatório: %s %s\n", firstName, lastName)
}
