package main

import (
	"fmt"

	"github.com/i-hit/go-lesson2.4.git/pkg/card"
)

func main() {
	transaction := card.MakeTransactions()

	fmt.Println("\n", card.GetMap(transaction, 1))
	fmt.Println("\n", card.GetMapByMutex(transaction, 1, 10000))
	fmt.Println("\n", card.GetMapByChannel(transaction, 1, 10000))
	fmt.Println("\n", card.GetMapByMutex2(transaction, 1, 10000))
}
