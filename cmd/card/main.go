package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"

	"github.com/i-hit/go-lesson2.4.git/pkg/card"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Print(err)
		}
	}()
	err = trace.Start(f)
	if err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

	transaction := card.MakeTransactions()

	fmt.Println("\n", card.GetMap(transaction, 1))
	fmt.Println("\n", card.GetMapByMutex(transaction, 1, 10000))
	fmt.Println("\n", card.GetMapByChannel(transaction, 1, 10000))
	fmt.Println("\n", card.GetMapByMutex2(transaction, 1, 10000))
}
