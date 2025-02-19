package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	firstNames := []string{"Alice", "Bob", "Charlie", "Diana"}
	lastNames := []string{"Smith", "Johnson", "Williams", "Brown"}

	firstName := firstNames[random.Intn(len(firstNames))]
	lastName := lastNames[random.Intn(len(lastNames))]

	fmt.Printf("Nome aleatório: %s %s\n", firstName, lastName)
}
